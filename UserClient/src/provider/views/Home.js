import React, { Component } from 'react';
import Ethereum from '../../Ethereum';
import Promise from 'bluebird';
import './home.css';

class Home extends Component {
  constructor () {
    super();
    this.state =  {
      proposedSignerAccounts: [],
      proposedSignerNames: [],
      proposedSignerVotes: [],
      kickedSignerAccounts: [],
      kickedSignerNames: [],
      kickedSignerVotes: [],
      signerNames: [],
      signerAccounts: [],
      proposalName: '',
      kickAddress: '',
      primaryAccount: '',
      futureProposer: '',
    };
    this.proposeSelf = this.proposeSelf.bind(this);
    this.kick = this.kick.bind(this);
    this.changeFieldById = this.changeFieldById.bind(this);
    this.sponsor = this.sponsor.bind(this);
    this.changeSignerVotes = this.changeSignerVotes.bind(this);
    this.voteForProposedSigner = this.voteForProposedSigner.bind(this);
    this.voteForKickedSigner = this.voteForKickedSigner.bind(this);
  }
  sponsor (event) {
    event.preventDefault();
    Promise.promisify(Ethereum.web3.eth.sendTransaction)({
      from: this.state.primaryAccount,
      to: this.state.futureProposer,
      value: 1e15, //arbitrary amount of eth that should be enough for a provider
    }).then(txid => Ethereum.waitForTx(txid))
      .then(() => this.setState({futureProposer: ''}));
  }
  proposeSelf (event) {
    event.preventDefault();
    Ethereum.getAccounts()
      .then(() => Ethereum.getAgentRegistry())
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        return agentRegistry.propose(this.state.proposalName);
      })
      .then(res => {
        this.setState({proposalName: ''});
        this.updateProspectivesList();
      }, err => {
        console.log('there was an error with the tx');
        console.log(err);
      });
  }
  kick (event) {
    event.preventDefault();
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        return agentRegistry.kick(this.state.kickAddress);
      })
      .then(res => {
        this.setState({kickAddress: ''});
        this.updateProspectivesList();
        this.updateSignersList();
      }, err => {
        console.log('there was an error with the tx');
        console.log(err);
      });
  }
  voteForProposedSigner (index) {
    return (event) => {
      event.preventDefault();

      Ethereum.getAgentRegistry()
        .then(reg => reg.deployed())
        .then(agentRegistry => {
          return agentRegistry.vote(this.state.proposedSignerAccounts[index],
            this.state.proposedSignerVotes[index]);
        });
    };
  }
  voteForKickedSigner (index) {
    return (event) => {
      event.preventDefault();
      Ethereum.getAgentRegistry()
        .then(reg => reg.deployed())
        .then(agentRegistry => {
          return agentRegistry.vote(this.state.kickedSignerAccounts[index],
            this.state.kickedSignerVotes[index]);
        }).then(res => {
          console.log(res);
        });
    };
  }
  changeFieldById (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  changeSignerVotes (index) {
    return (event) => {
      let state = this.state;
      state.proposedSignerVotes[index] = event.target.value === 'true';
      this.setState(state);
    };
  }
  changeKickedVotes (index) {
    return (event) => {
      let state = this.state;
      state.kickedSignerVotes[index] = event.target.value === 'true';
      this.setState(state);
    };
  }
  render () {
    let i;
    let currentSigners = [];
    for(i = 0; i < this.state.signerNames.length; i++) {
      currentSigners.push((
        <div key={this.state.signerAccounts[i]}>
          <span><strong>Name</strong>: {this.state.signerNames[i]} </span>
          <span><strong>Account</strong>: {this.state.signerAccounts[i]} </span>
        </div>
      ));
    }

    let proposedSigners = [];
    for(i = 0; i < this.state.proposedSignerNames.length; i++) {
      proposedSigners.push((
        <form onSubmit={this.voteForProposedSigner(i)} key={this.state.proposedSignerAccounts[i]}>
          <div>
            <span>Name: {this.state.proposedSignerNames[i]}</span>
            <span>Account: {this.state.proposedSignerAccounts[i]}</span>
          </div>
          <div>
            <label>
              <span>Approve</span>
              <input
                type="radio" name={'prospectiveVote' + i}
                value="true" onChange={this.changeSignerVotes(i)}
                checked={this.state.proposedSignerVotes[i]}/>
            </label>
            <label>
              <span>Reject</span>
              <input type="radio" name={'prospectiveVote' + i}
                value="false" onChange={this.changeSignerVotes(i)}
                checked={!this.state.proposedSignerVotes[i]}/>
            </label>
          </div>
          <button className="buttonStyle" type="submit">Vote</button>
        </form>
      ));
    }

    let kickedSigners = [];
    for(i = 0; i < this.state.kickedSignerNames.length; i++) {
      kickedSigners.push((
        <form onSubmit={this.voteForKickedSigner(i)} key={this.state.kickedSignerAccounts[i]}>
          <div>
            <span>Name: {this.state.kickedSignerNames[i]}</span>
            <span>Account: {this.state.kickedSignerAccounts[i]}</span>
          </div>
          <div>
            <label>
              <span>Approve</span>
              <input type="radio" name={'kickedVote'} value="true"
                onChange={this.changeKickedVotes(i)} checked={this.state.kickedSignerVotes[i]}/>
            </label>
            <label>
              <span>Reject</span>
              <input type="radio" name={'kickedVote'} value="false"
                onChange={this.changeKickedVotes(i)}
                checked={!this.state.kickedSignerVotes[i]}/>
            </label>
          </div>
          <button className="buttonStyle" type="submit">Vote</button>
        </form>
      ));
    }

    return  (
      <div className="mainPanel">
        <h2>Your provider account:</h2>
        {this.state.primaryAccount}
        <h2>Sponsor a new Provider</h2>
        <form onSubmit={this.sponsor}>
          <label>
            <input className="inputStyle" id="futureProposer" onChange={this.changeFieldById}
              value={this.state.futureProposer} placeholder="Address"/>
          </label>
          <button className="buttonStyle" type="submit"> Sponsor </button>
        </form>
        <div id="signersList">
          <h2>Registered Providers</h2>
          The current signers are the following:
          {currentSigners}
        </div>
        <div id="proposeKickSection">
          <div>
            <h2>Propose yourself as a provider</h2>
            <form onSubmit={this.proposeSelf}>
              <label>
                <input className="inputStyle" id="proposalName" onChange={this.changeFieldById}
                  value={this.state.proposalName} placeholder="Name"/>
              </label>
              <button className="buttonStyle" type="submit"> Propose Self </button>
            </form>
            <h2>Proposed Signers</h2>
            Here is a list of the proposed signers:
            {proposedSigners}
          </div>
          <div>
            <h2>Kick a registered provider</h2>
            <form onSubmit={this.kick}>
              <label>
                <input className="inputStyle" id="kickAddress" onChange={this.changeFieldById}
                  value={this.state.kickAddress} placeholder="Address" />
              </label>
              <button className="buttonStyle" type="submit"> Kick </button>
            </form>
            <h2>Kicked Signers</h2>
            Here is a list of the kicked signers:
            {kickedSigners}
          </div>
        </div>
      </div>
    );
  }
  updateProspectivesList () {
    let proposedSignerAccounts;
    let proposedSignerNames;
    let proposedSignerVotes;
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        Promise.resolve(agentRegistry.getNumProspectives()).then(numPros => {
          let prosPromises = [];
          for(let i = 0; i < numPros.toNumber(); i++) {
            prosPromises.push(agentRegistry.getProspective(i));
          }
          return prosPromises;
        }).spread((...prosac) => {
          proposedSignerAccounts = prosac;
          return proposedSignerAccounts.map(x => agentRegistry.getAgentName(x));
        }).spread((...pronam) => {
          proposedSignerNames = pronam;
          return Ethereum.getAccounts();
        }).then(accounts => {
          return proposedSignerAccounts.map(x => agentRegistry.getVoteInfo(x, accounts[0]));
        }).spread((...voat) => {
          proposedSignerVotes = voat;

          this.setState({
            proposedSignerAccounts,
            proposedSignerNames,
            proposedSignerVotes,
          });
        });
      });
  }
  updateKickedList () {
    let kickedSignerAccounts;
    let kickedSignerNames;
    let kickedSignerVotes;
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        Promise.resolve(agentRegistry.getNumKicked()).then(numKick => {
          let kickedPromises = [];
          for(let i = 0; i < numKick.toNumber(); i++) {
            kickedPromises.push(agentRegistry.getKicked(i));
          }
          return kickedPromises;
        }).spread((...kickacc) => {
          kickedSignerAccounts = kickacc;
          return kickedSignerAccounts.map(x => agentRegistry.getAgentName(x));
        }).spread((...kicknam) => {
          kickedSignerNames = kicknam;
          return Ethereum.getAccounts();
        }).then(accounts => {
          return kickedSignerAccounts.map(x => agentRegistry.getVoteInfo(x, accounts[0]));
        }).spread((...voat) => {
          kickedSignerVotes = voat;
          this.setState({
            kickedSignerAccounts,
            kickedSignerNames,
            kickedSignerVotes,
          });
        });
      });
  }
  updateSignersList () {
    let signerAccounts;
    let signerNames;
    Ethereum.getAgentRegistry()
      .then(reg => reg.deployed())
      .then(agentRegistry => {
        Promise.resolve(agentRegistry.getNumSigners()).then(numSigners => {
          let signerPromises = [];
          for(let i = 0; i < numSigners.toNumber(); i++) {
            signerPromises.push(agentRegistry.getSigner(i));
          }
          return signerPromises;
        }).spread((...sigac) => {
          signerAccounts = sigac;
          return signerAccounts.map(x => agentRegistry.getAgentName(x));
        }).spread((...signam) => {
          signerNames = signam;

          this.setState({
            signerAccounts,
            signerNames,
          });
        });
      });
  }

  updateprimaryAccount () {
    Ethereum.getAccounts()
      .then(accounts => {
        this.setState({primaryAccount: accounts[0]});
      });
  }

  componentDidMount () {
    this.updateSignersList();
    this.updateProspectivesList();
    this.updateKickedList();
    this.updateprimaryAccount();
  }
}

export default Home;
