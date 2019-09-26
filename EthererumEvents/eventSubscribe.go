package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
	
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		select {
		case <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			fmt.Println(vlog.TxHash.Hex())
			for _, t := range vlog.Topics {
				fmt.Println(t.Hex())
			}
		}
	}
}
