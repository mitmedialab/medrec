pragma solidity ^0.4.15;

//Represents both a Patient and a Provider,
//Patients list of relationships with different Providers
contract Agent {
  address public agent;
  bool public agentEnabled;
  address[] public custodians;
  bool[] public custodianEnabled;
  address[] public relationships; //list of Relationship contract addresses

  modifier isOwner() {
    bool enable;
    if(agentEnabled && msg.sender == agent) enable = true;
    for(uint i = 0; i < custodians.length; i++) {
      if(custodianEnabled[i] && msg.sender == custodians[i]) {
        enable = true;
        break;
      }
    }
    if(!enable) revert();
    _;
  }

  function Agent() public {
    agent = msg.sender;
    agentEnabled = true;
  }

  function setAgent(address addr) public isOwner {
    agent = addr;
  }

  function enableAgent(bool enable) public isOwner {
    agentEnabled = enable;
    require(getNumEnabledOwners() > 0);
  }

  function addCustodian(address addr) public isOwner {
    custodians.push(addr);
    custodianEnabled.push(true);
  }

  function removeCustodian(address addr) public isOwner {
    bool overwrite = false;
    for(uint i = 0; i < custodians.length; i++) {
      if(overwrite) {
        custodians[i - 1] = custodians[i];
        custodianEnabled[i-1] = custodianEnabled[i];
      }
      if(custodians[i] == addr) {
        overwrite = true;
      }
    }
    delete(custodians[custodians.length-1]);
    delete(custodianEnabled[custodianEnabled.length-1]);
    custodians.length -= 1;
    custodianEnabled.length -= 1;
    require(getNumEnabledOwners() > 0);
  }

  function enableCustodian(address addr, bool enable) public {
    for(uint i = 0; i < custodians.length; i++) {
      if(custodians[i] == addr) {
        custodianEnabled[i] = enable;
        break;
      }
    }
    require(getNumEnabledOwners() > 0);
  }

  function addRelationship(address r) public isOwner {
    relationships.push(r);
  }

  function getNumRelationships() public constant returns (uint) {
    return relationships.length;
  }

  function getNumEnabledOwners() public constant returns (uint) {
    uint num = 0;
    if(agentEnabled) num++;
    for(uint i = 0; i < custodianEnabled.length; i++) {
      if(custodianEnabled[i]) num++;
    }
    return num;
  }

  function getNumCustodians() public constant returns (uint) {
    return custodians.length;
  }
}
