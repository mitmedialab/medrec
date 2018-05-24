import React, { Component } from 'react';
import Tests from './Tests';
import RPCClient from '../../RPCClient';
import './home.css';

class Home extends Component {
  constructor () {
    super();
    this.state = {
      sponsor: '',
      hasAgent: false,
      accountAddress: '',
      agentAddress: 'No Agent Contract exists yet',
    };

    this.createAgent = this.createAgent.bind(this);
    this.valueChanged = this.valueChanged.bind(this);
  }
  createAgent (event) {
    event.preventDefault();
    if(this.state.hasAgent)return;
    let accounts;
    let host;
    let agentRegContract;
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed()).then(_regContract => {
        agentRegContract = _regContract;
        //get the current set of ethereum account
        return Ethereum.web3.eth.getAccounts();
      }).then(_acc => {
        accounts = _acc;
        return agentRegContract.getAgentHost(this.state.sponsor);
      }).then(_host => {
        host = _host;
        console.log('the agent host', this.state.sponsor);
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.PatientFaucet', {Account: this.state.sponsor});
      }).then(faucetRes => {
        console.log('have returned from faucet', faucetRes);
        //wait for the funding transaction to go through
        return Ethereum.waitForTx(faucetRes.Txid);
      }).then( () => {
        //create a new agent contract for the user
        return Ethereum.getAgent();
      }).then(agent => agent.new()).then((agentContract) => {
        //register the agent contract in the agent registry
        return agentRegContract.setAgentContractAddr(agentContract.address);
      }).then(txResult => {
        //wait for the registry tx to finish
        return Ethereum.waitForTx(txResult.tx);
      }).then(() => {
        console.log('Agent created and registered');
        console.log(this.state.contract);
        this.setState({
          sponsor: '',
          hasAgent: true,
        });
      });
  }
  valueChanged (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  render () {
    let firstTimeSection;
    if(!this.state.hasAgent) {
      firstTimeSection = (
        <div>
          <h2>First time account setup</h2>
          <p>your account must be sponsored by a provider, enter your sponsor's account below</p>
          <form>
            <label>
              <span>sponsor</span>
              <input id="sponsor" onChange={this.valueChanged} value={this.state.sponsor}/>
            </label>
            <input type="submit" onClick={this.createAgent}/>
          </form>
        </div>
      );
    }

    return  (
      <div className="mainPanel">
        <h2>Local Account</h2>
        <p>{this.state.accountAddress}</p>
        <h2>Agent Contract Address</h2>
        <p>{this.state.agentAddress}</p>
        {firstTimeSection}
        <h2>Your medical records overview</h2>
        <p>Here you can access your most recent medical records.</p>
        <Tests/>
      </div>
    );
  }
  componentDidMount () {
    let accounts;
    Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        this.setState({
          accountAddress: accounts[0],
        });
        return agentRegistry.getAgentContractAddr(accounts[0]);
      })
      .then(agentAddress => {
        console.log('agentAddress', agentAddress);
        if(parseInt(agentAddress, 16) != 0) {
          this.setState({
            hasAgent: true,
            agentAddress: agentAddress,
          });
        }
      });
  }
}

export default Home;
