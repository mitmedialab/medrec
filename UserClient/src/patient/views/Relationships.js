import React, {Component} from 'react';
import {connect} from 'react-redux';
import Ethereum from '../../Ethereum';
import RPCClient from '../../RPCClient';
import {SET_RELATIONSHIP, SET_VIEWER} from '../../constants';

class Permissions extends Component {
  constructor () {
    super();
    this.state = {
      permissionNames: [],
      permissionDurations: [],
      permissionStartTimes: [],
      permissionName: '',
      permissionStartTime: 0,
      permissionDuration: 0,
      host: '',
    };

    this.addPermission = this.addPermission.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
  }
  addPermission (event) {
    event.preventDefault();
    let accounts;
    let relationship;
    let providerAddr;
    return Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getRelationship();
      })
      .then(relationGenerator => {
        relationship = relationGenerator.at(this.props.relationshipAcc);
        return relationship.provider();
      }).then(_provAddr => {
        providerAddr = _provAddr;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentHost(providerAddr))
      .then(host => {
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.AddPermission', {
          AgentID: '0', //TODO: get the actual agentID
          ViewerGroup: this.props.viewerGroup,
          Name: this.state.permissionName,
          StartTime: this.state.permissionStartTime,
          DurationDays: this.state.permissionDuration,
        });
      })
      .then((permRes) => {
        this.setState({
          permissionName: '',
          permissionRead: true,
          permissionWrite: false,
          permissionDuration: '',
        });
        this.update(this.props);
      })
      .catch((err) => {
        console.log(err);
      });
  }
  changeFieldById (event) {
    let state = this.state;
    switch (event.target.id) {
      case 'permissionRead':
      case 'permissionWrite':
        state[event.target.id] = event.target.checked;
        break;
      case 'permissionDuration':
        if(/^[0-9]+$/.test(event.target.value)) {
          state[event.target.id] = parseInt(event.target.value);
        }else if(event.target.value === '') {
          state[event.target.id] = '';
        }
        break;
      default:
        state[event.target.id] = event.target.value;
    }
    this.setState(state);
    this.update(this.props);
  }
  render () {
    if(this.props.viewerGroup == undefined) {
      return (<div></div>);
    }

    if(this.props.status == true) {
      return (<div></div>);
    }
    let permissions = [];
    if(this.props.viewerGroup !== undefined) {
      for(let i = 0; i < this.state.permissionNames.length; i ++) {
        permissions.push(
          <div className="tab" key={i}>
            <span><strong>Name</strong>: {this.state.permissionNames[i]}</span>
            <span><strong>Duration</strong>: {this.state.permissionDurations[i]} days</span>
            <span><strong>Start Time</strong>: {this.state.permissionStartTimes[i]} days</span>
          </div>
        );
      }
    }
    return (
      <div className="settingsColumn">
        <div className="header">
          <p>Permissions</p>
        </div>
        <div className="content">
          {permissions}
        </div>
        <div className="edit">
          <p>Edit Permissions</p>
          <form onSubmit={this.addPermission}>
            <label>
              <input id="permissionName" onChange={this.changeFieldById}
                value={this.state.permissionName} placeholder="Name"/>
            </label>
            <label>
              <span>Duration</span>
              <input id="permissionDuration" onChange={this.changeFieldById}
                value={this.state.permissionDuration}/>
              <span>days</span>
            </label>
            <button className="buttonStyle" type="submit">Add</button>
          </form>
        </div>
      </div>
    );
  }
  componentDidMount () {
    if(this.props.relationshipAcc === undefined)return;

    let accounts;
    let relationship;
    let providerAddr;
    Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getRelationship();
      })
      .then(relationGenerator => {
        relationship = relationGenerator.at(this.props.relationshipAcc);
        return relationship.provider();
      }).then(_provAddr => {
        providerAddr = _provAddr;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentHost(providerAddr))
      .then(host => {
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.PatientFaucet', {Account: providerAddr});
      });

    this.update(this.props);
  }
  componentWillReceiveProps (nextProps) {
    this.update(nextProps);
  }
  update (props) {
    if(props.viewerGroup == undefined) {
      return;
    }
    let accounts;
    let relationship;
    let providerAddr;
    Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getRelationship();
      })
      .then(relationGenerator => {
        relationship = relationGenerator.at(this.props.relationshipAcc);
        return relationship.provider();
      }).then(_provAddr => {
        providerAddr = _provAddr;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentHost(providerAddr))
      .then(host => {
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.GetPermissions', {
          AgentID: '0', //TODO: get the actual agentID
          ViewerGroup: this.props.viewerGroup,
        });
      })
      .then(result => {
        let permissionNames = [];
        let permissionDurations = [];
        let permissionStartTimes = [];

        if(result.Permissions) {
          result.Permissions.forEach(perm => {
            permissionNames.push(perm.Name);
            permissionDurations.push(perm.Duration);
            permissionStartTimes.push(perm.StartTimes);
          });
        }

        this.setState({permissionNames, permissionDurations, permissionStartTimes});
      });
  }
}
Permissions = connect((state) => {
  return {
    relationshipAcc: state.patientReducer.relationshipAcc,
    viewerGroup: state.patientReducer.viewerGroup,
  };
})(Permissions);
class Viewers extends Component {
  constructor () {
    super();
    this.state = {
      viewerNames: '',
      viewerAccounts: '',
      viewerName: '',
      viewerAccount: '',
    };

    this.addViewer = this.addViewer.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
    this.selectViewer = this.selectViewer.bind(this);
  }
  addViewer (event) {
    event.preventDefault();
    let accounts;
    let relationship;
    let providerAddr;
    return Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getRelationship();
      })
      .then(relationGenerator => {
        relationship = relationGenerator.at(this.props.relationshipAcc);
        return relationship.provider();
      }).then(_provAddr => {
        providerAddr = _provAddr;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentHost(providerAddr))
      .then(host => {
        //send a message to the faucet to fund the new account
        return RPCClient.remote(host).send('MedRecRemote.PatientFaucet', {Account: providerAddr});
      })
      .then(faucetRes => {
        //wait for the funding transaction to go through
        return Ethereum.waitForTx(faucetRes.Txid);
      })
      .then( () => relationship.addViewerGroup())
      .then(() => {
        return relationship
          .addViewer(this.state.viewerName, this.state.viewerAccounts.length, this.state.viewerAccount);
      }).then((viewerRes) => {
        this.setState({viewerName: '', viewerAccount: ''});
        return Ethereum.waitForTx(viewerRes.tx);
      }).then(this.update(this.props));
  }
  changeFieldById (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  selectViewer (event) {
    this.props.dispatch({type: SET_VIEWER, viewerGroup: event.target.value});
  }
  render () {
    let viewers = [];
    if(this.props.status == true) {
      return (<div></div>);
    }

    if(this.props.relationshipAcc !== undefined) {
      for(let i = 0; i < this.state.viewerNames.length; i ++) {
        viewers.push(
          <div className="tab" key={i}>
            <span><input type="radio"
              checked={i == this.props.viewerGroup}
              onChange={this.selectViewer}
              value={i}></input>
            <strong>Name</strong> {this.state.viewerNames[i]}
            </span>
            <span><strong>Account</strong>{this.state.viewerAccounts[i]}</span>
          </div>
        );
      }
    }

    return (
      <div className="settingsColumn">
        <div className="header">
          <p>Viewers</p>
        </div>
        <div className="content">
          {viewers}
        </div>
        <div className="edit">
          <p>Add Viewer</p>
          <form onSubmit={this.addViewer}>
            <label>
              <input id="viewerName" onChange={this.changeFieldById}
                value={this.state.viewerName} placeholder="Enter account name"/>
            </label>
            <label>
              <input id="viewerAccount" onChange={this.changeFieldById}
                value={this.state.viewerAccount} placeholder="Enter account"/>
            </label>
            <button className="buttonStyle"  type="submit">Add</button>
          </form>
        </div>
      </div>
    );
  }
  componentDidMount () {
    this.update(this.props);
  }
  componentWillReceiveProps (nextProps) {
    this.update(nextProps);
  }
  update (props) {
    if(props.relationshipAcc == undefined) {
      return;
    }

    let relationship;
    let viewerAccounts;
    Ethereum.getRelationship()
      .then(relationGenerator => {
        relationship = relationGenerator.at(props.relationshipAcc);
        return relationship.getNumViewerGroups();
      })
      .then(numViewers => {
        let viewers = [];
        for(let i = 0; i < numViewers.toNumber(); i++) {
          viewers.push(relationship.getViewer(i, 0));
        }
        return viewers;
      })
      .spread((..._viewAcc) => {
        viewerAccounts = _viewAcc;
        let names = viewerAccounts.map(account => {
          return relationship.getViewerName(account);
        });
        return names;
      }).spread((...viewerNames) => {
        this.setState({viewerNames, viewerAccounts});
      });
  }
}
Viewers = connect((state) => {
  return {
    relationshipAcc: state.patientReducer.relationshipAcc,
    viewerGroup: state.patientReducer.viewerGroup,
  };
})(Viewers);

