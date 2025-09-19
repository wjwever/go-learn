// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//实现罗马数字转数整数, 用solidity实现
contract RomanToInteger {
    // 罗马字符到数字的映射
    mapping(bytes1 => uint256)  romanValues;
    
    constructor() {
        // 初始化罗马数字映射
        romanValues['I'] = 1;
        romanValues['V'] = 5;
        romanValues['X'] = 10;
        romanValues['L'] = 50;
        romanValues['C'] = 100;
        romanValues['D'] = 500;
        romanValues['M'] = 1000;
    }
    
    function romanToInt(string memory s) public view returns (uint256) {
        
        // 将字符串转换为bytes以便遍历
        bytes memory roman = bytes(s);

        uint256 sum = 0;
        for (uint256 i = 0; i < roman.length; i++) {
            uint256 cur = romanValues[roman[i]];
            sum += cur;
            if ( i  > 0 ) {
                uint256 pre = romanValues[roman[i-1]];
                if (cur > pre) {
                    sum -= pre * 2;  // 注意为负数的问题， 所以前面先进行了加法
                }
            }
        }
        return sum;
        
    }
}
