const agentRegistryJson = require('../SmartContracts/build/contracts/AgentRegistry.json');
const Web3 = require('web3');
const contract = require('truffle-contract');

let AgentRegistry = contract(agentRegistryJson);
AgentRegistry.setProvider(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));
let providerAcc = process.argv[2] || '0x0';
AgentRegistry.deployed()
  .then(agentRegistry => agentRegistry.getAgentHost(providerAcc))
  .then(host => {
    console.log(host);
  });
