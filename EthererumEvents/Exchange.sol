pragma solidity ^0.5.5;

contract Exchange {
    event Fill(
        address indexed maker,
        address indexed feeRecipientAddress,
        address takerAddress,
        address senderAddress,
        uint256 makerAssetFilledAmount,
        uint256 takerAssetFilledAmount,
        uint256 makerFeePaid,
        uint256 TakerFeePaid,
        bytes32 indexed orderHash,
        bytes makerAssetData,
        bytes takerAssetData

    );
    
    event Cancel(
        address indexed makerAddress,
        address indexed feeRecipientAddress,
        address senderAddress,
        bytes32 indexed orderHash,
        bytes makerAssetData,
        bytes takerAssetData
    );

}
