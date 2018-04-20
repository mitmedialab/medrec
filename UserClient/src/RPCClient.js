import Promise from 'bluebird';

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
  send (method, ...args) {

    return new Promise((resolve, reject) => {
      let id = requestNonce++;

      var headers = new Headers();
      headers.append('Content-Type', 'application/json');

      fetch(this.rpcAddress, {
        method: 'POST',
        body: JSON.stringify({'method': method, 'params': args, 'id': id}),
        headers: headers,
      }).then(res => res.json())
        .then((data) => {
          if(data.error !== null) {
            reject(data.error);
          }else {
            resolve(data.result);
          }
        });
    });
  }
}

let client = new RPCClient();

export default client;