class Relationships extends Component {
  constructor () {
    super();
    this.state = {
      providerNames: '',
      providerAccounts: '',
      providerAccount: '',
      providerName: '',
    };

    this.createRelationship = this.createRelationship.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
    this.selectRelationship = this.selectRelationship.bind(this);
  }
  createRelationship (event) {
    event.preventDefault();
    let agent;
    let agentRegistry;
    let relationship;
    let accounts;

    Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(_agentReg => {
        agentRegistry = _agentReg;
        return agentRegistry.getAgentContractAddr(accounts[0]);
      })
      .then(agentAddress => {
        return Ethereum.getAgent()
          .then(agentContract => agentContract.at(agentAddress));
      }).then(_agent => {
        agent = _agent;
        return agentRegistry.getAgentHost(this.state.providerAccount);
      })
      .then(host => RPCClient.remote(host)
        .send('MedRecRemote.PatientFaucet', {Account: this.state.providerAccount})
      )
      .then(faucetRes => {
        //wait for the funding transaction to go through
        return Ethereum.waitForTx(faucetRes.Txid);
      })
      .then( () => Ethereum.getRelationship())
      .then(relationGenerator => relationGenerator.new(this.state.providerAccount))
      .then(_relate => {
        relationship = _relate;
        //TODO: don't store this in plaintext
        return relationship.setProviderName(this.state.providerName);
      })
      .then(() => {
        return agent.addRelationship(relationship.address);
      })
      .then((relationTx) => {
        this.setState({
          providerAccount: '',
          providerName: '',
        });
        return Ethereum.waitForTx(relationTx.tx);
      }).then(this.update.bind(this));
  }
  changeFieldById (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  selectRelationship (event) {
    this.props.dispatch({type: SET_VIEWER, viewerGroup: null});
    this.props.dispatch({type: SET_RELATIONSHIP, relationshipAcc: event.target.value});
  }
  render () {
    let relationships = [];
    if(this.props.status == true) {
      return (<div></div>);
    }
    for(let i = 0; i < this.state.providerNames.length; i ++) {
      relationships.push((
        <div className="tab" key={i}>
          <span>
            <input type="radio"
              checked={this.state.providerAccounts[i]
                .localeCompare(this.props.relationshipAcc) === 0}
              onChange={this.selectRelationship}
              value={this.state.providerAccounts[i]}></input>
            <strong>Name</strong> {this.state.providerNames[i]}
          </span>
          <span><strong>Account</strong>{this.state.providerAccounts[i]}</span>
        </div>
      ));
    }

    return (
      <div className="settingsColumn">
        <div className="header">
          <p>Relationships</p>
        </div>
        <div className="content">
          {relationships}
        </div>
        <div className="edit">
          <p>Add Relationship</p>
          <form onSubmit={this.createRelationship}>
            <label>
              <input id="providerName" onChange={this.changeFieldById}
                value={this.state.providerName} placeholder="Enter Provider Name"/>
            </label>
            <label>
              <input id="providerAccount" onChange={this.changeFieldById}
                value={this.state.providerAccount} placeholder="Enter Provider Identifier"/>
            </label>
            <button className="buttonStyle" type="submit"> Create </button>
          </form>
        </div>
      </div>
    );
  }
  componentDidMount () {
    this.update();
  }
  update () {
    let accounts;
    let agent;
    let relationGenerator;
    let providerAccounts;
    return Ethereum.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentContractAddr(accounts[0]))
      .then(agentAddress => {
        return Ethereum.getAgent()
          .then(agentContract => agentContract.at(agentAddress));
      }).then(_agent => {
        agent = _agent;
        return Ethereum.getRelationship();
      })
      .then(_generator => {
        relationGenerator = _generator;
        return agent.getNumRelationships();
      })
      .then(numRelationships => {
        let relationships = [];
        for(let i = 0; i < numRelationships.toNumber(); i++) {
          relationships.push(agent.relationships(i));
        }
        return relationships;
      })
      .spread((..._provAcc) => {
        providerAccounts = _provAcc;
        return providerAccounts.map(account => {
          return relationGenerator.at(account).providerName();
        });
      }).spread((...providerNames) => {
        this.setState({providerNames, providerAccounts});
      });
  }
}
Relationships = connect((state) => {
  return {relationshipAcc: state.patientReducer.relationshipAcc};
})(Relationships);

