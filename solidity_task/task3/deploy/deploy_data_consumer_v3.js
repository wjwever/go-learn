
module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts(); // 确保这里能正确获取到账户

  // const { deployer } = await ethers.getSigners();

  // 确保 deployer 有值
  if (!deployer) {
    throw new Error("Deployer account not found");
  }

  await deploy("DataConsumerV3", {
    from: deployer,
    args: [], // 如果有构造函数参数，确保正确设置
    log: true,
  });
};

module.exports.tags = ["DataConsumerV3"];

