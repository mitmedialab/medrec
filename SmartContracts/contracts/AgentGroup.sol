pragma solidity ^0.4.15;

//Represents a set of agents who represent themself as one
contract AgentGroup {
  address[] public agents;

  modifier isOwner() {
    bool enable;
    for(uint i = 0; i < agents.length; i++) {
      if(msg.sender == agents[i]) {
        enable = true;
        break;
      }
    }
    if(!enable) revert();
    _;
  }

  function AgentGroup() public {
    agents.push(msg.sender);
  }

  function addAgent(address addr) public isOwner {
    agents.push(addr);
  }

  function removeAgent(address addr) public isOwner {
    bool overwrite = false;
    for(uint i = 0; i < agents.length; i++) {
      if(overwrite) {
        agents[i - 1] = agents[i];
      }
      if(agents[i] == addr) {
        overwrite = true;
      }
    }
    delete(agents[agents.length-1]);
    agents.length -= 1;
    require(getNumAgents() > 0);
  }

  function getNumAgents() public constant returns (uint) {
    return agents.length;
  }
}
