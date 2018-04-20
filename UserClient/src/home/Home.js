import React, { Component } from 'react';
import {connect} from 'react-redux';
import Switch from 'react-toggle-switch';
import RPCClient from '../RPCClient';
import Ethereum from '../Ethereum';
import './home.css';
import {SET_USER} from '../constants';
import classnames from 'classnames';

class Home extends Component {
  constructor (props) {
    super(props);
    this.state = {
      users: [],
      loginUser: '',
      loginPassword: '',
      fname: '',
      lname: '',
      username: '',
      password: '',
      agentID: '',
      seed: '',
      mode: 'patient',
      switched: false,
      enableModal: false,
      enableSeedModal: true,
      enableConfirmModal: false,
      enablePreviewModal: false,
      confirmSeed: '',
      confirmSeedError: '',
      sponsor: '',
      contract: 'agent',
      preview: 'agent',
    };
    this.createAgent = this.createAgent.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
    this.openConfirmModal = this.openConfirmModal.bind(this);
    this.closeModal = this.closeModal.bind(this);
    this.finishCreateAgent = this.finishCreateAgent.bind(this);
    this.login = this.login.bind(this);
    this.toggleSwitch = this.toggleSwitch.bind(this);
    this.selectContract = this.selectContract.bind(this);
    this.enablePreviewModal = this.enablePreviewModal.bind(this);
  }

  toggleSwitch () {
    this.setState(prevState => {
      return {
        switched: !prevState.switched,
        mode: !prevState.switched ? 'provider' : 'patient',
      };
    });
  }

  selectContract (event) {
    this.setState({contract: event.target.value});
  }

  enablePreviewModal (event) {
    this.setState({enablePreviewModal: true});
    switch (event.target.id) {
      case 'group':
        this.setState({preview: 'group'});
        break;
      case 'DMswitch':
        this.setState({preview: 'DMswitch'});
        break;
      default:
        this.setState({preview: 'agent'});
    }
  }

  componentDidMount () {
    RPCClient.send('MedRecLocal.GetUsernames')
      .then(res => this.setState({ users: res.Usernames}));
  }

  createAgent (event) {
    event.preventDefault();
    Ethereum.generateVault(this.state.password)
      .then(() => Ethereum.getSeed(this.state.password))
      .then((seed) => {
        //shows seed to user
        /*eslint max-len: ["error", { "code": 200 }]*/
        this.setState({
          seed: seed,
          enableModal: true,
          enableSeedModal: true,
        });
      });
  }

  changeFieldById (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }

  closeModal () {
    this.setState({
      enableModal: false,
      enableSeedModal: true,
      enableConfirmModal: false,
      enablePreviewModal: false,
    });
  }

  openConfirmModal () {
    this.setState({
      enableSeedModal: false,
      enableConfirmModal: true,
    });
  }

  finishCreateAgent (event) {
    event.preventDefault();
    //check the user entered seed
    let result = this.state.seed.localeCompare(this.state.confirmSeed);
    if(result != 0) {
      this.setState({confirmSeedError: 'Seed incorrect, please try again.'});
      return;
    }

    this.setState({
      enableConfirmModal: false,
      enableModal: false,
      confirmSeed: '',
      confirmSeedError: '',
    });
    //update the local database with the new user

    let agentRegContract;
    let accounts;
    RPCClient.send('MedRecLocal.NewUser', {
      FirstName: this.state.fname,
      LastName: this.state.lname,
      Username: this.state.username,
      Password: this.state.password,
      Seed: this.state.seed,
      Contract: this.state.contract,
    }).then( (res) => {
      if(res.Error !== '') {
        throw (res.Error);
      }
      return Ethereum.getAgentRegistry();
    }).then(reg => reg.deployed())
      .then(_regContract => {
        agentRegContract = _regContract;
        //get the current set of ethereum account
        return Ethereum.web3.eth.getAccounts();
      }).then(_acc => {
        accounts = _acc;
        //instantiate the account and UID in the key/val store

        RPCClient.remote('127.0.0.1').send('MedRecRemote.AddAccount', {
          Account: accounts[0],
          AgentID: this.state.agentID,
        });
        return agentRegContract.getAgentHost(this.state.sponsor);
      }).then(host => {
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.Faucet', {Account: accounts[0]});
      })
      .then(faucetRes => {
        //wait for the funding transaction to go through
        return Ethereum.waitForTx(faucetRes.Txid);
      })
      .then( () => {
        //create a new agent contract for the user
        return Ethereum.getAgent();
      })
      .then(agent => agent.new())
      .then((agentContract) => {
        //register the agent contract in the agent registry
        return agentRegContract.setAgentContractAddr(agentContract.address);
      })
      .then(txResult => {
        //wait for the registry tx to finish
        return Ethereum.waitForTx(txResult.tx);
      })
      .then(() => {
        console.log('Agent created and registered');
        console.log(this.state.contract);
        this.setState({
          enableModal: false,
          enableConfirmModal: false,
          fname: '',
          lname: '',
          username: '',
          password: '',
          sponsor: '',
          agentID: '',
        });
      });
  }

