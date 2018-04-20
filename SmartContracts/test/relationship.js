var Relationship = artifacts.require('./Relationship.sol');

contract('Relationship', function (accounts) {
  let relationship;
  var constants = require('./constants.js')(accounts);

  before(() => {
    return Relationship.new(constants.provider1, {from: constants.patron1}).then(_relationship => {
      relationship = _relationship;
    });
  });

  it('should have an patron', function () {
    return relationship.patron().then(patron => {
      assert.equal(patron, constants.patron1);
    });
  });

  it('should have an provider', function () {
    return relationship.provider().then(provider => {
      assert.equal(provider, constants.provider1);
    });
  });

  it('should set the provider name', function () {
    return relationship.setProviderName(constants.providerName1);
  });

  it('should have the correct provider name', function () {
    return relationship.providerName().then(name => {
      assert.equal(name, constants.providerName1);
    });
  });

  it('should fail adding permissions for non existant viewers', function () {
    return relationship.addPermission(constants.provider2, constants.perm1, constants.canRead, constants.canWrite, constants.duration1)
      .then(() => {assert(false);}, () => {assert(true);});
  });

  it('should add a new viewer group', function () {
    return Promise.all([
      relationship.addViewerGroup({from: constants.patron1}),
      relationship.addViewerGroup({from: constants.patron1}),
    ]);
  });

  //Test adding a new viewer
  it('should add a new viewer', () => {
    return Promise.all([
      relationship.addViewer(constants.providerName2, 0, constants.provider2, {from: constants.patron1}),
      relationship.addViewer(constants.providerName3, 1, constants.provider3, {from: constants.patron1}),
    ]);
  });

  it('should fail for duplicate viewers', () => {
    return relationship.addViewer(constants.providerName2, 0, constants.provider2, {from: constants.patron1})
      .then(() => {assert(false);}, () => {assert(true);});
  });

  it('should have all the viewers', () => {
    return relationship.getNumViewers().then(num => {
      assert.equal(num, 2);
    });
  });

  //Test adding a new permission
  it('let the patron add a permission', function () {
    return Promise.all([
      relationship.addPermission(0, constants.perm1,
        constants.canRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }),
      relationship.addPermission(0, constants.perm2,
        constants.cannotRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }),
      relationship.addPermission(1, constants.perm1,
        constants.canRead, constants.cannotWrite, constants.duration1, {
          from: constants.patron1,
        }),
    ]);
  });

  it('should not let other people add permissions', function () {
    return relationship.addPermission(constants.provider2, constants.perm1,
      constants.canRead, constants.canWrite, constants.duration1, {
        from: constants.patron2,
      })
      .then(() => {assert(false);}, () => {assert(true);});
  });
  it('should have the right number of permissions', () => {
    return Promise.all([
      relationship.getNumPermissions(0).then(num => {
        assert.equal(num, 2);
      }),
      relationship.getNumPermissions(1).then(num => {
        assert.equal(num, 1);
      }),
      relationship.getNumPermissions(2)
        .then(() => {assert(false);}, () => {assert(true);}),
    ]);
  });

  it('should properly store the permissions', () => {
    return Promise.all([
      relationship.checkPermission(0, constants.permIndex1).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.canWrite);
      }),
      relationship.checkPermission(0, constants.permIndex2).then(perm => {
        assert.equal(perm[0], constants.cannotRead);
        assert.equal(perm[1], constants.canWrite);
      }),
      relationship.checkPermission(1, constants.permIndex1).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.cannotWrite);
      }),
      relationship.checkPermission(1, constants.permIndex2).then(perm => {
        assert.equal(perm[0], constants.cannotRead);
        assert.equal(perm[1], constants.cannotWrite);
      }),
      relationship.checkPermission(2, constants.permIndex1).then(perm => {
        assert.equal(perm[0], constants.cannotRead);
        assert.equal(perm[1], constants.cannotWrite);
      }),
    ]);
  });

  it('should add a new viewer group', function () {
    return relationship.addViewerGroup({from: constants.patron1});
  });

  //Test removing a viewer and adding a different one
  it('should add new viewers', () => {
    return Promise.all([
      relationship.addViewer(constants.familyName1, 0, constants.family1, {from: constants.patron1}),
      relationship.addViewer(constants.familyName2, 1, constants.family2, {from: constants.patron1}),
      relationship.addViewer(constants.providerName4, 2, constants.provider4, {from: constants.patron1}),
    ]);
  });
  it('should remove a viewer', () => {
    return relationship.removeViewer(0, constants.provider2, {from: constants.patron1});
  });
  it('should not allow permissions for non-existent viewer groups', function () {
    return Promise.all([
      relationship.addPermission(5, constants.perm1,
        constants.canRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }).then(() => {assert(false);}, () => {assert(true);}),
      relationship.setPermissionReadAbility(5, constants.perm2,
        constants.canRead, {from: constants.patron1})
        .then(() => {assert(false);}, () => {assert(true);}),
      relationship.setPermissionWriteAbility(5, constants.perm2,
        constants.cannotWrite, {from: constants.patron1})
        .then(() => {assert(false);}, () => {assert(true);}),
      relationship.setPermissionDuration(5, constants.perm2,
        constants.durationForever, {from: constants.patron1})
        .then(() => {assert(false);}, () => {assert(true);}),
    ]);
  });
  it('let the patron add a permission', function () {
    return Promise.all([
      relationship.addPermission(0, constants.perm1,
        constants.canRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }),
      relationship.addPermission(0, constants.perm2,
        constants.cannotRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }),
      relationship.addPermission(2, constants.perm2,
        constants.canRead, constants.canWrite, constants.duration1, {
          from: constants.patron1,
        }),
      relationship.addPermission(2, constants.perm1,
        constants.canRead, constants.cannotWrite, constants.duration1, {
          from: constants.patron1,
        }),
    ]);
  });
  it('should remove a viewer', () => {
    return relationship.removeViewerGroup(1, {from: constants.patron1});
  });
  it('should properly have the right number of permissions', () => {
    return Promise.all([
      relationship.getNumPermissions(0).then(num => {
        assert.equal(num, 2);
      }),
      relationship.getNumPermissions(1).then(num => {
        assert.equal(num, 2);
      }),
      relationship.getNumPermissions(2)
        .then(() => {assert(false);}, () => {assert(true);}),
    ]);
  });

  it('should properly store the permissions', () => {
    return Promise.all([
      relationship.checkPermission(0, constants.permIndex1).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.canWrite);
      }),
      relationship.checkPermission(0, constants.permIndex2).then(perm => {
        assert.equal(perm[0], constants.cannotRead);
        assert.equal(perm[1], constants.canWrite);
      }),
      relationship.checkPermission(1, constants.permIndex1).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.canWrite);
      }),
      relationship.checkPermission(1, constants.permIndex2).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.cannotWrite);
      }),
    ]);
  });

  it('should have the rightviewers', () => {
    return Promise.all([
      relationship.getNumViewers().then(num => {
        assert.equal(num, 2);
      }),
      relationship.getViewerName(constants.provider3).then(name => {
        assert.equal(name, constants.providerName3);
      }),
      relationship.getViewerName(constants.family1).then(name => {
        assert.equal(name, constants.familyName1);
      }),
    ]);
  });

  //Test changing permissions
  it('should set a permission\'s access parameters', () => {
    return Promise.all([
      relationship.setPermissionReadAbility(0, constants.permIndex2,
        constants.canRead, {from: constants.patron1}),
      relationship.setPermissionWriteAbility(0, constants.permIndex2,
        constants.cannotWrite, {from: constants.patron1}),
      relationship.setPermissionDuration(0, constants.permIndex2,
        constants.durationForever, {from: constants.patron1}),
    ]);
  });
  it('should properly store the permissions', () => {
    return Promise.all([
      relationship.checkPermission(0, constants.permIndex2).then(perm => {
        assert.equal(perm[0], constants.canRead);
        assert.equal(perm[1], constants.cannotWrite);
      }),
      relationship.getPermissionReadAbility(0, constants.permIndex2).then(perm => {
        assert.equal(perm, constants.canRead);
      }),
      relationship.getPermissionWriteAbility(0, constants.permIndex2).then(perm => {
        assert.equal(perm, constants.cannotWrite);
      }),
      relationship.getPermissionDuration(0, constants.permIndex2).then(perm => {
        assert.equal(perm, constants.durationForever);
      }),
    ]);
  });
  it('should invalidate a permission by setting the duration', () => {
    return relationship.setPermissionDuration(0, constants.permIndex2,
      constants.duration0, {from: constants.patron1}).then(() => {
      return relationship.checkPermission(0, constants.permIndex2);
    }).then(perm => {
      assert.equal(perm[0], constants.cannotRead);
      assert.equal(perm[1], constants.cannotWrite);
      return relationship.getPermissionDuration(0, constants.permIndex2);
    }).then(perm => {
      assert.equal(perm, constants.duration0);
    });
  });
  it('should prevent randos from terminating the contract', () => {
    return relationship.terminate({from: constants.provider2})
      .then(() => {assert(false);}, () => {assert(true);});
  });
  it('should terminate the contract', () => {
    return relationship.terminate({from: constants.provider1});
  });
  it('should be dead', () => {
    return relationship.patron()
      .then(() => {assert(false);}, () => {assert(true);});
  });
});
