const agentJson = require('../SmartContracts/build/contracts/Agent.json');
const Web3 = require('web3');
const contract = require('truffle-contract');

let Agent = contract(agentJson);
Agent.setProvider(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));
let providerAcc = process.argv[2] || '0x0';
let agent;
Agent.new()
  .then(_agent => {
    agent = _agent;
    agent.setOwner(providerAcc);
  })
  .then(() => {
    process.stdout.write(agent);
  });
