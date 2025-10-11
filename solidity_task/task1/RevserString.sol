// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


contract ReverseString {
    function reverse(string calldata input) public pure returns(string memory) {
        bytes memory data = bytes(input);
        uint256 len = data.length;
        for (uint256 i = 0; i < len / 2; i++) {
            bytes1 tmp = data[i];
            data[i] = data[len - 1 -i];
            data[len - 1 -i] = tmp;
        }
        return string(data);
    }
}
