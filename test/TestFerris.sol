pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Ferris.sol";

contract TestFerris {
  
  function testFerrisBeneficiaryDeployed() {
  	Ferris ferris = Ferris(DeployedAddresses.Ferris());
    Assert.equal(ferris.beneficiary(), tx.origin, "yolo");
  }

  function testFerrisBeneficiaryNew() {
  	Ferris ferris = new Ferris();
    Assert.equal(ferris.beneficiary(), this, "yolo");
  }

}