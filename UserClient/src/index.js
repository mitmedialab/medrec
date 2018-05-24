import React, {Component} from 'react';
import { PersistGate } from 'redux-persist/es/integration/react';
import {Provider as ReduxProvider} from 'react-redux';
import {render as renderDOM} from 'react-dom';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import {store, persistor} from './reduxStore';
import Patient from './patient/Patient';
import Provider from './provider/Provider';
import Home from './home/Home';
import Ethereum from './Ethereum';

window.Ethereum = Ethereum;

class InitialLoad extends Component {
  constructor () {
    super();
    this.state = {
      ethLoading: true,
    };
  }

  render () {
    if(this.state.ethLoading) {
      return (<div></div>);
    }
    return this.props.children;

  }
  componentDidMount () {
    //if the user just refreshed the page get their data out of browser storage
    //and into the web app
    if(store.getState().homeReducer.username) {
      Ethereum.refreshVault();
    }
    Ethereum.waitForRPCConn().then(() => this.setState({ethLoading: false}));
  }
}
renderDOM(
  <PersistGate persistor={persistor}>
    <InitialLoad>
      <ReduxProvider store={store}>
        {/*<Header/>*/}
        <BrowserRouter>
          <Switch>
            <Route path="/patient" component={Patient}/>
            <Route path="/provider" component={Provider}/>
            <Route path="/" component={Home}/>
          </Switch>
        </BrowserRouter>
        {/*<Footer/>*/}
      </ReduxProvider>
    </InitialLoad>
  </PersistGate>,
  document.getElementById('root')
);
