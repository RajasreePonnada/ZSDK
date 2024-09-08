// Identity.sol
pragma solidity ^0.8.0;

contract Identity {
    struct IdentityData {
        bytes32 dataHash;
        uint256 lastUpdated;
    }

    mapping(address => IdentityData) public identities;

    event IdentityCreated(address indexed user, bytes32 dataHash);
    event IdentityUpdated(address indexed user, bytes32 newDataHash);

    function createIdentity(bytes32 _dataHash) external {
        require(identities[msg.sender].lastUpdated == 0, "Identity already exists");
        identities[msg.sender] = IdentityData(_dataHash, block.timestamp);
        emit IdentityCreated(msg.sender, _dataHash);
    }

    function updateIdentity(bytes32 _newDataHash) external {
        require(identities[msg.sender].lastUpdated != 0, "Identity does not exist");
        identities[msg.sender].dataHash = _newDataHash;
        identities[msg.sender].lastUpdated = block.timestamp;
        emit IdentityUpdated(msg.sender, _newDataHash);
    }

    // Add more functions as needed
}