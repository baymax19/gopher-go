// package main
//
// import (
// 	"context"
// 	// "encoding/json"
// 	"fmt"
// 	"log"
// 	// "math/big"
// 	// "strings"
//
// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
//
// 	// "github.com/ethereum/go-ethereum/accounts/abi"
// 	//            "github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
//
// 	// "./exchange"
//
// 	// "math/big"
// )
//
// // type LogFill struct {
// // 	Maker                  common.Address
// // 	Taker                  common.Address
// // 	FeeRecipient           common.Address
// // 	MakerToken             common.Address
// // 	TakerToken             common.Address
// // 	FilledMakerTokenAmount *big.Int
// // 	FilledTakerTokenAmount *big.Int
// // 	PaidMakerFee           *big.Int
// // 	PaidTakerFee           *big.Int
// // 	Tokens                 [32]byte
// // 	OrderHash              [32]byte
// // }
//
// func main() {
// 	fmt.Println("Event SubScription of Ethereum")
//
// 	client, err := ethclient.Dial("ws://localhost:8545")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// Signed LogFillOrder := 0x0bcc4c97732e47d9946f229edb95f5b6323f601300e4690de719993f3c371129
// 	// byte32 data 0xc711353550a0a496d40f6bb265741d01e7d6917ffa4eb1b79a0c4ab666b5eb96 (Hello World)
// 	contractAddress := common.HexToAddress("0x177b2f123b0c6449a2704ef31d7435d0bf14b41b")
// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{contractAddress},
// 	}
//
// 	// contractAbi, err := abi.JSON(strings.NewReader(string(exchange.StoreABI)))
// 	logs := make(chan types.Log)
// 	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
//
//
//
// 	// var event LogFill
//
// 	// for _, vLog := range logs {
// 	// 	err := contractAbi.Unpack(&event, "LogFill", vLog.Data)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	bz, err := json.Marshal(vLog)
// 	// 	fmt.Println(string(bz))
// 	// 	fmt.Println(vLog.TxHash.Hex())
// 	//
// 	// }
//
// 	// for {
// 	// 	select {
// 	// 	case err := <-sub.Err():
// 	// 		log.Fatal(err)
// 	// 	case vLog := <-logs:
// 	// 		err := contractAbi.Unpack(&event, "LogFill", vLog.Data)
// 	// 		if err != nil {
// 	// 			log.Fatal(err)
// 	// 		}
// 	// 		bz, err := json.Marshal(vLog)
// 	// 		fmt.Println(string(bz))
// 	// 		fmt.Println(vLog.TxHash.Hex())
// 	//
// 	// 	}
// 	// }
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:
// 			fmt.Println(vLog) // pointer to event log
// 		}
// 	}
// }

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	
	contractAddress := common.HexToAddress("0xbdacb395a6e516f4462befaebf490a76ed529b59")
	query := ethereum.FilterQuery{
		ToBlock:   big.NewInt(1344),
		FromBlock: big.NewInt(1344),
		Addresses: []common.Address{contractAddress},
	}
	query1 := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	
	logsQuery, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, vLog := range logsQuery {
		bz, _ := json.Marshal(vLog)
		fmt.Println(string(bz))
		fmt.Println(vLog.TxHash.Hex())
	}
	
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query1, logs)
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
	//
}
