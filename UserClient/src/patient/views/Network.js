import React, {Component} from 'react';
import Ethereum from '../../Ethereum';

import {InteractiveForceGraph, ForceGraphNode, ForceGraphLink} from 'react-vis-force';
import './chart.css';

class Network extends Component {
  constructor () {
    super();
    this.state = {
      nodes: [],
      links: [],
    };
    this.getProviders = this.getProviders.bind(this);
    this.getViewers = this.getViewers.bind(this);
  }
  render () {
    let colorProvider = 'rgb(186,203,204)';
    let colorViewer = 'rgb(207, 179, 211)';
    let radiusProvider = 10;
    let radiusViewer = 7;
    return (
      <div className="mainPanel">
        <h2>Network overview</h2>
        <p>You currently have {this.state.nodes.length - this.state.links.length} Provider <span className="colorBlock" style={{background: colorProvider}}></span> relationships and {this.state.links.length} Viewer <span className="colorBlock" style={{background: colorViewer}}></span> relationships.</p>
        <InteractiveForceGraph
          simulationOptions={{ animate: true, radiusMargin: 10, strength: { collide: 10 }, height: 400, width: 600 }}
          labelAttr="label"
          zoom
          showLabels
          highlightDependencies
          zoomOptions={{minScale: 1, maxScale: 5}} >
          {this.state.nodes.map((node) => <ForceGraphNode key={node.id} node={{ radius: node.group === 'provider' ? radiusProvider : radiusViewer, id: node.id, label: node.id}} fill={node.group === 'provider' ? colorProvider : colorViewer}  />)}
          {this.state.links.map((link) => <ForceGraphLink key={`${link.source}=>${link.target}`} link={{ source: link.source, target: link.target, strokeWidth: 2}} />)}
        </InteractiveForceGraph>
      </div>
    );
  }
  componentDidMount () {
    this.getProviders().then( () => {
      for(let i = 0; i < this.state.nodes.length; i ++) {
        this.getViewers(this.state.nodes[i].account, this.state.nodes[i].id);
      }
    });
  }
  getProviders () {
    let accounts;
    let agent;
    let relationGenerator;
    let providerAccounts;
    let relationships = [];
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
        for(let i = 0; i < providerNames.length; i++) {
          this.setState({
            nodes: [...this.state.nodes, { id: providerNames[i], group: 'provider', account: providerAccounts[i] }],
          });
        }
      });
  }
  getViewers (providerAccount, providerName) {
    if(providerAccount == undefined) {
      return;
    }
    let viewerAccounts;
    let relationship;
    Ethereum.getRelationship()
      .then(relationGenerator => {
        relationship = relationGenerator.at(providerAccount);
        return relationship.getNumViewers();
      })
      .then(numViewers => {
        let viewers = [];
        for(let i = 0; i < numViewers.toNumber(); i++) {
          viewers.push(relationship.viewers(i));
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
        for(let i = 0; i < viewerNames.length; i++) {
          this.setState({
            nodes: [...this.state.nodes, { id: viewerNames[i], group: 'viewer', account: viewerAccounts[i] }],
            links: [...this.state.links, { source: viewerNames[i], target: providerName, value: 1}],
          });
        }
      });
  }
}
export default Network;
