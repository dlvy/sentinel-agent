// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract SmartAccount {
    address public owner;
    mapping(address => uint256) public sessionKeys;

    constructor(address _owner) {
        owner = _owner;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    function setSessionKey(address key, uint256 expiresAt) external onlyOwner {
        sessionKeys[key] = expiresAt;
    }

    function execute(address target, bytes calldata data) external {
        require(
            msg.sender == owner || sessionKeys[msg.sender] > block.timestamp,
            "Unauthorized"
        );
        (bool success, ) = target.call(data);
        require(success, "Call failed");
    }
}
