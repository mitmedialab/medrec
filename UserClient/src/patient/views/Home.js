import React, { Component } from 'react';
import Tests from './Tests';
import Ethereum from '../../Ethereum';
import RPCClient from '../../RPCClient';
import './home.css';

class Home extends Component {
  constructor () {
    super();
    this.state = {
      sponsor: '',
      hasAgent: false,
      mainAccount: '',
      oneTimeAccount: '',
    };

    this.createAgent = this.createAgent.bind(this);
    this.generateAccount = this.generateAccount.bind(this);
    this.valueChanged = this.valueChanged.bind(this);
  }
  createAgent (event) {
    console.log('Creating Agent contract, please wait.....');
    event.preventDefault();
    if(this.state.hasAgent)return;
    let host;
    let agentRegContract;
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed()).then(_regContract => {
        agentRegContract = _regContract;
        //get the current set of ethereum account
        return agentRegContract.getAgentHost(this.state.sponsor);
      }).then(_host => {
        host = _host;
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host)
          .send('MedRecRemote.PatientFaucet', {Account: this.state.sponsor});
      })
      //wait for the funding transaction to go through
      .then(faucetRes => Ethereum.waitForTx(faucetRes.Txid))
      //create a new agent contract for the user
      .then(() => Ethereum.getAgent())
      .then(agent => agent.new())
      //register the agent contract in the agent registry
      .then((agentContract) => agentRegContract.setAgentContractAddr(agentContract.address))
      //wait for the registry tx to finish
      .then(txResult => Ethereum.waitForTx(txResult.tx)).then(() => {
        console.log('Agent created and registered');
        this.setState({
          sponsor: '',
          hasAgent: true,
        });
      });
  }
  generateAccount () {
    let uniqueAccount;
    Ethereum.generateAccount()
      .then(_account => {
        uniqueAccount = _account;
        return Ethereum.getAccounts();
      })
      .then(accounts => {
        return Ethereum.convertAddressToPub(accounts[0]);
      })
      .then(pubkey => {
        this.setState({oneTimeAccount: pubkey + ':' + uniqueAccount});
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
          <p>Give this unique account to your provider to initialize your account
            in the MedRec Network</p>
          {this.state.mainAccount}
          <p>When your provider tells you, continue by submitting your sponsor
            provider&#39; s account below</p>
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

    let returningSection;
    if(this.state.hasAgent) {
      returningSection = (
        <div>
          <h2>
            Single Use Account ID &nbsp;
            <span className="tooltip">
                ?
              <span className="tooltiptext">MedRec protects your privacy by
                 allowing you to generate a unique account id for every relationship.
                 There is no limit to how many can be generated. Never reuse an account id.</span>
            </span>
          </h2>
          <button onClick={this.generateAccount}>Generate</button>
          <p>{this.state.oneTimeAccount}</p>
        </div>
      );
    }

    return  (
      <div className="mainPanel">
        {firstTimeSection}
        {returningSection}
        <h2>Your medical records overview</h2>
        <p>Here you can access your most recent medical records.</p>
        <Tests/>
      </div>
    );
  }
  componentDidMount () {
    let accounts;
    Ethereum.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        this.setState({
          mainAccount: accounts[0],
        });
        return agentRegistry.getAgentContractAddr(accounts[0]);
      })
      .then(agentAddress => {
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
