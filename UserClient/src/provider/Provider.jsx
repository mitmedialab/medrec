import React, { Component } from 'react';
import {Route, Link, Switch} from 'react-router-dom';
import RPCClient from '../RPCClient';
import {connect} from 'react-redux';

import Home from './views/Home';
import PatientList from './views/PatientList';
import DropDownMenu from './DropDownMenu';

class Provider extends Component {

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
      <div id="providerStyle">
        <DropDownMenu history={this.props.history}/>
        <div className="sidebar">
          <div className="medrec-logo-provider"></div>
          <p>First name <span className="status">{this.state.fname}</span></p>
          <p>Last name <span className="status">{this.state.lname}</span></p>
          <p>Contract <span className="status">{this.props.contract}</span></p>
          <div className="button-group">
            <Link to="/provider/home">
              <button className="buttonHome">Home
              </button>
            </Link>
            <Link to="/provider/patientList">
              <button className="buttonList">Patient List
              </button>
            </Link>
          </div>
        </div>
        <div className="pane">
          <Switch>
            <Route path="/provider/patientList" component={PatientList}/>
            <Route component={Home}/>
          </Switch>
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
})(Provider);
