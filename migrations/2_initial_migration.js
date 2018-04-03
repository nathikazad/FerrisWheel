var FerrisToken = artifacts.require("FerrisToken");

module.exports = function(deployer) {
  deployer.deploy(FerrisToken);
};