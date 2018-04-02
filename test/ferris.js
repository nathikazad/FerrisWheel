require('truffle-test-utils').init();
var Ferris = artifacts.require("./Ferris.sol");

contract('Ferris', function(accounts) {
  it("should check beneficiary", async () => {
  	let ferris = await Ferris.deployed();
    let beneficiary = await ferris.beneficiary();
    assert.equal(beneficiary, accounts[0], "Address is different");
  });

  it("should bid correctly", async () => {
  	let ferris = await Ferris.deployed();
  	let actual = await ferris.getBid(accounts[0]);
  	var expected = 0;
  	assert.equal(actual, expected, "Value is not 0");
  	let result = await ferris.bid({value: 1});
  	assert.web3Event(result, {
	  event: 'NewBid',
	    args: {
	      bidder: accounts[0],
	      amount: 1 
	  }
	}, 'The event is emitted');
  	actual = await ferris.getBid(accounts[0]);
  	expected = 1;
  	assert.equal(actual, expected, "Value is not 1");
  });

  it("should accept correctly", async () => {
  	let ferris = await Ferris.deployed();
  	await ferris.bid({value: 1});
  	let result = await ferris.accept(accounts[0], 1, {from:accounts[0]});
  	assert.web3Event(result, {
	  event: 'AcceptedBid',
	    args: {
	      bidder: accounts[0],
	      amount: 1 
	  }
	}, 'The event is emitted');
  });

  it("should not accept from non beneficiary", async () => {
  	let ferris = await Ferris.deployed();
  	await ferris.bid({value: 1});
  	let err = null;
  	try {
  		await ferris.accept(accounts[0], 1, {from:accounts[1]});
  	} catch (error) {
  		assert(
          error.message.search('revert'),
          "Expected revert, got '" + error + "' instead",
        );
        return;
  	}
  	assert.fail('Expected throw not received');
  });

  it("should not accept if less money", async () => {
  	let ferris = await Ferris.new();
  	let actual = await ferris.getBid(accounts[0]);
  	var expected = 0;
  	assert.equal(actual, expected, "Value is not 0");
  	let err = null;
  	try {
  		await ferris.accept(accounts[0], 1, {from:accounts[0]});
  	} catch (error) {
  		assert(
          error.message.search('revert'),
          "Expected revert, got '" + error + "' instead",
        );
        return;
  	}
  	assert.fail('Expected throw not received');
  });

  it("should increase bid when user adds", async () => {
    let ferris = await Ferris.new();
    await ferris.bid({value: 1});
    let actual = await ferris.getBid(accounts[0]);
    var expected = 1;
    assert.equal(actual, expected, "Value is not 1");
    await ferris.bid({value: 1});
    actual = await ferris.getBid(accounts[0]);
    expected = 2;
    assert.equal(actual, expected, "Value is not 2");
  });

  it("should decrease bid when ben. accepts", async () => {
    let ferris = await Ferris.new();
    await ferris.bid({value: 2});
    let actual = await ferris.getBid(accounts[0]);
    var expected = 2;
    assert.equal(actual, expected, "Value is not 2");
    await ferris.accept(accounts[0], 1, {from:accounts[0]});
    actual = await ferris.getBid(accounts[0]);
    expected = 1;
    assert.equal(actual, expected, "Value is not 1");
  });

  it("should withdraw correctly", async () => {
    let ferris = await Ferris.new();
    await ferris.bid({value: 2});
    let actual = await ferris.getBid(accounts[0]);
    var expected = 2;
    assert.equal(actual, expected, "Value is not 2");
    let result = await ferris.withdraw();
    assert.web3Event(result, {
      event: 'WithdrewBid',
        args: {
          bidder: accounts[0],
          amount: 2 
      }
    }, 'The event is emitted');
    actual = await ferris.getBid(accounts[0]);
    expected = 0;
    assert.equal(actual, expected, "Value is not 0");
  });

  it("should not withdraw if no money", async () => {
    let ferris = await Ferris.new();
    let actual = await ferris.getBid(accounts[0]);
    var expected = 0;
    assert.equal(actual, expected, "Value is not 0");
    let err = null;
    try {
      await ferris.withdraw();
    } catch (error) {
      assert(
          error.message.search('revert'),
          "Expected revert, got '" + error + "' instead",
        );
        return;
    }
    assert.fail('Expected throw not received');
  });



});

//Console variables
//var ferris = Ferris.deployed();
//web3.eth.accounts
//Ferris.deployed().then( function(instance) { return instance.beneficiary.call() })
