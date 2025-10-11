// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;
import "./MyNFT.sol";

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract Auction {
    address _owner;
    mapping(uint256 => Tx) _txes;
    uint256 _txId;
    struct Tx {
        //创建者
        address _creator;
        // nft合约地址
        address _nftAddr;
        // nft tokenId
        uint256 _nftTokenId;
        // 开始时间
        uint256 _startTime;
        // 锁定期
        uint256 _lockTime;
        // 当前最高价 us dollar, 初始要设置一个起拍价
        int256 highestUSD;
        // 最高出价者
        address _highestBidder;
        // ERC20 合约地址
        address _tokenAddr;
        // 最高出价
        uint256 _amount;
        // 整个交易是否结束
        bool _ended;
    }

    //  address(0) => sepolia eth
    mapping(address => AggregatorV3Interface) internal _priceFeeds;

    constructor() {
        _owner = msg.sender;
    }

    function owner() public view returns (address) {
        return _owner;
    }

    function setPriceFeeder(address addr, address aggAddr) external {
        _priceFeeds[addr] = AggregatorV3Interface(aggAddr);
    }

    // 接受NFT
    // function onERC721Received(
    //     address operator,
    //     address from,
    //     uint256 tokenId,
    //     bytes calldata data
    // ) external pure override returns (bytes4) {
    //     return IERC721Receiver.onERC721Received.selector;
    // }

    // 创建拍卖
    function createAuction(
        uint256 lockTime,
        int256 startPrice,
        uint256 nftTokenId,
        address nftAddress
    ) external {
        // owner 可以创建拍卖
        // require(msg.sender == _owner, "Only owner!");
        require(lockTime > 0, "lockTime > 0");
        require(startPrice > 0, "startPrice > 0");
        // 是否取得授权
        require(
            MyNFT(nftAddress).getApproved(nftTokenId) == address(this),
            "NFT not approved"
        );

        // 把nft转移到本合约, msg.sender 必须持有NFT
        // MyNFT(nftAddress).safeTransferFrom(
        //     msg.sender,
        //     address(this),
        //     nftTokenId
        // );

        ++_txId;
        _txes[_txId] = Tx({
            _creator: msg.sender,
            _nftAddr: nftAddress,
            _nftTokenId: nftTokenId,
            _startTime: block.timestamp,
            _lockTime: lockTime,
            highestUSD: startPrice,
            _highestBidder: address(0),
            _tokenAddr: address(0),
            _amount: 0,
            _ended: false
        });
    }

    function getUSD(
        uint256 amount,
        address tokenAddr
    ) internal view returns (int256) {
        AggregatorV3Interface priceFeed = _priceFeeds[tokenAddr];
        // prettier-ignore
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();
        // TODO overflow check
        return int256(amount) * answer;
    }

    // 拍卖
    // TODO
    // 1.   支持ERC20
    // 2.   汇率转换
    function bid(
        uint256 aucId,
        address tokenAddr,
        uint256 amount
    ) external payable {
        Tx storage curTx = _txes[aucId];
        // 是否有效拍卖
        require(curTx._nftAddr != address(0), "Unknown Aucion");
        // 拍卖是否结束
        require(
            block.timestamp < curTx._startTime + curTx._lockTime,
            "Auction Ended"
        );
        //是否是支持的货币类型
        require(
            address(_priceFeeds[tokenAddr]) != address(0),
            "Invalid Currency"
        );
        if (tokenAddr == address(0)) {
            amount = msg.value;
        }
        // 检查出价是否大于最高价
        int256 payValue = getUSD(amount, tokenAddr);
        require(payValue > curTx.highestUSD, "Low Bid");

        // 转移token 到拍卖账户
        if (tokenAddr != address(0)) {
            IERC20(tokenAddr).transferFrom(msg.sender, address(this), amount);
        }
        // 退还之前的ETH或者token
        if (curTx._tokenAddr != address(0)) {
            // 退回token
            IERC20(curTx._tokenAddr).transfer(
                curTx._highestBidder,
                curTx._amount
            );
        } else {
            payable(curTx._highestBidder).transfer(curTx._amount);
        }
        curTx._highestBidder = msg.sender;
        curTx.highestUSD = payValue;
        curTx._tokenAddr = tokenAddr;
        curTx._amount = amount;
    }

    // 结束拍卖
    function end(uint256 aucId) external {
        Tx storage curTx = _txes[aucId];
        require(msg.sender == curTx._creator, "Not allowed");
        require(
            block.timestamp >= curTx._startTime + curTx._lockTime,
            "Auction not ended!"
        );
        require(curTx._ended == false, "The auction had been ended!");
        curTx._ended = true;
        require(curTx._highestBidder != address(0), "No one bids");
        // 转移Token
        if (curTx._tokenAddr != address(0)) {
            IERC20(curTx._tokenAddr).transfer(msg.sender, curTx._amount);
        }
        //转移nft
        address nft_owner = MyNFT(curTx._nftAddr).ownerOf(curTx._nftTokenId);
        MyNFT(curTx._nftAddr).safeTransferFrom(
            nft_owner,
            curTx._highestBidder,
            curTx._nftTokenId
        );
    }

    // 查看最高出价者
    function highestBidAddr(uint256 id) external view returns (address) {
        return _txes[id]._highestBidder;
    }

    // 查看最高价
    function highestBidAmount(uint256 id) external view returns (uint256) {
        return _txes[id]._amount;
    }
}
