// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import './tokens/ERC721URIStorage.sol';
import './interfaces/IERC20.sol';
import './utils/ownable.sol';
import './libraries/Strings.sol';
import './libraries/SafeTransfers.sol';

contract IdentDeFi is ERC721URIStorage, Ownable, ChainlinkClient {
  using Chainlink for Chainlink.Request;
  using Strings for address;

  string public path;
  address public oracle;
  bytes32 public jobId;
  uint public fee;
  uint public price;
  mapping(address => uint) private _tokens;
  mapping(bytes32 => address) private requestToMint;
  uint private _counter;

  modifier onlyOwnerOrHolder(uint tokenId) {
    require(msg.sender == owner || tokenOf(msg.sender) == tokenId, "IdentDeFI::onlyOwnerOrHolder: Invalid signer address.");
    _;
  }

  constructor(
    string memory _path,
    address _oracle,
    uint _fee,
    uint _price
  )
  ERC721("IdentDeFi", "IDF")
  {
    setPublicChainlinkToken();
    path = _path;
    oracle = _oracle;
    jobId = 'bc746611ebee40a3989bbe49e12a02b9';
    fee = _fee;
    price = _price;
  }

  function tokenOf(address account) public view returns (uint) {
    return _tokens[account];
  }

  function verified(address account) public view returns (bool) {
    return tokenOf(account) > 0;
  }

  function revoke(uint tokenId) external onlyOwnerOrHolder(tokenId) {
    address owner = ownerOf(tokenId);
    require(owner != address(0), "IdentDeFI::revoke: Invalid token ID");
    delete _tokens[owner];
    emit TokenRevoked(tokenId, owner, msg.sender);
  }

  function mintVerification() public payable {
    require(msg.value >= price, "IdentDeFI::mintVerification: Paid value is insufficiant");
    require(balanceOf(msg.sender) < 1, "IdentDeFI::mintVerification: Only 1 token per address allowed");
    bytes32 requestId = requestValidation();
    requestToMint[requestId] = msg.sender;
  }

  function requestValidation() internal returns (bytes32 requestId) {
    Chainlink.Request memory req = buildChainlinkRequest(jobId, address(this), this.fulfill.selector);
    // Set the URL to perform the GET request on
    req.add("get", string(abi.encodePacked(path, "/account/", msg.sender.addressToString())));
    // Define path to value in JSON
    req.add("path", "data.valid");
    // Sends the request
    return sendChainlinkRequestTo(oracle, req, fee);
  }

  function fulfill(bytes32 _requestId, bool _valid) public recordChainlinkFulfillment(_requestId) {
    address account = requestToMint[_requestId];

    if (_valid) {
      mint(account);
    } else {
      emit InvalidValidation(account);
    }
  }

  function mint(address account) private {
    _counter += 1;
    _safeMint(account, _counter);
    _setTokenURI(_counter, string(abi.encodePacked(path, "/account/", account.addressToString(), "/metadata")));
    _tokens[account] = _counter;

    emit ValidationSuccess(account, _counter);
  }

  function _transfer(
    address from,
    address to,
    uint tokenId
  ) internal override {
    revert('IdentDeFI::transfer: Not allowed');
  }

  function withdraw(uint _amount) external onlyOwner {
    uint balance = address(this).balance;
    require(_amount <= balance, "IdentDeFI::withdraw: Insufficient balance");
    SafeTransfers.safeTransferETH(msg.sender, balance);
  }

  function tokenBalance(address _tokenContract) view public returns (uint) {
    IERC20 token = IERC20(_tokenContract);
    return token.balanceOf(address(this));
  }

  function withdrawToken(address _tokenContract, uint _amount) external onlyOwner {
    require(_amount <= tokenBalance(_tokenContract), "IdentDeFI::withdrawToken: Insufficient balance");
    SafeTransfers.safeTransfer(_tokenContract, msg.sender, _amount);
  }

  receive() external payable {
    mintVerification();
  }

  event ValidationSuccess(address indexed _account, uint indexed _tokenId);
  event InvalidValidation(address indexed _account);
  event TokenRevoked(uint indexed _tokenId, address indexed _account, address indexed _revoker);
}
