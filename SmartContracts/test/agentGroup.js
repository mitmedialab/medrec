var AgentGroup = artifacts.require('./AgentGroup.sol');

contract('AgentGroup', function (accounts) {
  let agentGroup;
  var constants = require('./constants.js')(accounts);

  before(() => {
    return AgentGroup.new({from: constants.agent1}).then(_agentGroup => {
      agentGroup = _agentGroup;
    });
  });

  it('should have an owner', function () {
    return agentGroup.agents(0).then(agent => {
      assert.equal(agent, constants.agent1);
    });
  });

  it('should not let the agent remove the only agent', function () {
    return agentGroup.removeAgent(constants.agent1, false, {from: constants.agent1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should let the agent add an agent', function () {
    return agentGroup.addAgent(constants.agent2, {from: constants.agent1});
  });

  it('should have the right number of agents', function () {
    return agentGroup.getNumAgents().then(agents => {
      assert.equal(agents, 2);
    });
  });

  it('should have teh correct owners', function () {
    return Promise.all([
      agentGroup.agents(0).then(agent => {
        assert.equal(agent, constants.agent1);
      }),
      agentGroup.agents(1).then(agent => {
        assert.equal(agent, constants.agent2);
      }),
    ]);
  });

  it('should not let other people add agents', function () {
    return agentGroup.addAgent(constants.family1, {from: constants.family1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should let the second agent add an agent', function () {
    return agentGroup.addAgent(constants.family1, {from: constants.agent1});
  });

  it('should have the right number of agents', function () {
    return agentGroup.getNumAgents().then(agents => {
      assert.equal(agents, 3);
    });
  });

  it('should let agents remove an agent', function () {
    return agentGroup.removeAgent(constants.agent1, {from: constants.family1});
  });

  it('should have the right number of owners', function () {
    return agentGroup.getNumAgents().then(agents => {
      assert.equal(agents, 2);
    });
  });
});
