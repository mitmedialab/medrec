pragma solidity ^0.4.15;

//Contains
contract AgentRegistry {
  struct Agent {
    string name;
    address contractAddr;
    string host;
  }
  mapping(address => Agent) agentInfo;
  mapping(string => address) agentByName;
  mapping(address => address) public agentFromContract;

  //the current set of signers
  address[] signers;
  mapping(address => bool) public isSigner;
  //keeps track of votes per signer and the yay count
  mapping( address => address[]) voters;
  mapping( address => mapping(address => bool)) hasVoted;
  mapping( address => mapping(address => bool)) voteInfo;
  mapping( address => uint) yayVotes;
  //prospectives is a list of all prospective waiting to be voted
  address[] prospectives;
  //kicked is a list of all signers on the chopping block to be voted out
  address[] kicked;
  mapping(address => bool) public  isProspective;
  mapping(address => bool) public  isKicked;
  //keeps track of who's idea each proposal was
  mapping(address => address) proposer;

  event AddSigner(address indexed addr);
  event RemoveSigner(address indexed addr);

  modifier onlySigners() {
    require(isSigner[msg.sender]);
    _;
  }

  function AgentRegistry(string name, address contractAddr, string host) public {
    signers.push(msg.sender);
    agentInfo[msg.sender] = Agent(name, contractAddr, host);
    agentByName[name] = msg.sender;
    agentFromContract[contractAddr] = msg.sender;
    isSigner[msg.sender] = true;
  }

  // to prevent fraud only signers are allowed to set a name
  function setAgentName(string name) onlySigners() public {
      //throw if the proposed name is already taken
      require(agentByName[name] == address(0));

      agentInfo[msg.sender].name = name;
      agentByName[name] = msg.sender;
  }

  function setAgentContractAddr(address contractAddr) public {
      agentInfo[msg.sender].contractAddr = contractAddr;
      agentFromContract[contractAddr] = msg.sender;
  }

  function setAgentHost(string host) public {
      agentInfo[msg.sender].host = host;
  }

  //the caller proposes to add themself to the set of signers
  //the name is used to identify the new signer
  function propose(string name) public {
    require(!isSigner[msg.sender]);
    require(!isProspective[msg.sender]);

    prospectives.push(msg.sender);
    isProspective[msg.sender] = true;
    proposer[msg.sender] = msg.sender;

    //throw if the proposed name is already taken
    require(agentByName[name] == address(0));
    agentInfo[msg.sender].name = name;
  }

  //propose a signer to be kicked
  function kick (address rip) public onlySigners {
    require(isSigner[msg.sender]);
    require(!isKicked[rip]);
    //the blockchain cannot survive with less than 2 signers
    require(signers.length > 2);

    kicked.push(rip);
    isKicked[rip] = true;
    proposer[rip] = msg.sender;
    vote(rip, true);
  }

  //the caller recinds their proposal, deleting all cast votes
  function rescind (address prospective) public {
    require(proposer[prospective] == msg.sender);
    clearVotes(prospective);
  }

  function clearVotes(address prospective) internal {
    uint i;
    for(i = 0; i < voters[prospective].length; i++) {
      delete voteInfo[prospective][voters[prospective][i]];
      delete hasVoted[prospective][voters[prospective][i]];
    }
    delete voters[prospective];
    delete proposer[prospective];
    delete yayVotes[prospective];

    bool overwrite = false;
    if(isProspective[prospective]) {
      delete isProspective[prospective];
      for(i = 0; i < prospectives.length; i++) {
        if(overwrite) {
          prospectives[i - 1] = prospectives[i];
        }
        if(prospectives[i] == prospective) {
          overwrite = true;
        }
      }
      delete(prospectives[prospectives.length-1]);
      prospectives.length -= 1;
    } else {
      delete isKicked[prospective];
      for(i = 0; i < kicked.length; i++) {
        if(overwrite) {
          kicked[i - 1] = kicked[i];
        }
        if(kicked[i] == prospective) {
          overwrite = true;
        }
      }
      delete(kicked[kicked.length-1]);
      kicked.length -= 1;
    }
  }

  //an existing singer can vote or revote on an existing proposal
  function vote(address prospective, bool value) public onlySigners() {
    require(isProspective[prospective] || isKicked[prospective]);

    //if the signer voted yes before and votes no now remove their old vote
    if(voteInfo[prospective][msg.sender] && !value) {
      yayVotes[prospective] -= 1;
    }
    //if the signer is voting yes and hadn't before add a yay vote
    if(!voteInfo[prospective][msg.sender] && value) {
      yayVotes[prospective] += 1;
    }
    voteInfo[prospective][msg.sender] = value;
    if(!hasVoted[prospective][msg.sender]) {
      voters[prospective].push(msg.sender);
    }
    hasVoted[prospective][msg.sender] = true;

    //if this was a no vote it won't trigger anything and no further processing is needed
    if(!value) return;
    //if there aren't enought votes for a majority nothing more to do
    if(yayVotes[prospective] < (signers.length+1)/2) return;

    if(isKicked[prospective]) {
      bool overwrite = false;
      for(uint i = 0; i < signers.length; i++) {
        if(overwrite) {
          signers[i - 1] = signers[i];
        }
        if(signers[i] == prospective) {
          overwrite = true;
        }
      }
      delete(signers[signers.length-1]);
      signers.length -= 1;
      isSigner[prospective] = false;
      RemoveSigner(prospective);
    } else {
      signers.push(prospective);
      isSigner[prospective] = true;
      AddSigner(prospective);
    }

    clearVotes(prospective);
  }

  function getAgentByName(string name) public constant returns (address) {
    return agentByName[name];
  }

  function getAgentName(address addr) public constant returns (string) {
    return agentInfo[addr].name;
  }

  function getAgentContractAddr(address addr) public constant returns (address) {
    return agentInfo[addr].contractAddr;
  }

  function getAgentHost(address addr) public constant returns (string) {
    return agentInfo[addr].host;
  }

  function getNumSigners() public constant returns (uint) {
    return signers.length;
  }

  function getSigner(uint idx) public constant returns (address) {
    return signers[idx];
  }

  function getNumVoters(address prospective) public constant returns (uint) {
    return voters[prospective].length;
  }

  function getVoter(address prospective, uint idx) public constant returns (address) {
    return voters[prospective][idx];
  }

  function getVoteInfo(address prospective, address signer) public constant returns (bool) {
    return voteInfo[prospective][signer];
  }

  function getNumYayVotes(address prospective) public constant returns (uint) {
    return yayVotes[prospective];
  }

  function getNumProspectives() public constant returns (uint) {
    return prospectives.length;
  }

  function getProspective(uint idx) public constant returns (address) {
    return prospectives[idx];
  }

  function getNumKicked() public constant returns (uint) {
    return kicked.length;
  }

  function getKicked(uint idx) public constant returns (address) {
    return kicked[idx];
  }

  function getProposer(address prospective) public constant returns (address) {
    return proposer[prospective];
  }
}
