// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract Math {
    error Overflow();
    error Underflow();
    error DivisionByZero();

    function sqrt(uint256 a) public pure returns (uint256) {
        // binary search
        if (a == 0) {
            return 0;
        }
        uint256 low = 1;
        uint256 hi = a;
        while (low <= hi) {
            uint256 mid = low + (hi - low) / 2;
            if (mid == a / mid) {
                return mid;
            } else if (mid > a / mid) {
                hi = mid - 1;
            } else {
                low = low + 1;
            }
        }
        return hi;
    }

    // function add(uint256 a, uint256 b) public pure returns (uint256) {
    //     uint256 c = a + b;
    //     if (c < a) {
    //         revert Overflow();
    //     }
    //     return c;
    // }
    //
    // function sub(uint256 a, uint256 b) public pure returns (uint256) {
    //     if (b > a) {
    //         revert Underflow();
    //     }
    //     return a - b;
    // }
    //
    // function mul(uint256 a, uint256 b) public pure returns (uint256) {
    //     uint256 c = a * b;
    //     if (c / a != b) {
    //         revert Overflow();
    //     }
    //
    //     return c;
    // }
    //
    // function div(uint256 a, uint256 b) public pure returns (uint256) {
    //     if (b == 0) {
    //         revert DivisionByZero();
    //     }
    //     return a / b;
    // }
}
