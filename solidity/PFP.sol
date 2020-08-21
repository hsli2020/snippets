//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract PFP is ERC721Enumerable, Ownable {
    using Counters for Counters.Counter;
    using Strings for uint256;
    uint256 MAX_TO_MINT = 5000;
    uint256 totalMinted;
    string public URI;
    bool public REVEAL = false;
    uint256 tokenPrice;

    constructor(
        string memory name,
        string memory symbol,
        string memory initialURI,
        uint256 initialPrice
    ) ERC721(name, symbol) {
        URI = initialURI;
        setPrice(initialPrice);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override
        returns (string memory)
    {
        if (REVEAL) {
            return string(abi.encodePacked(URI, tokenId.toString()));
        }
        return URI;
    }

    function toggleReveal(string memory updatedURI) public onlyOwner {
        REVEAL = !REVEAL;
        URI = updatedURI;
    }

    function setPrice(uint256 price) public onlyOwner {
        tokenPrice = price;
    }

    function mint(address addr) public payable returns (uint256) {
        require(msg.value >= tokenPrice, "insufficient funds");
        require(totalMinted < MAX_TO_MINT, "Would exceed max supply");
        uint256 tokenId = totalMinted + 1;
        totalMinted += 1;
        _safeMint(addr, tokenId);

        return tokenId;
    }
}