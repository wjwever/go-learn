// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

contract SafeMath {
    function sqrt(uint256 a) public pure returns (uint256) {
        uint256 c = Math.sqrt(a);
        return c;
    }
}
