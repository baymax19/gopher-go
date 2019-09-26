package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	
	"./exchange"
)



type LogError struct {
	ErrorId   uint8
	OrderHash [32]byte
}

func main() {
	fmt.Println("Ox Exchange SubScription")
	
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	
	contractAddress := common.HexToAddress("0x48bacb9266a570d521063ef5dd96e61686dbe788")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1401),
		ToBlock:   big.NewInt(1401),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	
	query1 := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}
	
	logS := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query1, logS)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("logs ", logS, sub)
	contractAbi, err := abi.JSON(strings.NewReader(exchange.ExchangeABI))
	if err != nil {
		log.Fatal(err)
	}
	
	logFillSign := []byte("Fill(address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes,bytes)")
	logFillHash := crypto.Keccak256Hash(logFillSign)
	
	logCancelSign := []byte("Cancel(address,address,address,bytes32,bytes,bytes)")
	logCancelHash := crypto.Keccak256Hash(logCancelSign)
	fmt.Println(logCancelHash.Hex())
	
	logErrorSign := []byte("Error(uint8,bytes32")
	logErrorHash := crypto.Keccak256Hash(logErrorSign)
	
	for _, vLog := range logs {
		
		switch vLog.Topics[0].Hex() {
		case logFillHash.Hex():
			
			fmt.Println("Log Name: LogFill\n")
			var fillEvent LogFill
			
			err := contractAbi.Unpack(&fillEvent, "Fill", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			
			fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.FeeRecipientAddress = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.OrderHash = vLog.Topics[3]
			
			fmt.Printf("Maker Address: %s\n", fillEvent.Maker.Hex())
			fmt.Printf("Taker Address: %s\n", fillEvent.TakerAddress.Hex())
			fmt.Printf("Fee Recipient Address : %s\n", fillEvent.FeeRecipientAddress.Hex())
			fmt.Printf("Sender Address: %s\n", fillEvent.SenderAddress.Hex())
			fmt.Printf("Maker Asset Data: %s\n", hexutil.Encode(fillEvent.MakerAssetData))
			fmt.Printf("Taker AssetData: %s\n", hexutil.Encode(fillEvent.TakerAssetData))
			fmt.Printf("Maker Asset Filled Amount: %s\n", fillEvent.MakerAssetFilledAmount.String())
			fmt.Printf("Taker Asset Filled  Amount: %s\n", fillEvent.TakerAssetFilledAmount.String())
			fmt.Printf("Maker Fee Paid: %s\n", fillEvent.MakerFeePaid.String())
			fmt.Printf("Taker Fee Paid : %s\n", fillEvent.TakerFeePaid.String())
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(fillEvent.OrderHash[:]))
		
		case logCancelHash.Hex():
			
			fmt.Printf("Log Name: LogCancel\n")
			var cancelEvent LogCancel
			
			err := contractAbi.Unpack(&cancelEvent, "Cancel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			
			cancelEvent.MakerAddress = common.HexToAddress(vLog.Topics[1].Hex())
			cancelEvent.FeeRecipientAddress = common.HexToAddress(vLog.Topics[2].Hex())
			cancelEvent.OrderHash = vLog.Topics[3]
			
			fmt.Printf("Maker: %s\n", cancelEvent.MakerAddress.Hex())
			fmt.Printf("Fee Recipient: %s\n", cancelEvent.FeeRecipientAddress.Hex())
			fmt.Printf("Maker Asset Data: %s\n", string(cancelEvent.MakerAssetData))
			fmt.Printf("Sender Address: %s\n", cancelEvent.SenderAddress.Hex())
			fmt.Printf("Taker Asset Data: %s\n", string(cancelEvent.TakerAssetData))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))
		
		case logErrorHash.Hex():
			fmt.Printf("Log Name: LogError\n")
			
			errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			
			errorEvent := &LogError{
				ErrorId:   uint8(errorID),
				OrderHash: vLog.Topics[2],
			}
			
			fmt.Printf("Error ID: %d\n", errorEvent.ErrorId)
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))
		}
		
		fmt.Printf("\n\n")
		
	}
}
