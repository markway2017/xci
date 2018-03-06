pragma solidity ^0.4.18;

contract PermissionedWhiteList {

    address _owner;
    uint256 _nodeId;
    mapping (string => uint256) _whiteList;
    mapping (uint256 => string) _dids;

    function PermissionedWhiteList () public {
        _owner = msg.sender;
        _nodeId = 1;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    function addNewNode(string enode, string DIDJson) external onlyOwner {
        _whiteList[enode] = _nodeId;
        _dids[_nodeId] = DIDJson;

        _nodeId += 1;
    }

    function inWhiteList(string enode) public constant returns (bool) {
        return _whiteList[enode] > 0;
    }

    function getDID(string enode) public constant returns (string) {
        uint256 nodeId = _whiteList[enode];
        return _dids[nodeId];
    }
}