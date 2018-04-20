var Agent = artifacts.require('./Agent.sol');
var AgentRegistry = artifacts.require('./AgentRegistry.sol');

contract('AgentRegistry', function (accounts) {
  let agentRegistry;
  let agent;
  var constants = require('./constants.js')(accounts);

  before(() => {
    return Agent.new({from: constants.provider1}).then(_agent => {
      agent = _agent;
      return AgentRegistry.new(constants.providerName1, agent.address, constants.host1, {from: constants.provider1});
    }).then(_agentRegistry => {
      agentRegistry = _agentRegistry;
    });
  });

  it('should have a signer', () => {
    return agentRegistry.getSigner(0).then(signer => {
      assert.equal(signer, constants.provider1);
    });
  });

  it('should have only one singer', () => {
    return agentRegistry.getNumSigners().then(value => {
      assert.equal(value, 1);
    });
  });

  it('should store the signer\'s name', () => {
    return agentRegistry.getAgentName(constants.provider1).then(name => {
      assert.equal(name, constants.providerName1);
    });
  });

  it('should store the signer\'s agent contract address', () => {
    return agentRegistry.getAgentContractAddr(constants.provider1).then(value => {
      assert.equal(value, agent.address);
    });
  });

  it('should reject proposals for an existing name', () => {
    return agentRegistry.propose(constants.providerName1, {from: constants.provider2})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should allow someone to propose themself', () => {
    return agentRegistry.propose(constants.providerName4, {from: constants.provider2});
  });

  it('should store the prospective\'s name', () => {
    return agentRegistry.getAgentName(constants.provider2).then(name => {
      assert.equal(name, constants.providerName4);
    });
  });

  it('shouldn\'t allow prospectives to readd themself', () => {
    return agentRegistry.propose({from: constants.provider2})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('shouldn\'t allow signers to repropose themself', () => {
    return agentRegistry.propose(constants.providerName4, {from: constants.provider1})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('shouldn\'t allow someone else to rescind the proposal', () => {
    return agentRegistry.rescind(constants.provider2, {from: constants.provider1})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should set the prospective contract address', () => {
    return agentRegistry.setAgentContractAddr(constants.provider2,
      {from: constants.provider2}
    );
  });

  it('should store the prospective\'s contract address', () => {
    return agentRegistry.getAgentContractAddr(constants.provider2).then(value => {
      assert.equal(value, constants.provider2);
    });
  });

  it('should reject votes from non signers', () => {
    return agentRegistry.vote(constants.provider2, true, {from: constants.provider2})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should accept the signers vote', () => {
    return agentRegistry.vote(constants.provider2, true, {from: constants.provider1})
      .then(result => {
        assert.equal(result.logs.length, 1);
        assert.equal(result.logs[0].event, 'AddSigner');
        assert.equal(result.logs[0].args.addr, constants.provider2);
      });
  });

  it('should reject further votes on a completed proposal', () => {
    return agentRegistry.vote(constants.provider2, true, {from: constants.provider1})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should reject name changs to an existing name', () => {
    return agentRegistry.setAgentName(constants.providerName1, {from: constants.provider2})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should change the new signer\'s name', () => {
    return agentRegistry.setAgentName(constants.providerName2,
      {from: constants.provider2}
    );
  });

  it('should store the signers\'s name', () => {
    return agentRegistry.getAgentName(constants.provider2).then(name => {
      assert.equal(name, constants.providerName2);
    });
  });

  it('should set the new signer\'s name', () => {
    return agentRegistry.setAgentHost(constants.host2,
      {from: constants.provider2}
    );
  });

  it('should store the signers\'s host', () => {
    return agentRegistry.getAgentHost(constants.provider2).then(name => {
      assert.equal(name, constants.host2);
    });
  });

  it('should reverse lookup the signer by name', () => {
    return agentRegistry.getAgentByName(constants.providerName2).then(name => {
      assert.equal(name, constants.provider2);
    });
  });

  it('should set the prospective contractAddr', () => {
    return agentRegistry.setAgentContractAddr(constants.provider4,
      {from: constants.provider4}
    );
  });

  it('should allow someone to propose themself', () => {
    //truffle add Promises to the global namespace
    return Promise.all([
      agentRegistry.propose('', {from: constants.provider3}),
      agentRegistry.propose('', {from: constants.provider4}),
      agentRegistry.propose('', {from: constants.pharmacy1}),
    ]);
  });

  it('should properly store info about the pending prospectives', () => {
    return Promise.all([
      agentRegistry.getNumProspectives().then(value => {
        assert.equal(value.toNumber(), 3);
      }),
      agentRegistry.getProspective(0).then(value => {
        assert.equal(value, constants.provider3);
      }),
      agentRegistry.getProposer(constants.provider3).then(value => {
        assert.equal(value, constants.provider3);
      }),
    ]);
  });

  //this signer should also get automatically confirmed by 1/2 of the signers
  it('should accept the first signers vote', () => {
    return agentRegistry.vote(constants.provider3, true, {from: constants.provider1})
      .then(result => {
        assert.equal(result.logs.length, 1);
        assert.equal(result.logs[0].event, 'AddSigner');
        assert.equal(result.logs[0].args.addr, constants.provider3);
      });
  });

  it('should properly store info about the pending transactions', () => {
    return Promise.all([
      agentRegistry.getNumProspectives().then(value => {
        assert.equal(value.toNumber(), 2);
      }),
      agentRegistry.getProspective(1).then(value => {
        assert.equal(value, constants.pharmacy1);
      }),
      agentRegistry.getProposer(constants.pharmacy1).then(value => {
        assert.equal(value, constants.pharmacy1);
      }),
    ]);
  });

  //these signers won't get automatically confirmed, they need 2/3 confirmations
  it('should accept the first signer\'s vote', () => {
    return Promise.all([
      agentRegistry.vote(constants.provider4, true, {from: constants.provider1}),
      agentRegistry.vote(constants.pharmacy1, true, {from: constants.provider1}),
    ]);
  });

  it('should allow someone rescind their proposal', () => {
    return agentRegistry.rescind(constants.pharmacy1, {from: constants.pharmacy1});
  });

  it('should properly store info about the pending transactions', () => {
    return Promise.all([
      agentRegistry.getNumProspectives().then(value => {
        assert.equal(value.toNumber(), 1);
      }),
      agentRegistry.getProspective(0).then(value => {
        assert.equal(value, constants.provider4);
      }),
      agentRegistry.getProposer(constants.provider4).then(value => {
        assert.equal(value, constants.provider4);
      }),
    ]);
  });

  it('should accept the first signer change their vote', () => {
    return agentRegistry.vote(constants.provider4, false, {from: constants.provider1});
  });

  it('should store the right number of voters', () => {
    return agentRegistry.getNumVoters(constants.provider4).then(value => {
      assert.equal(value.toNumber(), 1);
    });
  });

  it('should store the first voter in the list of voters', () => {
    return agentRegistry.getVoter(constants.provider4, 0).then(value => {
      assert.equal(value, constants.provider1);
    });
  });

  it('should store a no vote for the first signer', () => {
    return agentRegistry.getVoteInfo(constants.provider4, constants.provider1).then(value => {
      assert.isFalse(value);
    });
  });

  it('should store no yay vote', () => {
    return agentRegistry.getNumYayVotes(constants.provider4).then(value => {
      assert.equal(value.toNumber(), 0);
    });
  });

  it('should accept both signers\' vote', () => {
    return Promise.all([
      agentRegistry.vote(constants.provider4, true, {from: constants.provider1}),
      agentRegistry.vote(constants.provider4, true, {from: constants.provider2}).then(result => {
        assert.equal(result.logs.length, 1);
        assert.equal(result.logs[0].event, 'AddSigner');
        assert.equal(result.logs[0].args.addr, constants.provider4);
      }),
    ]);
  });

  it('should store no yay vote', () => {
    return agentRegistry.getNumYayVotes(constants.provider4).then(value => {
      assert.equal(value.toNumber(), 0);
    });
  });

  it('should register the new signer', () => {
    Promise.all([
      agentRegistry.setAgentName(constants.providerName4,
        {from: constants.provider4}
      ),
      agentRegistry.setAgentContractAddr(constants.providerName4,
        {from: constants.provider4}
      ),
      agentRegistry.setAgentHost(constants.host1,
        {from: constants.provider4}
      ),
    ]);
  });

  it('should properly update the signer info', () => {
    return Promise.all([
      agentRegistry.getNumSigners().then(value => {
        assert.equal(value.toNumber(), 4);
      }),
      agentRegistry.getAgentName(constants.provider4).then(name => {
        assert.equal(name, constants.providerName4);
      }),
      agentRegistry.getAgentHost(constants.provider4).then(name => {
        assert.equal(name, constants.host1);
      }),
      agentRegistry.getAgentContractAddr(constants.provider4).then(value => {
        assert.equal(value, constants.provider4);
      }),
      agentRegistry.getSigner(3).then(value => {
        assert.equal(value, constants.provider4);
      }),
      agentRegistry.isSigner.call(constants.provider2).then(value => {
        assert.isTrue(value);
      }),
      agentRegistry.getNumProspectives().then(value => {
        assert.equal(value.toNumber(), 0);
      }),
    ]);
  });

  it('shouldn\'t allow randos to kick people', () => {
    return agentRegistry.kick(constants.provider1, {from: constants.pharmacy1})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should allow a signer to kick someone', () => {
    return agentRegistry.kick(constants.provider1, {from: constants.provider2});
  });

  it('shouldn\'t allow duplicate kicks', () => {
    return agentRegistry.kick(constants.provider1, {from: constants.provider3})
      .then(() => {assert(false);}, () => {
        assert(true);
      });
  });

  it('should properly store info about the pending transactions', () => {
    return Promise.all([
      agentRegistry.getNumKicked().then(value => {
        assert.equal(value.toNumber(), 1);
      }),
      agentRegistry.getKicked(0).then(value => {
        assert.equal(value, constants.provider1);
      }),
      agentRegistry.getProposer(constants.provider1).then(value => {
        assert.equal(value, constants.provider2);
      }),
    ]);
  });

  it('should allow the other signers to vote', () => {
    return agentRegistry.vote(constants.provider1, true, {from: constants.provider3})
      .then(result => {
        assert.equal(result.logs.length, 1);
        assert.equal(result.logs[0].event, 'RemoveSigner');
        assert.equal(result.logs[0].args.addr, constants.provider1);
      });
  });

  it('should properly update the signer info', () => {
    return Promise.all([
      agentRegistry.getNumSigners().then(value => {
        assert.equal(value.toNumber(), 3);
      }),
      agentRegistry.getSigner(2).then(value => {
        assert.equal(value, constants.provider3);
      }),
      agentRegistry.isSigner.call(constants.provider1).then(value => {
        assert.isFalse(value);
      }),
      agentRegistry.getNumProspectives().then(value => {
        assert.equal(value.toNumber(), 0);
      }),
    ]);
  });
});
