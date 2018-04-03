require('truffle-test-utils').init();
var Ferris = artifacts.require("./Ferris.sol");
var FerrisToken = artifacts.require("./FerrisToken.sol");

contract('Ferris', function(accounts) {
  it("should check beneficiary", async () => {
    let ferris = await Ferris.deployed();
    let beneficiary = await ferris.beneficiary();
    assert.equal(beneficiary, accounts[0], "Address is different");
  });

  it("should check ferris token address", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.deployed();
    var expected = await ferris.getFerrisTokenAddress();
    assert.equal(ferrisToken.address, expected, "ferris token address different");
  });

  it("should bid correctly", async () => {
    var bidAmount = 1;
    var userAccount = web3.eth.coinbase;
    let ferrisToken = await FerrisToken.deployed();
  	let ferris = await Ferris.deployed();
    let userTokenBalance = await ferrisToken.balanceOf(userAccount);
    let ferrisBalance = await ferrisToken.balanceOf(ferris.address);
    let userBidBalance = await ferris.getBid(userAccount);
    await ferrisToken.approve(ferris.address, bidAmount, {from:userAccount});
  	let result = await ferris.bid(bidAmount);
  	// Ferris Token Assert Userbalance reduce and ferris balance increase
    let actual = await ferrisToken.balanceOf(userAccount);
    let expected = userTokenBalance.sub(bidAmount);
    assert.ok(actual.eq(expected), "User token Balance is not "+expected.toString());
    actual = await ferrisToken.balanceOf(ferris.address);
    expected = ferrisBalance.plus(bidAmount);
    assert.ok(actual.eq(expected), "Ferris token Balance is not "+expected.toString());
    // Ferris assert User bid increase
    actual = await ferris.getBid(userAccount);
  	expected = userBidBalance.plus(bidAmount);
  	assert.ok(actual.eq(expected), "User Bid Balance is not "+expected.toString());
  });

    it("should emit bid event correctly", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.deployed();
    let actual = await ferris.getBid(accounts[0]);
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount);
    let result = await ferris.bid(bidAmount);
    assert.web3Event(result, {
      event: 'NewBid',
        args: {
          bidder: accounts[0],
          amount: bidAmount 
      }
    }, 'The event is emitted');
  });

  it("should accept correctly", async () => {
    let ferrisToken = await FerrisToken.deployed();
  	let ferris = await Ferris.deployed();
    let beneficiary = await ferris.beneficiary();
    var userAccount = web3.eth.coinbase;
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount);
  	assert.isOk(await ferris.bid(bidAmount));
    let beneficiaryTokenBalance = await ferrisToken.balanceOf(beneficiary);
    let ferrisBalance = await ferrisToken.balanceOf(ferris.address);
    let userBidBalance = await ferris.getBid(userAccount);
  	let result = await ferris.accept(accounts[0], bidAmount);
    // Ferris Token Assert beneficiary balance increase and ferris balance decrease
    let actual = await ferrisToken.balanceOf(beneficiary);
    let expected = beneficiaryTokenBalance.plus(bidAmount);
    assert.ok(actual.eq(expected), "Beneficiary token Balance is not "+expected.toString());
    actual = await ferrisToken.balanceOf(ferris.address);
    expected = ferrisBalance.sub(bidAmount);
    assert.ok(actual.eq(expected), "Ferris token Balance is not "+expected.toString());
    // Ferris assert
    actual = await ferris.getBid(userAccount);
    expected = userBidBalance.sub(bidAmount);
    assert.ok(actual.eq(expected), "User Bid Balance is not "+expected.toString());
  });

    it("should emit accept events correctly", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.deployed();
    var userAccount = web3.eth.coinbase;
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount);
    assert.isOk(await ferris.bid(bidAmount));
    let result = await ferris.accept(userAccount, bidAmount);
    assert.web3Event(result, {
      event: 'AcceptedBid',
        args: {
          bidder: userAccount,
          amount: bidAmount 
        }
    }, 'The event is emitted');
  });

  it("should not accept from non beneficiary", async () => {
  	let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.deployed();
    var userAccount = web3.eth.coinbase;
    var nonBeneficiary = accounts[1];
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount);
  	await ferris.bid(bidAmount);
  	let err = null;
  	try {
  		await ferris.accept(userAccount, bidAmount, {from:nonBeneficiary});
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
    let ferrisToken = await FerrisToken.deployed();
  	let ferris = await Ferris.new(ferrisToken.address);
    var userAccount = web3.eth.coinbase;
    var bidAmount = 1;
  	let err = null;
  	try {
  		await ferris.accept(userAccount, bidAmount, {from:userAccount});
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
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.new(ferrisToken.address);
    var userAccount = web3.eth.coinbase;
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount * 2);
    await ferris.bid(bidAmount);
    let actual = await ferris.getBid(userAccount);
    var expected = 1;
    assert.equal(actual, expected, "Value is not 1");
    await ferris.bid(bidAmount);
    actual = await ferris.getBid(userAccount);
    expected = bidAmount * 2;
    assert.equal(actual, expected, "Value is not 2");
  });

  it("should decrease bid when ben. accepts", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.new(ferrisToken.address);
    var userAccount = web3.eth.coinbase;
    var bidAmount = 1;
    await ferrisToken.approve(ferris.address, bidAmount * 2);
    await ferris.bid(bidAmount * 2);
    let actual = await ferris.getBid(userAccount);
    var expected = 2;
    assert.equal(actual, expected, "Value is not 2");
    await ferris.accept(userAccount, bidAmount);
    actual = await ferris.getBid(userAccount);
    expected = bidAmount;
    assert.equal(actual, expected, "Value is not 1");
  });

  it("should withdraw correctly", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.new(ferrisToken.address);
    var userAccount = web3.eth.coinbase;
    var bidAmount = 2;
    await ferrisToken.approve(ferris.address, bidAmount);
    await ferris.bid(bidAmount);

    let userTokenBalance = await ferrisToken.balanceOf(userAccount);
    let ferrisBalance = await ferrisToken.balanceOf(ferris.address);
    let userBidBalance = await ferris.getBid(userAccount);
    await ferris.withdraw();
    // Ferris Token Assert Userbalance reduce and ferris balance increase
    let actual = await ferrisToken.balanceOf(userAccount);
    let expected = userTokenBalance.plus(bidAmount);
    assert.ok(actual.eq(expected), "User token Balance is not "+expected.toString());
    actual = await ferrisToken.balanceOf(ferris.address);
    expected = ferrisBalance.sub(bidAmount);
    assert.ok(actual.eq(expected), "Ferris token Balance is not "+expected.toString());
    // Ferris assert User bid increase
    actual = await ferris.getBid(userAccount);
    expected = userBidBalance.sub(bidAmount);
    assert.ok(actual.eq(expected), "User Bid Balance is not "+expected.toString());
  });

  it("should emit withdraw event correctly", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.new(ferrisToken.address);
    var userAccount = web3.eth.coinbase;
    var bidAmount = 2;
    await ferrisToken.approve(ferris.address, bidAmount);
    await ferris.bid(bidAmount);
    let result = await ferris.withdraw();
    assert.web3Event(result, {
      event: 'WithdrewBid',
        args: {
          bidder: userAccount,
          amount: bidAmount 
      }
    }, 'The event is emitted');
  });

  it("should not withdraw if no money", async () => {
    let ferrisToken = await FerrisToken.deployed();
    let ferris = await Ferris.new(ferrisToken.address);
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