class DeadManSwitch extends Component {
  constructor () {
    super();
    this.state = {
      providerNames: '',
      providerAccounts: '',
      providerAccount: '',
      providerName: '',
    };

    this.createRelationship = this.createRelationship.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
    this.selectRelationship = this.selectRelationship.bind(this);
    this.DMswitch = this.DMswitch.bind(this);
  }
  createRelationship (event) {
    event.preventDefault();
    let agent;
    let agentRegistry;
    let relationship;
    let accounts;

    Ethereum.web3.eth.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(_agentReg => {
        agentRegistry = _agentReg;
        return agentRegistry.getAgentContractAddr(accounts[0]);
      })
      .then(agentAddress => {
        return Ethereum.getAgent()
          .then(agentContract => agentContract.at(agentAddress));
      }).then(_agent => {
        agent = _agent;
        return agentRegistry.getAgentHost(this.state.providerAccount);
      })
      .then(host => RPCClient.remote(host)
        .send('MedRecRemote.PatientFaucet', {Account: this.state.providerAccount})
      )
      .then(faucetRes => {
        //wait for the funding transaction to go through
        return Ethereum.waitForTx(faucetRes.Txid);
      })
      .then( () => Ethereum.getRelationship())
      .then(relationGenerator => relationGenerator.new(this.state.providerAccount))
      .then(_relate => {
        relationship = _relate;
        return relationship.setProviderName(this.state.providerName);
      }).then(() => {
        return agent.addRelationship(relationship.address);
      }).then((relationTx) => {
        this.setState({
          providerAccount: '',
          providerName: '',
        });
        return Ethereum.waitForTx(relationTx.tx);
      }).then(this.update.bind(this));
  }
  changeFieldById (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  selectRelationship (event) {
    this.props.dispatch({type: SET_VIEWER, viewerGroup: null});
    this.props.dispatch({type: SET_RELATIONSHIP, relationshipAcc: event.target.value});
  }
  DMswitch () {
    console.log('Dead Man switch button was clicked today');
  }
  render () {
    let relationships = [];
    if(this.props.status == true) {
      return (<div></div>);
    }
    for(let i = 0; i < this.state.providerNames.length; i ++) {
      relationships.push((
        <div className="tab" key={i}>
          <span>
            <input type="radio"
              checked={this.state.providerAccounts[i]
                .localeCompare(this.props.relationshipAcc) === 0}
              onChange={this.selectRelationship}
              value={this.state.providerAccounts[i]}></input>
            <strong>Name</strong> {this.state.providerNames[i]}
          </span>
          <span><strong>Account</strong>{this.state.providerAccounts[i]}</span>
        </div>
      ));
    }

    return (
      <div className="DMswitch">
        <div className="settingsColumn">
          <button className="DMbutton" onClick={this.DMswitch}>Click this once per day</button>
          <p>If you fail to click this button every day, the Dead Man Switch contract will be activated.</p>
        </div>
        <div className="settingsColumn">
          <div className="header">
            <p>Relationships</p>
          </div>
          <div className="content">
            {relationships}
          </div>
          <div className="edit">
            <p>Add Relationship</p>
            <form onSubmit={this.createRelationship}>
              <label>
                <input id="providerName" onChange={this.changeFieldById}
                  value={this.state.providerName} placeholder="Enter Provider Name"/>
              </label>
              <label>
                <input id="providerAccount" onChange={this.changeFieldById}
                  value={this.state.providerAccount} placeholder="Enter Provider Identifier"/>
              </label>
              <button className="buttonStyle" type="submit"> Add </button>
            </form>
          </div>
        </div>
      </div>
    );
  }
  componentDidMount () {
    this.update();
  }
  update () {
    let accounts;
    let agent;
    let relationGenerator;
    let providerAccounts;
    return Ethereum.getAccounts()
      .then(_acc => {
        accounts = _acc;
        return Ethereum.getAgentRegistry();
      })
      .then(reg => reg.deployed())
      .then(agentRegistry => agentRegistry.getAgentContractAddr(accounts[0]))
      .then(agentAddress => {
        return Ethereum.getAgent()
          .then(agentContract => agentContract.at(agentAddress));
      }).then(_agent => {
        agent = _agent;
        return Ethereum.getRelationship();
      })
      .then(_generator => {
        relationGenerator = _generator;
        return agent.getNumRelationships();
      })
      .then(numRelationships => {
        let relationships = [];
        for(let i = 0; i < numRelationships.toNumber(); i++) {
          relationships.push(agent.relationships(i));
        }
        return relationships;
      })
      .spread((..._provAcc) => {
        providerAccounts = _provAcc;
        return providerAccounts.map(account => {
          return relationGenerator.at(account).providerName();
        });
      }).spread((...providerNames) => {
        this.setState({providerNames, providerAccounts});
      });
  }
}
DeadManSwitch = connect((state) => {
  return {relationshipAcc: state.patientReducer.relationshipAcc};
})(DeadManSwitch);

class MainWindow extends Component {
  render () {
    let relationshipsHidden = false;
    let viewersHidden = false;
    let permissionsHidden = false;
    let switchHidden = true;
    if(this.props.contract === 'DMswitch') {
      permissionsHidden = true;
      viewersHidden = true;
      relationshipsHidden = true;
      switchHidden = false;
    }
    return (
      <div className="mainPanel">
        <h2>Contract overview</h2>
        <div className="relationshipsTable">
          <DeadManSwitch status={switchHidden} />
          <Relationships status={relationshipsHidden} />
          <Viewers status={viewersHidden} />
          <Permissions status={permissionsHidden} />
        </div>
      </div>
    );
  }
}

export default MainWindow;
