pragma solidity ^0.4.18;

contract Relationship {
    address public patron;
    address public provider; // this is a unique address for the provider used only for this relationship
    string public providerAddr; //encrypted provider address
    string public providerName; //encrypted provider name

    struct Viewer {
        string name;
        string providerAddr; //the real provider address encypted so only the viewer can read
    }

    address[][] viewerGroups;
    mapping(address => bool) public isViewer;
    mapping(address => Viewer) viewerInfo;
    mapping(string => address) viewerByName;

    uint256 constant UINT256_MAX = ~uint256(0);

    modifier isPatron() {
        if(msg.sender != patron) revert();
        _;
    }

    function Relationship(address _provider) public {
        patron = msg.sender;
        provider = _provider;
    }

/****These functions should be left commented out until a use case for them arises
  function setPatron(address addr) isPatron {
    patron = addr;
  }
  function setProvider(address addr) isPatron {
    provider = addr;
  }
******************/

function setProviderAddress(string addr) public {
    providerAddr = addr;
}

function setProviderName(string name) public {
  providerName = name;
}

function addViewerGroup() public isPatron {
    viewerGroups.length += 1;
}

function removeViewerGroup(uint viewerGroup) public isPatron {
    uint numViewers = viewerGroups[viewerGroup].length;
    uint i;
    for(i = 0; i < numViewers; i++) {
        isViewer[viewerGroups[viewerGroup][i]] = false;
    }

    uint numGroups = viewerGroups.length;
    for(i = viewerGroup+1; i < numGroups; i++) {
        viewerGroups[i - 1] = viewerGroups[i];
    }
    delete(viewerGroups[numGroups-1]);
    viewerGroups.length -= 1;
}

function addViewer(string name, uint viewerGroup, address viewer, string provAddr) public isPatron {
    require(!isViewer[viewer]);

    isViewer[viewer] = true;
    viewerGroups[viewerGroup].push(viewer);
    viewerInfo[viewer] = Viewer(name, provAddr);
}

function removeViewer(uint viewerGroup, address viewer) public isPatron {
    require(isViewer[viewer]);

    isViewer[viewer] = false;
    uint numViewers = viewerGroups[viewerGroup].length;
    bool overwrite = false;
    for(uint i = 0; i < numViewers; i++) {
        if(overwrite) {
            viewerGroups[viewerGroup][i - 1] = viewerGroups[viewerGroup][i];
        }
        if(viewerGroups[viewerGroup][i] == viewer) {
            overwrite = true;
        }
    }
    delete(viewerGroups[viewerGroup][numViewers-1]);
    viewerGroups[viewerGroup].length -= 1;
    delete(viewerInfo[viewer]);
}

function getNumViewerGroups() public constant returns(uint) {
    return viewerGroups.length;
}

function getNumViewers(uint group) public constant returns(uint) {
    return viewerGroups[group].length;
}

function getViewer(uint group, uint index) public constant returns(address) {
  return viewerGroups[group][index];
}

function getViewerByName(string name) public constant returns(address) {
    return viewerByName[name];
}

function getViewerName(address addr) public constant returns(string) {
    return viewerInfo[addr].name;
}

function terminate() public {
    if(msg.sender != patron && msg.sender != provider) revert();
    selfdestruct(patron);
}
}
