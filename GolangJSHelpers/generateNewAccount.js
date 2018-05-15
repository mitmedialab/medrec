import Web3 from 'web3';
import Promise from 'bluebird';
import {keystore, signing} from 'eth-lightwallet';
const fs = require('fs');

let password = process.argv[2] || '0x0';

let home;
if(process.env.APPDATA) {
  home = process.env.APPDATA + '/MedRec';
}else {
  home = process.env.HOME;
  home += process.platform == 'darwin' ? '/Library/Preferences' : '/.medrec';
}
let path = home + '/providerKeystore.key';

//TODO there is a potential race condition if this script is called in rapid succession
//TODO this should be fixed by making a single node server handle requests
let serializedKeystore = fs.readSync(path);
let vault = keystore.deserialize(serializedKeystore);

vault.keyFromPassword(password, (_, pwDerivedKey) => {
  //generate new address/private key pair
  vault.generateNewAddress(pwDerivedKey, 1);
  let address = vault.getAddresses()[0];
  fs.wrteSync(path, vault.serialize());
  process.stdout.write(address);
});
