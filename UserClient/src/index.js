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

//registerServiceWorker();
console.log('hello world');
class InitialLoad extends Component {
  render () {
    return this.props.children;
  }
  componentDidMount () {
    if(store.getState().homeReducer.username) {
      Ethereum.refreshVault();
    }
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
