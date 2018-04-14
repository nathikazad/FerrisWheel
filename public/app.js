
$("[type='number']").keypress(function (evt) {
    evt.preventDefault();
});

ws = new WebSocket('ws://' + window.location.host + '/ws');
ws.addEventListener('message', function(e) {
    var msg = JSON.parse(e.data);
    if(msg.method == "spin") {
        rotateAnimation(msg.arg0);
    }
});

function rotateAnimation(degrees){
    document.getElementById("Wheel").style.transform  = "rotate("+degrees+"deg)";
    document.getElementById("Chair1").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair2").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair3").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair4").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair5").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair6").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair7").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair8").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("degrees").innerHTML      = degrees;
}

window.addEventListener('load', function() {
    // Check if Web3 has been injected by the browser:
  if (typeof web3 !== 'undefined') {
    // You have a web3 browser! Continue below!
    startApp(web3);
  } else {
     // Warn the user that they need to get a web3 browser
     // Or install MetaMask, maybe with a nice graphic.
  }

  document.getElementById("getFT").onclick = function () { console.log('hello!'); };
  document.getElementById("approve-ft-button").onclick = approveFT;
  document.getElementById("increase-bid-button").onclick = increaseBid;
})

const Eth = require('ethjs');
const ferrisAddress = '0x2328ef76C4c55B317573f176b3C751522e7acFD7';
const ferrisContractABI = [{"constant":true,"inputs":[],"name":"beneficiary","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"withdraw","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"bid","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getFerrisTokenAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"getBid","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"chosenBidder","type":"address"},{"name":"amount","type":"uint256"}],"name":"accept","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"ferrisTokenAddress","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"eventId","type":"uint256"},{"indexed":false,"name":"bidder","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"NewBid","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"eventId","type":"uint256"},{"indexed":false,"name":"bidder","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"AcceptedBid","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"eventId","type":"uint256"},{"indexed":false,"name":"bidder","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"WithdrewBid","type":"event"}];
const ferrisTokenContractABI = [{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"INITIAL_SUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}];


async function startApp(web3) {
    eth = new Eth(web3.currentProvider);
    ferris = eth.contract(ferrisContractABI).at(ferrisAddress);
    const ferrisTokenAddress = await ferris.getFerrisTokenAddress();
    ferrisToken = eth.contract(ferrisTokenContractABI).at(ferrisTokenAddress[0]);
    coinbase = await eth.coinbase();
    updateView();
}

async function updateView() {
    document.getElementById("ferris-address").innerHTML = ferrisAddress;
    
    const beneficiary = await ferris.beneficiary();
    document.getElementById("beneficiary-address").innerHTML = beneficiary[0];
    
    document.getElementById("user-address").innerHTML = coinbase;
    
    const userBalance = await ferrisToken.balanceOf(coinbase);
    document.getElementById("user-balance").innerHTML = userBalance[0].toNumber();
    
    const userApprovedFTs = await(ferrisToken.allowance(coinbase, ferrisAddress))
    document.getElementById("user-approved-fts").innerHTML = userApprovedFTs[0].toNumber();
    document.getElementById("approve-ft-field").setAttribute("max", userBalance[0].toNumber() - userApprovedFTs[0].toNumber());

    const userBid = await ferris.getBid(coinbase);
    document.getElementById("user-bid").innerHTML = userBid[0].toNumber();
    document.getElementById("increase-bid-field").setAttribute("max", userApprovedFTs[0].toNumber() - userBid[0].toNumber());
}

async function waitForTxToBeMined (txHash) {
  $(':button').prop('disabled', true);
  let txReceipt
  while (!txReceipt) {
    try {
      txReceipt = await eth.getTransactionReceipt(txHash)
    } catch (err) {
      console.log("Tx Mine failed");
    }
  }
  console.log("Tx Mined");
  $(':button').prop('disabled', false);
  updateView();
}

function getFTs(){
    console.log("you wish");
}

async function approveFT(){
    console.log("trying to approve");
    const amount = parseInt(document.getElementById("approve-ft-field").value);
    document.getElementById("approve-ft-field").value = 0;
    
    ferrisToken.increaseApproval(ferrisAddress, amount, { from: coinbase }).then(function (txHash) {
      console.log('Transaction sent');
      console.dir(txHash);
      waitForTxToBeMined(txHash);
    });
}

async function increaseBid(){
    console.log("trying to increase bid");
    const amount = parseInt(document.getElementById("increase-bid-field").value);
    document.getElementById("increase-bid-field").value = 0;
    
    ferris.bid(amount, { from: coinbase }).then(function (txHash) {
      console.log('Transaction sent');
      console.dir(txHash);
      waitForTxToBeMined(txHash);
    });
}


