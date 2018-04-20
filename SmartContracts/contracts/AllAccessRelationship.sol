pragma solidity ^0.4.18;

contract AllAccessRelationship {
  address public patron;
  address public provider;
  string public providerName;

  struct Viewer {
    string name;
    address addr;
  }

  address[] public viewers;
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

  function setProviderName(string name) public {
    providerName = name;
  }

  function addViewer(string name, address viewer) public isPatron {
    require(!isViewer[viewer]);

    isViewer[viewer] = true;
    viewers.push(viewer);
    viewerInfo[viewer] = Viewer(name, viewer);
  }

  function removeViewer(address viewer) public isPatron {
    require(isViewer[viewer]);

    isViewer[viewer] = false;
    bool overwrite = false;
    for(uint i = 0; i < viewers.length; i++) {
      if(overwrite) {
        viewers[i - 1] = viewers[i];
      }
      if(viewers[i] == viewer) {
        overwrite = true;
      }
    }
    delete(viewers[viewers.length-1]);
    delete(viewerInfo[viewer]);
  }

  function getNumViewers() public constant returns(uint) {
    return viewers.length;
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
