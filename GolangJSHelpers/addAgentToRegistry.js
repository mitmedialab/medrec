const agentRegistryJson = require('../SmartContracts/build/contracts/AgentRegistry.json');
const Web3 = require('web3');
const contract = require('truffle-contract');

let Agent = contract(agentRegistryJson);
Agent.setProvider(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));
let contractAddr = process.argv[2] || '0x0';

AgentRegistry.deployed()
  .then(agentRegistry => agentRegistry.setAgentContractAddr(contractAddr))
  .then(() => {  });
