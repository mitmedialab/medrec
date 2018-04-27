import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Link,
} from 'react-router-dom';



class SidePanel extends Component {

  render () {

    return  (
      <div className="pane-sm sidebar">
        <div className="medrec-logo"></div>
        <p>PATIENT INFO </p>
        <p>First Name</p>
        <p>Last Name</p>
        <p>DOB</p>
        <p>Contact</p>
        <p>PCP</p>
        <p>Hospital</p>
      </div>
    );
  }
}

export default SidePanel;
