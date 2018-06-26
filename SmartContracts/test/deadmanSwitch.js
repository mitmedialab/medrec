var DeadmanSwitch = artifacts.require('./DeadmanSwitch.sol');

contract('DeadmanSwitch', function (accounts) {
  let agent;
  var constants = require('./constants.js')(accounts);

  before(() => {
    return DeadmanSwitch.new({from: constants.agent1}).then(_agent => {
      agent = _agent;
    });
  });

  it('should have an owner', function () {
    return agent.agent().then(agent => {
      assert.equal(agent, constants.agent1);
    });
  });

  it('let the owner add relationships', function () {
    return agent.addRelationship(constants.provider1, {from: constants.agent1});
  });

  it('should transfer the agent', function () {
    return agent.setAgent(constants.agent2, {from: constants.agent1});
  });

  it('should touch the switch', function () {
    return agent.touch({from: constants.agent2});
  });

  it('should set the timeout to 5 seconds', function () {
    return agent.setTimeout(5000, {from: constants.agent2});
  });

  it('should have the right timeout time', function () {
    return  agent.timeout().then(time => {
      assert.equal(time, 5000);
    });
  });

  it('should stil have an alive agent', function () {
    return agent.isAlive().then(alive => {
      assert.isTrue(alive);
    });
  });

  it('should be dead after 5 seconds', function (done) {
    agent.touch({from: constants.agent2}).then(() => {
      web3.currentProvider.sendAsync({
        jsonrpc: '2.0',
        method: 'evm_increaseTime',
        params: [6000],
        id: new Date().getTime(),
      }, (err, res) => {
        if(err !== null)throw err;

        agent.isAlive().then(alive => {
          assert.isFalse(alive);
          done();
        });
      });
    });
  });
});
