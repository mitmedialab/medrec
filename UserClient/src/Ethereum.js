import Web3 from 'web3';
import contract from 'truffle-contract';
import Promise from 'bluebird';
import {keystore, signing} from 'eth-lightwallet';
import ProviderEngine from 'web3-provider-engine';
import FilterSubprovider from 'web3-provider-engine/subproviders/filters.js';
import VmSubprovider from 'web3-provider-engine/subproviders/vm.js';
import HookedWalletSubprovider from 'web3-provider-engine/subproviders/hooked-wallet.js';
import NonceSubprovider from 'web3-provider-engine/subproviders/nonce-tracker.js';
import RpcSubprovider from 'web3-provider-engine/subproviders/rpc.js';
//import the contracts descriptors
import agentJson from '../../SmartContracts/build/contracts/Agent.json';
import agentRegistryJson from '../../SmartContracts/build/contracts/AgentRegistry.json';
import relationshipJson from '../../SmartContracts/build/contracts/Relationship.json';
import Transaction from 'ethereumjs-tx';
import Utils from 'ethereumjs-util';
import {store} from './reduxStore';

class Ethereum {
  constructor () {
    this.engine = new ProviderEngine();
    this.web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));
    //this.utils = new Utils();
    this.pendingTransactions = [];
    this.addWeb3Commands();
    this.initWeb3Provider();
    //convert the contracts into usable objects
    this.Agent = contract(agentJson);
    this.AgentRegistry = contract(agentRegistryJson);
    this.Relationship = contract(relationshipJson);

    //Provision the contract with a web3 provider
    this.Agent.setProvider(this.engine);
    this.AgentRegistry.setProvider(this.engine);
    this.Relationship.setProvider(this.engine);

    //the value stores access to privateKeys and addresses
    this.vault = null;

    //saves the user password so they don't have to enter it for every transaction
    this.pwDerivedKey = null;
  }

  waitForVault () {
    return new Promise((resolve, reject) => {
      let check = () => {
        if(this.vault) {
          resolve(this.vault);
        }else {
          setTimeout(check, 200);
        }
      };
      check();
    });
  }

  waitForRPCConn () {
    return new Promise((resolve, reject) => {
      let timer;
      let check = () => {
        this.web3.eth.net.isListening().then(listening => {
          if(listening) {
            clearTimeout(timer);
            resolve();
          }
        }).catch(() => {});
      };
      timer = setInterval(check, 2000);
      check();
    });
  }

  //add some non standard web3 commands to the web3 api
  addWeb3Commands () {
    function toStringVal (val) {
      return String(val);
    }

    function toBoolVal (val) {
      if(String(val) == 'true') {
        return true;
      }
      return false;

    }

    function toIntVal (val) {
      return parseInt(val);
    }

    this.web3.extend({
      property: 'personal',
      methods: [new this.web3.extend.Method({
        name: 'unlockAccount',
        call: 'personal_unlockAccount',
        params: 3,
        inputFormatter: [this.web3.extend.utils.toAddress, toStringVal, toIntVal],
        outputFormatter: toBoolVal,
      })],
    });

    this.web3.extend({
      property: 'personal',
      methods: [new this.web3.extend.Method({
        name: 'importRawKey',
        call: 'personal_importRawKey',
        params: 2,
        inputFormatter: [this.web3.extend.utils.toAddress, toStringVal],
      })],
    });

    this.web3.extend({
      property: 'miner',
      methods: [new this.web3.extend.Method({
        name: 'start',
        call: 'miner_start',
        params: 0,
      })],
    });
  }

  //initialize the modular Web3 API provider
  initWeb3Provider () {
    //filters
    this.engine.addProvider(new FilterSubprovider());

    //pending nonce
    this.engine.addProvider(new NonceSubprovider());

    //vm
    this.engine.addProvider(new VmSubprovider());

    ////accounts management
    this.engine.addProvider(new HookedWalletSubprovider({
      getAccounts: (cb) => {
        this.getVault().then(vault => {
          if(vault) {
            cb(null, vault.getAddresses());
          }else {
            cb('no keys available, login frst', null);
          }
        });
      },
      signMessage: (options, cb) => {
        this.getVault().then(vault => {
          var secretKey = vault.exportPrivateKey(options.from, this.pwDerivedKey);
          var msg = new Buffer(options.data.replace('0x', ''), 'hex');

          let msgHash = Utils.hashPersonalMessage(msg);
          let sgn = Utils.ecsign(msgHash, new Buffer(secretKey, 'hex'));
          let signedMsgHex = Utils.toRpcSig(sgn.v, sgn.r, sgn.s);
          cb(null, signedMsgHex);
        });
      },
      signTransaction: (txObject, cb) => {
        this.getVault().then(vault => {
          txObject.gas = '0x1000000';
          txObject.gasLimit = '0x1000000';

          let tx = new Transaction(txObject);
          let rawTX = tx.serialize().toString('hex');

          let signedTx = '0x' + signing.signTx(
            vault,
            this.pwDerivedKey,
            rawTX,
            txObject.from.toLowerCase()
          );
          cb(null, signedTx);
        });
      },
    }));

    //data source
    //this must come as the last component so it doesn't overshadow the wallet
    this.engine.addProvider(new RpcSubprovider({
      rpcUrl: 'http://localhost:8545',
    }));

    //log new blocks
    this.engine.on('block', (block) => {
      let newPendTrans = [];
      this.pendingTransactions.forEach(txObj => {
        this.web3.eth.getTransactionReceipt(txObj.txid).then(receipt => {
          if(receipt && receipt.blockHash) {
            txObj.resolve(receipt);
          }else {
            newPendTrans.push(txObj);
          }
        });
      });
      this.pendingTransactions = newPendTrans;
    });

    //network connectivity error
    this.engine.on('error', function (err) {
      //report connectivity errors
      console.error(err.stack);
    });

    //wait for the ethereum client to be connected, then start polling for blocks
    this.waitForRPCConn().then(() => {
      this.web3 = new Web3(this.engine);
      this.engine.start();
    });
  }

  //gets the current user's account
  getAccounts () {
    return new Promise((resolve, reject) => {
      this.web3.eth.getAccounts((err, accounts) => {
        resolve(accounts);
      });
    });
  }

  signMessage (address, dataToSign) {
    return new Promise((resolve, reject) => {
      this.web3.eth.sign((err, signedMsgHex) => {
        resolve(signedMsgHex);
      });
    });
  }

  getAgentRegistry () {
    return new Promise((resolve, reject) => {
      this.getVault().then(vault => {
        this.AgentRegistry.web3.eth.defaultAccount = vault.getAddresses()[0];
        resolve(this.AgentRegistry);
      });
    });
  }

  getAgent () {
    return new Promise((resolve, reject) => {
      this.getVault().then(vault => {
        this.Agent.web3.eth.defaultAccount = vault.getAddresses()[0];
        resolve(this.Agent);
      });
    });
  }

  getRelationship () {
    return new Promise((resolve, reject) => {
      this.getVault().then(vault => {
        this.Relationship.web3.eth.defaultAccount = vault.getAddresses()[0];
        resolve(this.Relationship);
      });
    });
  }

  //if the user is coming back after refreshing the page then refresh the vault
  refreshVault () {
    return new Promise((resolve, reject) => {
      let data = store.getState().homeReducer;
      if(data.seed != undefined) {
        this.generateVault(data.password, data.seed, 1).then(resolve);
        return;
      }
      reject('no user is logged in');
    });
  }

  //generate a new vault
  //password: mandatory
  //seed: used to generate HD private keys, default is random generation
  //numAddresses: number of addresses to generate, default is 1
  generateVault (password, seed, numAddresses) {
    seed = seed || keystore.generateRandomSeed();
    numAddresses = numAddresses || 1;
    return new Promise((resolve, reject) => {
      //the seed is stored encrypted by a user-defined password
      keystore.createVault({
        password: password,
        seedPhrase: seed,
        hdPathString: "m/44'/60'/0'/0",
      }, (err, vault) => {
        if(err)throw err;

        //get a copy of the encryption key
        vault.keyFromPassword(password, (_, pwDerivedKey) => {
          if(err)throw err;
          //store the key for later unlocking
          //generate new address/private key pair
          //the corresponding private keys are also encrypted
          vault.generateNewAddress(pwDerivedKey, numAddresses);
          this.vault = vault;

          //reset the waitForVault function so it takes the new vault value
          //this.waitForVault = new Promise((vaultResolve) => {
          //vaultResolve(this.vault);
          //});
          //TODO: make sure the above isn't needed any more

          this.pwDerivedKey = pwDerivedKey;
          resolve(vault);
        });
      });
    });
  }

  //get the vault
  getVault () {
    return new Promise((resolve, reject) => {
      this.waitForVault().then(resolve);
    });
  }

  //generates a seed if necessary
  getSeed (password) {
    return new Promise((resolve, reject) => {
    //get the private key vault
      this.getVault(password).then(vault => {
        //get a copy of the encryption key
        vault.keyFromPassword(password, function (err, pwDerivedKey) {
          let seed = vault.getSeed(pwDerivedKey);
          resolve(seed);
        });
      });
    });
  }

  //given a transaction id, waits for the tx to make it into a block
  waitForTx (txid) {
    return new Promise((resolve, reject) => {
      this.pendingTransactions.push({txid, resolve});
    });
  }
}

export default (new Ethereum());
