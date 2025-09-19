// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MergeSortedArray {
    function merge(uint256[] memory arr1, uint256[] memory arr2) public pure  returns(uint256[] memory) {
        uint256 len1 = arr1.length;
        uint256 len2 = arr2.length;
        uint256[] memory res = new uint256[](len1 + len2);

        uint256 id1 = 0;
        uint256 id2 = 0;
        uint256 id3 = 0;
        while(id1 < len1 && id2 < len2) {
            if (arr1[id1] < arr2[id2]) {
             res[id3++] = arr1[id1++];
            } else {
             res[id3++] = arr2[id2++];
            }
        }

        while(id1 < len1) {
             res[id3++] = arr1[id1++];
        }

        while (id2 < len2) {
             res[id3++] = arr2[id2++];
        }
        return res;
    }
}
