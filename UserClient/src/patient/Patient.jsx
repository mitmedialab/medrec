import React, { Component } from 'react';
import {Route, Link} from 'react-router-dom';
import RPCClient from '../RPCClient';
import {connect} from 'react-redux';

import Home from './views/Home';
import Network from './views/Network';
import Relationships from './views/Relationships';
import DropDownMenu from './DropDownMenu';

class Patient extends Component {

  constructor (props) {
    super(props);
    this.state = {
      fname: '',
      lname: '',
      username: '',
    };
  }

  componentDidMount () {
    RPCClient.send('MedRecLocal.GetUserDetails', {
      username: this.props.username,
      contract: this.props.contract,
    }).then( (res) => {
      this.setState({
        fname: res.FirstName,
        lname: res.LastName,
      });
    });
  }

  render () {
    return (
      <div id="patientStyle">
        <DropDownMenu history={this.props.history}/>
        <div className="sidebar">
          <div className="medrec-logo-patient"></div>
          <p>First name <span className="status">{this.state.fname}</span></p>
          <p>Last name <span className="status">{this.state.lname}</span></p>
          <p>Contract   <span className="status">{this.props.contract}</span></p>
          <div className="button-group">
            <Link to="/patient/home">
              <button className="buttonHome">Home
              </button>
            </Link>
            <Link to="/patient/network">
              <button className="buttonAbout">Your network
              </button>
            </Link>
            <Link to="/patient/relationships">
              <button className="buttonRelationships">Edit Relationships
              </button>
            </Link>
          </div>
        </div>
        <div>
          <Route path="/patient/home" component={Home}/>
          <Route path="/patient/network" component={Network}/>
          <Route path="/patient/relationships" render={(props) => (
            <Relationships contract={this.props.contract} {...props} />
          )} />
        </div>
      </div>
    );
  }
}

export default connect(state => {
  return {
    username: state.homeReducer.username,
    contract: state.homeReducer.contract,
  };
})(Patient);
