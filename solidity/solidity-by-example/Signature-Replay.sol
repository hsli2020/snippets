//Signature Replay
//
//Signing messages off-chain and having a contract that requires that signature before 
//executing a function is a useful technique.
//
//For example this technique is used to:
//  - reduce number of transaction on chain
//  - gas-less transaction, called meta transaction
//
//Vulnerability
//
//Same signature can be used multiple times to execute a function. This can be harmful if 
//the signer's intention was to approve a transaction once.

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
pragma experimental ABIEncoderV2;

import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/
        contracts/utils/cryptography/ECDSA.sol";

contract MultiSigWallet {
    using ECDSA for bytes32;
    address[2] public owners;

    constructor(address[2] memory _owners) payable {
        owners = _owners;
    }

    function deposit() external payable {}

    function transfer(address _to, uint _amount, bytes[2] memory _sigs) external {
        bytes32 txHash = getTxHash(_to, _amount);
        require(_checkSigs(_sigs, txHash), "invalid sig");

        (bool sent, ) = _to.call{value: _amount}("");
        require(sent, "Failed to send Ether");
    }

    function getTxHash(address _to, uint _amount) public view returns (bytes32) {
        return keccak256(abi.encodePacked(_to, _amount));
    }

    function _checkSigs(bytes[2] memory _sigs, bytes32 _txHash)
        private view returns (bool)
    {
        bytes32 ethSignedHash = _txHash.toEthSignedMessageHash();

        for (uint i = 0; i < _sigs.length; i++) {
            address signer = ethSignedHash.recover(_sigs[i]);
            bool valid = signer == owners[i];

            if (!valid) {
                return false;
            }
        }

        return true;
    }
}

//Preventative Techniques
//   Sign messages with nonce and address of the contract.

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
pragma experimental ABIEncoderV2;

import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/
        contracts/utils/cryptography/ECDSA.sol";

contract MultiSigWallet {
    using ECDSA for bytes32;

    address[2] public owners;
    mapping(bytes32 => bool) public executed;

    constructor(address[2] memory _owners) payable { owners = _owners; }

    function deposit() external payable {}

    function transfer(address _to, uint _amount, uint _nonce, bytes[2] memory _sigs) external {
        bytes32 txHash = getTxHash(_to, _amount, _nonce);
        require(!executed[txHash], "tx executed");
        require(_checkSigs(_sigs, txHash), "invalid sig");

        executed[txHash] = true;

        (bool sent, ) = _to.call{value: _amount}("");
        require(sent, "Failed to send Ether");
    }

    function getTxHash(address _to, uint _amount, uint _nonce) public view returns (bytes32) {
        return keccak256(abi.encodePacked(address(this), _to, _amount, _nonce));
    }

    function _checkSigs(bytes[2] memory _sigs, bytes32 _txHash) private view returns (bool)
    {
        bytes32 ethSignedHash = _txHash.toEthSignedMessageHash();

        for (uint i = 0; i < _sigs.length; i++) {
            address signer = ethSignedHash.recover(_sigs[i]);
            bool valid = signer == owners[i];

            if (!valid) {
                return false;
            }
        }

        return true;
    }
}

/*
// owners
0xe19aea93F6C1dBef6A3776848bE099A7c3253ac8
0xfa854FE5339843b3e9Bfd8554B38BD042A42e340

// to
0xe10422cc61030C8B3dBCD36c7e7e8EC3B527E0Ac
// amount
100
// nonce
0
// tx hash
0x12a095462ebfca27dc4d99feef885bfe58344fb6bb42c3c52a7c0d6836d11448

// signatures
0x120f8ed8f2fa55498f2ef0a22f26e39b9b51ed29cc93fe0ef...c40bb8720de6d8f2523077ad6d1138a60923b801c
0xa240a487de1eb5bb971e920cb0677a47ddc6421e38f7b048f...18469f9eec2c59e87201bbc8db0e4d9796935cb1b
*/
