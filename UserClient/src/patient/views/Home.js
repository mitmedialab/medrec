import React, { Component } from 'react';
import Tests from './Tests';
import './home.css';

class Home extends Component {

  render () {
    return  (
      <div className="mainPanel">
        <h2>Your medical records overview</h2>
        <p>Here you can access your most recent medical records.</p>
        <Tests/>
      </div>
    );
  }
}

export default Home;
