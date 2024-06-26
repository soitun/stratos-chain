import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { proxyAdmin, deployer } = await getNamedAccounts();

  await deploy("SystemContract", {
    from: deployer,
    proxy: {
      owner: proxyAdmin,
      execute: {
        init: {
          methodName: "initialize",
          args: [],
        },
      },
      proxyContract: "OpenZeppelinTransparentProxy",
    },
    skipIfAlreadyDeployed: false,
    log: true,
  });
};
func.tags = ["0001_deploy_system_contract"];

export default func;
