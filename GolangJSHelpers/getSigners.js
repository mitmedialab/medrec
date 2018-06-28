const agentRegistryJson = require('../SmartContracts/build/contracts/AgentRegistry.json');
const Web3 = require('web3');
const contract = require('truffle-contract');
const Promise = require('bluebird');

let AgentRegistry = contract(agentRegistryJson);
AgentRegistry.setProvider(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));
let agentRegistry;
Promise.resolve(AgentRegistry.deployed())
  .then(_reg => {
    agentRegistry = _reg;
    return agentRegistry.getNumSigners();
  }).then(numSigners => {
    let signerPromises = [];
    for(let i = 0; i < numSigners.toNumber(); i++) {
      signerPromises.push(agentRegistry.getSigner(i));
    }
    return signerPromises;
  }).spread((...signerAccounts) => {
    process.stdout.write(JSON.stringify(signerAccounts));
  });
