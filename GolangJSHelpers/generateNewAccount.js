const Web3  = require('web3');
const Promise  = require('bluebird');
const {keystore, signing} = require('eth-lightwallet');
const fs = require('fs');

let path = process.argv[2] || 'path';
let password = process.argv[3] || '';

//TODO there is a potential race condition if this script is called in rapid succession
//TODO this should be fixed by making a single node server handle requests
let serializedKeystore = fs.readFileSync(path, {flag: 'r'});
let vault = keystore.deserialize(serializedKeystore);

vault.keyFromPassword(password, (_, pwDerivedKey) => {
  //generate new address/private key pair
  vault.generateNewAddress(pwDerivedKey, 1);
  let addresses = vault.getAddresses();
  fs.writeFileSync(path, vault.serialize(), {mode: 0o666});
  process.stdout.write(addresses[addresses.length - 1]);
});
