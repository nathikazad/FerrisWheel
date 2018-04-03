pragma solidity ^0.4.17;
import 'zeppelin-solidity/contracts/token/ERC20/StandardToken.sol';

contract FerrisToken is StandardToken {
	string public name = 'FerrisToken';
	string public symbol = 'FT';
	uint8 public decimals = 2;
	uint public INITIAL_SUPPLY = 12000;

	function FerrisToken() public {
	  totalSupply_ = INITIAL_SUPPLY;
	  balances[msg.sender] = INITIAL_SUPPLY;
	}
}