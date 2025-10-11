// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//实现罗马数字转数整数, 用solidity实现
contract IntToRoman{
    struct Info {
        uint256 value;
        string symbol;
    }

    Info[] _infos;
    constructor() {
        _infos.push(Info(1000, "M"));
        _infos.push(Info(900, "CM"));
        _infos.push(Info(500, "D"));
        _infos.push(Info(400, "CD"));
        _infos.push(Info(100, "C"));
        _infos.push(Info(90, "XC"));
        _infos.push(Info(50, "L"));
        _infos.push(Info(40, "XL"));
        _infos.push(Info(10, "X"));
        _infos.push(Info(9, "IX"));
        _infos.push(Info(5, "V"));
        _infos.push(Info(4, "IV"));
        _infos.push(Info(1, "I"));
    }

    function intToRoman(uint256 num) public view returns(string memory) {
        string memory ret = "";
        while ( num != 0) {
            for (uint256 i = 0; i < _infos.length; i++) {
                if (num >= _infos[i].value) {
                    ret = string.concat(ret, _infos[i].symbol);
                    num -= _infos[i].value;
                    break;
                }
            }
        }
        return ret;
    }
}

