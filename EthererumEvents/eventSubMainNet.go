package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	
	// "strings"
	
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	
	"github.com/baymax19/gopher-go/EthererumEvents/exchange"
)

type LogFill struct {
	Maker                  common.Address
	FeeRecipientAddress    common.Address
	TakerAddress           common.Address
	SenderAddress          common.Address
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
	OrderHash              [32]byte
	MakerAssetData         []byte
	TakerAssetData         []byte
}

type LogCancel struct {
	MakerAddress        common.Address
	FeeRecipientAddress common.Address
	SenderAddress       common.Address
	OrderHash           [32]byte
	MakerAssetData      []byte
	TakerAssetData      []byte
}

func main() {
	
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	
	contractAddress := common.HexToAddress("0x4F833a24e1f95D70F028921e27040Ca56E09AB0b")
	if err != nil {
		log.Fatal(err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	
	contractAbi, err := abi.JSON(strings.NewReader(exchange.ExchangeABI))
	if err != nil {
		log.Fatal(err)
	}
	
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	
	logFillSign := []byte("Fill(address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes,bytes)")
	logFillHash := crypto.Keccak256Hash(logFillSign)
	
	logCancelSign := []byte("Cancel(address,address,address,bytes32,bytes,bytes)")
	logCancelHash := crypto.Keccak256Hash(logCancelSign)
	
	for {
		select {
		case <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.TxHash.Hex())
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
				fmt.Printf("Maker Asset Data: %s\n", hexutil.Encode(cancelEvent.MakerAssetData))
				fmt.Printf("Sender Address: %s\n", cancelEvent.SenderAddress.Hex())
				fmt.Printf("Taker Asset Data: %s\n", hexutil.Encode(cancelEvent.TakerAssetData))
				fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))
			}
			fmt.Println()
		}
	}
}
