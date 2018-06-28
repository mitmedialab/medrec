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

  it('should add a new viewer group', function () {
    return Promise.all([
      relationship.addViewerGroup({from: constants.patron1}),
      relationship.addViewerGroup({from: constants.patron1}),
    ]);
  });

  //Test adding a new viewer
  it('should add a new viewer', () => {
    return Promise.all([
      relationship.addViewer(constants.providerName2, 0, constants.provider2, constants.providerName1, {from: constants.patron1}),
      relationship.addViewer(constants.providerName3, 1, constants.provider3, constants.providerName1, {from: constants.patron1}),
    ]);
  });

  it('should fail for duplicate viewers', () => {
    return relationship.addViewer(constants.providerName2, 0, constants.provider2, constants.providerName1, {from: constants.patron1})
      .then(() => {assert(false);}, () => {assert(true);});
  });

  it('should have all the viewers', () => {
    return relationship.getNumViewerGroups().then(num => {
      assert.equal(num, 2);
    });
  });


  it('should add a new viewer group', function () {
    return relationship.addViewerGroup({from: constants.patron1});
  });

  //Test removing a viewer and adding a different one
  it('should add new viewers', () => {
    return Promise.all([
      relationship.addViewer(constants.familyName1, 0, constants.family1, constants.providerName1, {from: constants.patron1}),
      relationship.addViewer(constants.familyName2, 1, constants.family2, constants.providerName1, {from: constants.patron1}),
      relationship.addViewer(constants.providerName4, 2, constants.provider4, constants.providerName1, {from: constants.patron1}),
    ]);
  });

  it('should remove a viewer', () => {
    return relationship.removeViewer(0, constants.provider2, {from: constants.patron1});
  });

  it('should remove a viewer', () => {
    return relationship.removeViewerGroup(1, {from: constants.patron1});
  });

  it('should have the rightviewers', () => {
    return Promise.all([
      relationship.getNumViewerGroups().then(num => {
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
