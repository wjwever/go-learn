//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract MyToken is IERC20 {
    string private _name;
    string private _symbol;
    address private _owner;
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) _allowances;
    uint256 private _totalSupply;

    // modifier
    modifier onlyOwner() {
        require(msg.sender == _owner, "Only owner can do this opertion");
        _;
    }

    // 错误
    error ERC20InvalidReceiver(address receiver);
    error ERC20InsufficientBalance(
        address owner,
        uint256 balance,
        uint256 needed
    );
    error ERC20InsufficientAllowance(
        address owner,
        address spender,
        uint256 allowance,
        uint256 needed
    );

    /// functions
    constructor(string memory name, string memory symbol) {
        _name = name;
        _symbol = symbol;
        _owner = msg.sender;
    }

    // 合约所有者增发代币
    function mint(uint256 value) external onlyOwner {
        _totalSupply += value;
        _balances[_owner] += value;
    }

    // 查询总代币供应量
    function totalSupply() external view returns (uint256) {
        return _totalSupply;
    }

    // 查询账户余额
    function balanceOf(address account) external view returns (uint256) {
        return _balances[account];
    }

    // 转账
    function transfer(address to, uint256 value) external returns (bool) {
        address from = msg.sender;
        uint256 balance = _balances[from];
        if (balance < value) {
            revert ERC20InsufficientBalance(from, balance, value);
        }
        _balances[from] -= value;
        _balances[to] += value;
        emit Transfer(from, to, value);
        return true;
    }

    function allowance(
        address owner,
        address spender
    ) external view returns (uint256) {
        return _allowances[owner][spender];
    }

    // 授权
    function approve(address spender, uint256 value) external returns (bool) {
        if (spender == address(0)) {
            revert ERC20InvalidReceiver(spender);
        }

        address owner = msg.sender;
        uint256 balance = _balances[owner];
        if (balance < value) {
            revert ERC20InsufficientBalance(owner, balance, value);
        }
        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
        return true;
    }

    // 代扣转账
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external returns (bool) {
        address spender = msg.sender;
        uint256 allow = _allowances[from][spender];
        if (allow < value) {
            revert ERC20InsufficientAllowance(from, spender, allow, value);
        }

        uint256 balance = _balances[from];
        if (balance < value) {
            revert ERC20InsufficientBalance(from, balance, value);
        }

        _balances[from] -= value;
        _balances[to] += value;
        _allowances[from][spender] -= value;
        emit Transfer(from, to, value);
        return true;
    }
}
