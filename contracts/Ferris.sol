pragma solidity ^0.4.17;

contract Ferris {
  // Parameters of the auction. 
  address public beneficiary;


  // Allowed withdrawals of previous bids
  mapping(address => uint) bids;


  // Events that will be fired on changes.
  event NewBid(address bidder, uint amount);
  event AcceptedBid(address bidder, uint amount);

  /// Create a Ferris with a single beneficiary
  function Ferris() {
    beneficiary = msg.sender;
  }

  /// Bid for a ferris ride
  function bid() payable {
    bids[msg.sender] += msg.value;
    emit NewBid(msg.sender, msg.value);
  }

  /// Withdraw amount.
  function withdraw() returns (bool) {
    require (bids[msg.sender] > 0);
    uint amount = bids[msg.sender];
    bids[msg.sender] = 0;
    require(msg.sender.send(amount));
    return true;
  }

  /// Accept the funds of the chosen bidder
  function accept(address chosenBidder, uint amount) {
    require (msg.sender == beneficiary);
    require (bids[chosenBidder] >= amount);
    require (beneficiary.send(amount));
    bids[chosenBidder] -= amount;
    emit AcceptedBid(chosenBidder, amount);
  }
  
  function getBid(address addr) public view returns(uint) {
    return bids[addr];
  }
}


  