import React, { Component } from 'react';
import {connect} from 'react-redux';
import './dropdownmenu.css';
import {CLEAR_USER, CLEAR_RELATIONSHIP} from '../constants';

class DropDownMenu extends Component {
  constructor (props) {
    super(props);

    this.state = {
      dropped: false,
    };

    this.toggleDropped = this.toggleDropped.bind(this);
    this.signout = this.signout.bind(this);
  }

  toggleDropped (event) {
    this.setState({dropped: event.target.checked});
  }

  signout () {
    this.props.dispatch({type: CLEAR_USER});
    this.props.dispatch({type: CLEAR_RELATIONSHIP});
    this.props.history.push('/');
  }

  render () {
    return (
      <div className="signout">
        <button className="buttonStyle" onClick={this.signout} href="#">Sign out</button>
      </div>
    );
  }
}

export default connect()(DropDownMenu);
