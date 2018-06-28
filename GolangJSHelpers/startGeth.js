const Manager = require('ethereum-client-binaries').Manager;
const config = require('./clientBinaries.json');
const mgr =  new Manager(config);
const { spawn, spawnSync } = require('child_process'); var op = require('openport');
let home;
if(process.env.APPDATA) {
  home = process.env.APPDATA + '/MedRec';
}else {
  home = process.env.HOME;
  home += process.platform == 'darwin' ? '/Library/Preferences' : '/.medrec';
}

let runGeth = (binaryPath) => {
  op.find(
    {
      startingPort: 6338,
      endingPort: 10000,
    },
    function (err, port) {
      if(err) {
        console.error('Could not find open port for the Ethereum client');
        process.exit(1);
      }
      let bootnode = 'enode://757b0f2c9e2bede3a2554985880fae15bdf6ce90b375ec9cb025a397499b75f7994ee05edbdaddaf469edb72198220f2a32c216925ca50a5626f497de889d652@34.239.63.141:6338';
      let instance = spawn(binaryPath, [
        '--datadir=' + home,
        '--bootnodes=' + bootnode,
        '--rpc',
        '--rpcapi=eth,miner,personal,net,web3,clique',
        '--rpccorsdomain=*', //TODO: fix this security vulnerability and specify exact domains
        '--ws',
        '--wsapi=eth,miner,personal,net,web3,clique',
        '--wsorigins=*', //TODO: fix this security vulnerability and specify exact domains
        '--port=' + port, '--networkid=633732',
        '--targetgaslimit=10000000000000',
      ]);

      instance.stderr.on('data', (data) => {
        process.stdout.write(`Ethereum: ${data}`);
      });
    }
  );

};

let initGeth = (binaryPath) => {
  spawnSync(binaryPath, ['--datadir', home, 'init', 'EthereumClient/medrec-genesis.json']);
  runGeth(binaryPath);
};

mgr.init({
  folders: [
    home + '/Geth/',
  ],
})
  .then(() => {
    if(mgr.clients.Geth.state.available) {
      initGeth(mgr.clients.Geth.activeCli.fullPath);
    }else {
      mgr.download('Geth', {
        downloadFolder: home,
      }).then(file => initGeth(file.unpackFolder + '/geth'));
    }
  })
  .catch(err => {
    console.log(err);
    process.exit(1);
  });
