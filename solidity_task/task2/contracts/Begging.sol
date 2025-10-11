//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "hardhat/console.sol";

contract Begging {
    address _owner;
    mapping (address => uint256) _donations;
    address[] _addrs;
    bool _widthdrawed; // 记录捐赠是否提取

    // 日期锁
    uint256 _lockTime; // 锁定期
    uint256 _startTime; 

    // 捐赠事件
    event Donate(address sender, uint256 amount);
    // 提取事件
    event Withdraw(address reciver, uint256 amount);

    // 构造函数
    constructor(uint256 lockTime) {
        _owner = msg.sender; 
        _widthdrawed = false;
        _startTime = block.timestamp;
        _lockTime = lockTime;
    }

    //允许用户向合约发送以太币，并记录捐赠信息。 
    function donate() external payable {
        require (block.timestamp  < _startTime + _lockTime, "The donation ended now");
        if (_donations[msg.sender] == 0) {
            _addrs.push(msg.sender);
        }
        _donations[msg.sender] += msg.value;
        emit Donate(msg.sender, msg.value);
    }

    //合约所有者提取所有资金
    function withdraw() public onlyOwner {
        // 过了捐款期才可以提款
        require (block.timestamp  > _startTime + _lockTime, "Locked now");

        // 不能重复提取
        require(_widthdrawed == false, "You have noting to withdraw");

        _widthdrawed = true;
        uint256 balance = address(this).balance;
        payable(msg.sender).transfer(balance); 
        emit Withdraw(msg.sender, balance);
    }

    // 查询某个地址的捐赠金额
    function getDonation(address addr) public view returns(uint256) {
        return _donations[addr];
    }

    // 查询top3捐赠者
    function getTop3Donors() public view returns(address[3] memory) {
        address[3]  memory addrs;
        uint256[3]  memory donations;
        for (uint256 i = 0; i < 3; i++) {
            addrs[i] = address(0);
            donations[i] = 0;
        }
        for (uint256 i = 0 ; i < _addrs.length; i++) {
            address cur = _addrs[i];
            uint256 amount = _donations[cur];
            for (uint256 j = 0; j < 3; j++) {
                if (amount > donations[j]) {
                    for (uint256 k = 2; k > j; k--) {
                        console.log("move ", k - 1, "->", k);
                        donations[k] = donations[k - 1];
                        addrs[k] = addrs[k - 1];
                    }
                    addrs[j] = cur;
                    donations[j] = amount;
                    console.log("put " , amount, "to ", j);
                    break;
                }
            }
        }
        return  addrs;
    }

    function owner() external view returns(address) {
        return _owner;
    }

    function canDonate() external view returns(bool) {
        return block.timestamp < _startTime + _lockTime;
    }

    modifier  onlyOwner {
        require(msg.sender == _owner, "Only owner can do this opertion");
        _;
    }
}
