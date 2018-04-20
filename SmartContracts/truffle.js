module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 8545,
      network_id: '*', //Match any network id
    },
    production: {
      host: 'localhost',
      port: 8545,
      network_id: 633732, //matches the MedRec PoA blockchain
    },
  },
};
