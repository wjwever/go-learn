//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

//import "@openzeppelin/contracts/utils/Counters.sol";

contract MyNFT is ERC721, ERC721URIStorage {
    address private _owner;
    uint256 public _counter;
    string public constant URI =
        "https://bafybeiaal6hb63cqbnedvwderv6s7ny3thsdiu6slqabgv3ldcwliekeea.ipfs.dweb.link?filename=my_nft_meta.json";

    constructor() ERC721("MyNFT", "MyNFT") {
        _owner = msg.sender;
        _counter = 0;
    }

    function safeMint(
        address recipient
    )
        external
        returns (
            // string memory token_uri
            uint256
        )
    {
        ++_counter;
        _mint(recipient, _counter);
        _setTokenURI(_counter, URI);
        return _counter;
    }

    function tokenURI(
        uint256 tokenId
    ) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
