pragma solidity ^0.4.17;

import 'zeppelin-solidity/contracts/token/ERC20/StandardToken.sol';
import './FerrisToken.sol';

contract Ferris {
  // Parameters of the auction. 
  address public beneficiary;
  FerrisToken ferrisToken;

  // Allowed withdrawals of previous bids
  mapping(address => uint) bids;


  // Events that will be fired on changes.
  event NewBid(address bidder, uint amount);
  event AcceptedBid(address bidder, uint amount);
  event WithdrewBid(address bidder, uint amount);

  /// Create a Ferris with a single beneficiary
  function Ferris(address ferrisTokenAddress) {
    beneficiary = msg.sender;
    ferrisToken = FerrisToken(ferrisTokenAddress);
  }

  /// Bid for a ferris ride
  function bid(uint amount) returns (bool){
    require(ferrisToken.transferFrom(msg.sender, this, amount));
    bids[msg.sender] += amount;
    emit NewBid(msg.sender, amount);
  }

  /// Withdraw amount.
  function withdraw() returns (uint) {
    require (bids[msg.sender] > 0);
    uint amount = bids[msg.sender];
    bids[msg.sender] = 0;
    require(ferrisToken.transfer(msg.sender, amount));
    emit WithdrewBid(msg.sender, amount);
    return amount;
  }

  /// Accept the funds of the chosen bidder
  function accept(address chosenBidder, uint amount) returns (bool){
    require (msg.sender == beneficiary);
    require (bids[chosenBidder] >= amount);
    require(ferrisToken.transfer(beneficiary, amount));
    bids[chosenBidder] -= amount;
    emit AcceptedBid(chosenBidder, amount);
    return true;
  }
  
  function getBid(address addr) public view returns(uint) {
    return bids[addr];
  }

  function getFerrisTokenAddress() public view returns(address) {
    return address(ferrisToken);
  }
}


  