var FerrisToken = artifacts.require("FerrisToken");
var Ferris = artifacts.require("Ferris");

module.exports = function(deployer) {
	deployer.deploy(Ferris, FerrisToken.address);
};
