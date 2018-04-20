import React, { Component } from 'react';
import RPCClient from '../../RPCClient';
import Ethereum from '../../Ethereum';
import './home.css';

class Tests extends Component {
  constructor (props) {
    super(props);
    this.state = {
      patientID: '',
      documents: [],
      username: '',
      password: '',
      tableState: 'isHidden',
      providerHost: '',
      providerAddress: '',
    };
    this.fetchDocument = this.fetchDocument.bind(this);
    this.signMessage = this.signMessage.bind(this);
  }

  //fetch documents
  fetchDocument () {
    RPCClient.remote('127.0.0.1').send('MedRecRemote.PatientDocuments', {
      PatientID: 1,
    }).then( (res) => {
      if(res.Error !== '') {
        throw (res.Error);
      }else {
        this.setState({documents: res.Documents});
        this.setState({tableState: 'isVisible'});
      }
    });
  }

  signMessage () {
    Ethereum.getAccounts().then((Accounts) => {
      Ethereum.web3.eth.sign('message', Accounts[0]).then((sig) => {
        RPCClient.remote('127.0.0.1').send('MedRecRemote.Recover', {
          MsgHex: sig.msgHex,
          Signature: sig.signedMsgHex,
        }).then( (res) => {
          if(res.Error !== '') {
            throw (res.Error);
          }else {
            this.setState({address: res.Address});
          }
        });
      });
    });
  }


  render () {
    return  (
      <div>
        <button
          onClick={this.fetchDocument}
          className="buttonStyle"
        > Fetch records</button>
        <br/>
        <br/>
        <table className="table-striped">
          <thead>
            <tr>
              <th>DocumentID</th>
              <th>docdatetime</th>
              <th>patientid</th>
              <th>practiceid</th>
              <th>recvddatetime</th>
            </tr>
          </thead>
          <tbody>
            {this.state.documents.map((d, i)=>
              <tr key={i}>
                <td>{d.DocumentID}</td>
                <td>{d.docdatetime}</td>
                <td>{d.patientid}</td>
                <td>{d.practiceid}</td>
                <td>{d.recvddatetime}</td>
              </tr>)}
          </tbody>
        </table>
      </div>
    );
  }
}

export default Tests;
