import Promise from 'bluebird';
import Ethereum from './Ethereum';

let requestNonce = 0;

class RPCClient {
  constructor () {
    let hostname  =  window.location.hostname;
    this.host = hostname !== '' ? hostname : '127.0.0.1';
    this.port = 6337;
    this.rpcAddress = 'http://' + this.host + ':' + this.port + '/localRPC';
  }
  remote (host) {
    let client = new RPCClient();
    client.host = host;
    client.rpcAddress = 'http://' + host + ':' + this.port + '/remoteRPC';
    return client;
  }
  send (method, args, signed = true) {
    args = args || {};
    return new Promise((resolve, reject) => {

      if(signed) {
      //sign the message with the current time before sending
        args.Time = '' + (new Date).getTime();
        Ethereum.getAccounts()
          .then(accounts => {
            return  Ethereum.web3.eth.sign(args.Time, accounts[0]);
          }).then(sig => {
            args.Signature = sig;
            this._send(method, args, resolve, reject);
          });
      }else {
        this._send(method, args, resolve, reject);
      }
    });
  }
  _send (method, args, resolve, reject) {
    let id = requestNonce++;
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    fetch(this.rpcAddress, {
      method: 'POST',
      body: JSON.stringify({'method': method, 'params': [args], 'id': id}),
      headers: headers,
    }).then(res => res.json())
      .then((data) => {
        if(data.error !== null) {
          reject(data.error);
        }else {
          resolve(data.result);
        }
      });
  }
}

let client = new RPCClient();

export default client;
