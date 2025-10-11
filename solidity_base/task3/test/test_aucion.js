const { ethers, deployments, getNamedAccounts } = require("hardhat")
const { expect } = require("chai")


describe("拍卖测试", function() {
  let NFT, AUC, DataFeed;
  this.timeout(60 * 1000); // 60s
  beforeEach(async () => {
    await deployments.fixture(["MyNFT", "Auction", "DataConsumerV3"]);
    const NFTDeployment = await deployments.get("MyNFT");
    NFT = await ethers.getContractAt("MyNFT", NFTDeployment.address);
    const AUCDeployment = await deployments.get("Auction");
    AUC = await ethers.getContractAt("Auction", AUCDeployment.address);
    const DataFeedDeployment = await deployments.get("DataConsumerV3");
    DataFeed = await ethers.getContractAt("DataConsumerV3", DataFeedDeployment.address);
  }
  );

  it("部署成功", async function() {
    const { deployer } = await getNamedAccounts(); // 确保这里能正确获取到账户
    expect(await NFT.owner()).to.equal(deployer);
    expect(await AUC.owner()).to.equal(deployer);
  });

  it("预言机测试", async function() {
    const val = await DataFeed.getChainlinkDataFeedETH2USD();
    // const val = await DataFeed.getChainlinkDataFeedLatestAnswer();
    console.log(`ETH2USD: ${val}`);
    expect(val).to.greaterThan(0);
  });

  it("NFT基本操作", async () => {
    const { deployer, user1, user2, user3 } = await getNamedAccounts(); // 确保这里能正确获取到账户
    // 给user1 铸造一个nft
    const URI = "https://bafybeiaal6hb63cqbnedvwderv6s7ny3thsdiu6slqabgv3ldcwliekeea.ipfs.dweb.link?filename=my_nft_meta.json";
    await NFT.safeMint(user1, URI);

    // 验证tokenId计数器增加
    const counter = await NFT._counter();
    expect(counter).to.equal(1);

    // 验证NFT所有权
    const ownerOfToken = await NFT.ownerOf(1);
    expect(ownerOfToken).to.equal(user1);

    // 验证Token URI
    const retrievedURI = await NFT.tokenURI(1);
    expect(retrievedURI).to.equal(URI);
  });

  it("测试拍卖", async () => {
    // 获取账户
    const { deployer, user1, user2, user3 } = await getNamedAccounts(); // 确保这里能正确获取到账户
    // user1铸造一个NFT
    const URI = "https://bafybeiaal6hb63cqbnedvwderv6s7ny3thsdiu6slqabgv3ldcwliekeea.ipfs.dweb.link?filename=my_nft_meta.json";
    await NFT.safeMint(user1, URI);

    // user1 给拍卖合约授权
    const deployerS = await ethers.getSigner(deployer);
    const user1S = await ethers.getSigner(user1);
    const user2S = await ethers.getSigner(user2);
    const user3S = await ethers.getSigner(user3);
    await NFT.connect(user1S).approve(AUC.getAddress(), 1);

    // 创建一个拍卖, 1ETH起拍
    await AUC.connect(deployerS).createAuction(30, ethers.parseEther("1"), 1, NFT.getAddress());

    // deployer 出价 3 ETH
    await AUC.connect(deployerS).bid(1, { value: ethers.parseEther("3") });

    // user2 出价 2wei, 非法
    try {
      await AUC.connect(user2S).bid(1, { value: ethers.parseEther("2") });
    } catch (error) {
      console.log(error.message);
    }

    // user3 出价5 ETH 
    await AUC.connect(user3S).bid(1, { value: ethers.parseEther("5") });

    // 当前的最高出价者，应该是user3
    expect(await AUC.highestBidAddr(1)).to.equal(user3);

    // 当前最高价 150
    expect(await AUC.highestBidPrice(1)).to.equal(ethers.parseEther("5"));

    //查看各个账户余额 
    const accounts = await getNamedAccounts(); // 确保这里能正确获取到账户
    let i = 0;
    for (const acc in accounts) {
      const balance = await ethers.provider.getBalance(accounts[acc])
      const balanceEth = ethers.formatEther(balance); // 自动转换为ETH单位
      console.log(`account${i} balance:${balanceEth}`)
      i++
    }
    // user1 进行结束交易，报错
    try {
      await AUC.connect(user1S).end(1);
    } catch (error) {
      console.log(error.message);
    }
    // deploy 进行结束交易，报错，
    try {
      await AUC.connect(deployerS).end(1);
    } catch (error) {
      console.log(error.message);
    }
    //等待 30s
    console.log("waiting for 30s...");
    await new Promise(resolve => { setTimeout(resolve, 30 * 1000) });

    // deploy 进行结束交易
    await AUC.connect(deployerS).end(1);
    // 此时token 1 应该属于user3
    expect(await NFT.ownerOf(1)).to.equal(user3);
  });

  // it('should complete auction', function(done) {  // 非async函数
  //   auction.start().then(() => done()).catch(done);
  // });
});

