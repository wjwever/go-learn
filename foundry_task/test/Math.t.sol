// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {Math} from "../src/Math.sol";
import {console} from "forge-std/console.sol"; // 必须导入

contract MathTest is Test {
    Math public math;

    function setUp() public {
        math = new Math();
    }

    function test(uint256 val) public view {
        vm.assume(val < 10 ** 9);
        uint256 c = math.sqrt(val);
        assertLe(c * c, val);
        assertGt((c + 1) * (c + 1), val);
    }
}
