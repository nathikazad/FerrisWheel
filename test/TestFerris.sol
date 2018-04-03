pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import '../contracts/FerrisToken.sol';
import "../contracts/Ferris.sol";

contract TestFerris {
  
  function testFerrisBeneficiaryDeployed() {
  	Ferris ferris = Ferris(DeployedAddresses.Ferris());
    Assert.equal(ferris.beneficiary(), tx.origin, "Beneficiary mismatch");
  }

  function testFerrisBeneficiaryNew() {
  	FerrisToken ferristoken = FerrisToken(DeployedAddresses.FerrisToken());
  	Ferris ferris = new Ferris(ferristoken);
    Assert.equal(ferris.beneficiary(), this, "Beneficiary mismatch");
  }

}