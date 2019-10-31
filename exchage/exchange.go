package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

func main(){

	takerAsset := types.NewInt64Coin("usd",2)
	makerAsset := types.NewInt64Coin("stake",1)
	takerFillAmount := types.NewInt64Coin("usd",3)

	takerGetAmount := makerAsset.Amount.ToDec().Mul(takerFillAmount.Amount.ToDec()).Quo(takerAsset.Amount.ToDec())
	fmt.Println(takerGetAmount)

	remain := takerFillAmount.Amount.ToDec().Sub(takerGetAmount)
	fmt.Println(remain,remain.TruncateInt64())

}
