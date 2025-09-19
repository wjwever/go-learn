// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BinarySearch {
    function search(int256[] memory data, int256 target) public pure returns(int256) {
        uint256 len = data.length;
        int256 left = 0;
        int256 right = int256(len) -1;

        while(left <= right) {
            int256 mid =  (left + right) >> 1;
            if (data[uint256(mid)] == target) {
                return mid;
            } else if (data[uint256(mid)] > target) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        return -1;
    }
}
