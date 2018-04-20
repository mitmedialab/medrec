var Agent = artifacts.require('Agent');
var AgentRegistry = artifacts.require('AgentRegistry');

module.exports = function (deployer) {
  deployer.deploy(Agent).then(a => {
    return deployer.deploy(AgentRegistry, 'TheFirstProvider', Agent.address, '127.0.0.1');
  });
};
