import React, { Component } from 'react';
import {connect} from 'react-redux';
import RPCClient from '../../RPCClient';

class PatientList extends Component {
  constructor () {
    super();
    this.state = {
      uid: '',
      account: '',
    };

    this.addPatient = this.addPatient.bind(this);
    this.valueChanged = this.valueChanged.bind(this);
  }
  addPatient (event) {
    event.preventDefault();
    RPCClient.send('MedRecLocal.AddAccount', {
      UniqueID: this.state.uid,
      Account: this.state.account,
      Username: this.props.username,
      Password: this.props.password,
    }).then(() => {
      this.setState({
        uid: '',
        account: '',
      });
    });
  }
  valueChanged (event) {
    let state = this.state;
    state[event.target.id] = event.target.value;
    this.setState(state);
  }
  render () {
    return (
      <div className="mainPanel">
        <form>
          <h2>Add a patient</h2>
          <label>
            <span>Patient Account</span>
            <input id="account" value={this.state.account} onChange={this.valueChanged}></input>
          </label>
          <label>
            <span>Unique ID</span>
            <input id="uid" value={this.state.uid} onChange={this.valueChanged}/>
          </label>
          <input type="submit" onClick={this.addPatient}/>
        </form>
        <h2>List of patients</h2>
        <table className="table-striped">
          <thead>
            <tr>
              <th>First name</th>
              <th>Last name</th>
              <th>DOB</th>
              <th>Insurance</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Random</td>
              <td>Person</td>
              <td>3/8/1998</td>
              <td>An insurance</td>
            </tr>
            <tr>
              <td>Another</td>
              <td>Person</td>
              <td>3/8/1998</td>
              <td>Another insurance</td>
            </tr>
          </tbody>
        </table>
      </div>
    );
  }
}

export default connect(state => {
  return {
    username: state.homeReducer.username,
    password: state.homeReducer.password,
  };
})(PatientList);
