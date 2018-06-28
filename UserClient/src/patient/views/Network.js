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
    let nodes = this.state.nodes.map((node) => <ForceGraphNode
      key={node.id}
      node={{
        radius: node.group === 'provider' ? radiusProvider : radiusViewer,
        id: node.id,
        label: node.id,
      }} fill={node.group === 'provider' ? colorProvider : colorViewer}  />
    );
    let links = this.state.links.map((link) => <ForceGraphLink
      key={`${link.source}=>${link.target}`}
      link={{ source: link.source, target: link.target, strokeWidth: 2}} />
    );
    return (
      <div className="mainPanel">
        <h2>Network overview</h2>
        <p>You currently have {this.state.nodes.length - this.state.links.length} Provider
          <span className="colorBlock" style={{background: colorProvider}}></span>
          relationships and {this.state.links.length} Viewer
          <span className="colorBlock" style={{background: colorViewer}}></span> relationships.</p>
        <InteractiveForceGraph
          simulationOptions={{
            animate: true,
            radiusMargin: 10,
            strength: { collide: 10 },
            height: 400,
            width: 600,
          }}
          labelAttr="label"
          zoom
          showLabels
          highlightDependencies
          zoomOptions={{minScale: 1, maxScale: 5}} >
          {nodes}
          {links}
        </InteractiveForceGraph>
      </div>
    );
  }
  componentDidMount () {
    this.getProviders().then( () => {
      for(let i = 0; i < this.state.nodes.length; i ++) {
        this.getViewers(this.state.nodes[i].relationship, this.state.nodes[i].id);
      }
    });
  }
  getProviders () {
    let accounts;
    let agent;
    let relationGenerator;
    let relationshipAccounts;
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
      .spread((..._rels) => {
        relationshipAccounts = _rels;
        return relationshipAccounts.map(account => {
          return relationGenerator.at(account).providerName();
        });
      })
      .spread((...providerNames) => {
        return providerNames.map(name => Ethereum.decrypt(accounts[0], name));
      })
      .spread((...providerNames) => {
        let nodes = this.state.nodes;
        for(let i = 0; i < providerNames.length; i++) {
          nodes.push({
            id: providerNames[i],
            group: 'provider',
            relationship: relationshipAccounts[i],
          });
        }
        this.setState({nodes});
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
        viewerAccounts = _viewAcc.filter(account => account.localeCompare('0x') !== 0);
        let names = viewerAccounts
          .map(account => {
            return relationship.getViewerName(account);
          });
        return names;
      }).spread((...viewerNames) => {
        let nodes = this.state.nodes;
        let links = this.state.links;
        for(let i = 0; i < viewerNames.length; i++) {
          nodes.push({
            id: viewerNames[i],
            group: 'viewer',
            account: viewerAccounts[i],
          });
          links.push({
            source: viewerNames[i],
            target: providerName,
            value: 1,
          });
        }
        this.setState({nodes, links});
      });
  }
}
export default Network;