  login () {
    RPCClient.send('MedRecLocal.GetSeed', {
      Username: this.state.loginUser,
      Password: this.state.loginPassword,
    }).then( (res) => {
      if(res.Error !== '') {
        console.log(res.Error);
      }else {
        this.props.dispatch({
          type: SET_USER,
          username: this.state.loginUser,
          password: this.state.loginPassword,
          seed: res.Seed,
          contract: this.state.contract,
        });
        //update the vault with the new user credentials
        Ethereum.refreshVault().then(() => {
          this.props.history.push('/' + this.state.mode + '/home');
        });
      }
    });
  }

  render () {
    let loginStyle = this.state.mode === 'patient' ? 'loginPatient' : 'loginProvider';
    let identityStyle = this.state.mode === 'patient' ? 'identityPatient' : 'identityProvider';
    let contractStyle = this.state.mode === 'patient' ? 'contractPatient' : 'contractProvider';
    return (
      <main>
        <div id="homepageWrapper">
          <div id="sidePane">
            <h1>MedRec</h1>
            <h2>A decentralized medical records management system.</h2>
            <p>v1.0</p>
            <div id="logoSprite"></div>
          </div>
          <div  id="loginPane" >
            <div id={loginStyle}>
              <Switch onClick={this.toggleSwitch} on={this.state.switched}/>
              <h3>Login as {this.state.mode}</h3>
              {/*{users}*/}
              <div id="loginUserPassContainer">
                <div id="loginUserPass">
                  <input className="inputStyle" id="loginUser" onChange={this.changeFieldById} placeholder="username"
                    value={this.state.loginUser}/>
                  <input className="inputStyle" id="loginPassword" onChange={this.changeFieldById} type="password" placeholder="password" value={this.state.loginPassword}/>
                  <button className="loginStyle" onClick={this.login}>Login</button>
                </div>
              </div>
            </div>
            <div id={contractStyle}>
              <h3>Select contract to be deployed</h3>
              <div>
                <input type="radio" checked={this.state.contract === 'agent'} value="agent" onChange={this.selectContract}></input>Default
                <span className="preview" id="agent" onClick={this.enablePreviewModal}>Preview</span>
              </div>
              <div>
                <input type="radio" checked={this.state.contract === 'group'} value="group" onChange={this.selectContract}></input>Group
                <span className="preview" id="group" onClick={this.enablePreviewModal}>Preview</span>
              </div>
              <div>
                <input type="radio" checked={this.state.contract === 'DMswitch'} value="DMswitch" onChange={this.selectContract}></input>DM Switch
                <span className="preview" id="DMswitch" onClick={this.enablePreviewModal}>Preview</span>
              </div>
            </div>
            <div id={identityStyle}>
              <h3>Create an Identity using the contract</h3>
              <form onSubmit={this.createAgent}>
                <input className="inputStyle" id="sponsor" placeholder="sponsor's account" onChange={this.changeFieldById} value={this.state.sponsor}/>
                <br/>
                <input className="inputStyle" id="agentID" placeholder="unique ID" onChange={this.changeFieldById} value={this.state.agentID}/>
                <br/>
                <input className="inputStyle" id="fname" placeholder="First name" onChange={this.changeFieldById} value={this.state.fname}/>
                <br/>
                <input className="inputStyle" id="lname" placeholder="Last name" onChange={this.changeFieldById} value={this.state.lname}/>
                <br/>
                <input className="inputStyle" id="username" placeholder="username" onChange={this.changeFieldById} value={this.state.username}/>
                <br/>
                <input className="inputStyle" id="password" type="password" placeholder="password" onChange={this.changeFieldById} value={this.state.password}/>
                <button className="loginStyle" type="submit">Create</button>
              </form>
            </div>
          </div>
        </div>
        <div id="seedModal" className={classnames({display: this.state.enableModal})}>
          <div onClick={this.closeModal} className="modalBackground"></div>
          <div className={classnames({modalContent: true, display: this.state.enableSeedModal})}>
            <button className="close" onClick={this.closeModal}></button>
            <p>The following seed can be used to recover your account, be sure to write it down and keep it somewhere safe.</p>
            <p className="seed">{this.state.seed}</p>
            <button className="buttonStyle" onClick={this.openConfirmModal}>Next</button>
          </div>
          <div className={classnames({modalContent: true, display: this.state.enableConfirmModal})}>
            <p>{this.state.confirmSeedError}</p>
            <button className="close" onClick={this.closeSeedModal}></button>
            <p>Please reenter your seed, to ensure you copied it correctly.</p>
            <input className="inputStyle" id="confirmSeed" onChange={this.changeFieldById} value={this.state.confirmSeed}/>
            <button className="buttonStyle" onClick={this.finishCreateAgent}>Finish</button>
          </div>
        </div>
        <div id="previewModal" className={classnames({display: this.state.enablePreviewModal})}>
          <div onClick={this.closeModal} className="modalBackground"></div>
          <button className="close" onClick={this.closeModal}></button>
          <div className={classnames({modalContent: true, display: this.state.enablePreviewModal})}>
            <button className="close" onClick={this.closeModal}></button>
            <div className="wrapper"><div className={this.state.preview}></div></div>
          </div>
        </div>
      </main>
    );
  }
}
export default connect()(Home);
