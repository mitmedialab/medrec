var Agent = artifacts.require('./Agent.sol');

contract('Agent', function (accounts) {
  let agent;
  var constants = require('./constants.js')(accounts);

  before(() => {
    return Agent.new({from: constants.agent1}).then(_agent => {
      agent = _agent;
    });
  });

  it('should have an owner', function () {
    return agent.agent().then(agent => {
      assert.equal(agent, constants.agent1);
    });
  });

  it('should enable the agent', function () {
    return agent.agentEnabled().then(enable => {
      assert.equal(enable, true);
    });
  });

  it('should have the right number of owners', function () {
    return agent.getNumEnabledOwners().then(owners => {
      assert.equal(owners, 1);
    });
  });

  it('let the owner add relationships', function () {
    return agent.addRelationship(constants.provider1, {from: constants.agent1});
  });

  it('should transfer the agent', function () {
    return agent.setAgent(constants.agent2, {from: constants.agent1});
  });

  it('should not let other people add relationships', function () {
    return agent.addRelationship(constants.provider1, {from: constants.agent1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should let the agent add a relationship', function () {
    return agent.addRelationship(constants.pharmacy1, {from: constants.agent2});
  });

  it('should not let the agent remove the only owner', function () {
    return agent.enableAgent(false, {from: constants.agent1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should let the agent add a custodian', function () {
    return agent.addCustodian(constants.family1, {from: constants.agent2});
  });

  it('should have the right number of owners', function () {
    return agent.getNumEnabledOwners().then(owners => {
      assert.equal(owners, 2);
    });
  });

  it('should not let other people add custodians', function () {
    return agent.addCustodian(constants.family1, {from: constants.agent1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should disable the agent', function () {
    return agent.enableAgent(false, {from: constants.family1}).then(() => {
      return agent.agentEnabled();
    }).then((enable) => {
      assert.equal(enable, false);
    });
  });

  it('should not let the custodian disable the only owner', function () {
    return agent.enableCustodian(constants.family1, false, {from: constants.family1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should not let the custodian remove the only owner', function () {
    return agent.removeCustodian(constants.family1, false, {from: constants.family1})
      .then(() => {assert(false);}, (test, err) => {
        assert(true);
      });
  });

  it('should let the custodian add a custodian', function () {
    return agent.addCustodian(constants.family2, {from: constants.family1});
  });

  it('should have the right number of owners', function () {
    return agent.getNumEnabledOwners().then(owners => {
      assert.equal(owners, 2);
    });
  });

  it('should let the custodian add a relationship', function () {
    return agent.addRelationship(constants.provider2, {from: constants.family1});
  });

  it('should let the custodian disable a custodian', function () {
    return agent.enableCustodian(constants.family1, false, {from: constants.family1});
  });

  it('should have the right number of owners', function () {
    return agent.getNumEnabledOwners().then(owners => {
      assert.equal(owners, 1);
    });
  });

  it('should have the right number of relationships', function () {
    return agent.getNumRelationships().then(relations => {
      assert.equal(relations, 3);
    });
  });
});
