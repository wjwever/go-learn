// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Voting {
    mapping(string  =>uint256 ) _votes;
    string [] _candidates;
    address _owner;

    constructor() {
        _owner = msg.sender;
    }

    function vote(string memory candi) public {
        if (_votes[candi] == 0) {
            _candidates.push(candi);
        }
        _votes[candi]++;
    }

    function getVotes(string calldata candi) public view  returns(uint256){
        return _votes[candi];
    }

    function resetVotes() public {
        require(msg.sender == _owner, "Pemission Denied");
        for (uint256 i = 0; i < _candidates.length; i++) {
            _votes[_candidates[i]] = 0;
        }
    }
}
