package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
	"sort"
)

type Order struct {
	Name   string
	Value1 *sdk.Int
	Value2 *sdk.Int
}

type Orders []Order

var _ sort.Interface = Orders{}

func main() {
	a := sdk.NewUint(1124356712314901199)
	b := sdk.NewUint(1234567891234512349)
	i := new(sdk.Int)
	divMul := i.Mul(a, b)
	fmt.Println(i, divMul)
	i.Div(i, sdk.NewUint(1234567891234512349))
	fmt.Println(divMul)
	var datas = []Order{
		//        { Name :  "test1",  Value1 : 12, Value2 : 13 },
		//        { Name :  "testt2",  Value1 : 10, Value2 : 14 },
		{Name: "test1", Value1: sdk.NewUint(1012789378078089989), Value2: sdk.NewUint(13)},
		{Name: "testt2", Value1: divMul, Value2: sdk.NewUint(14)},
	}
	fmt.Sprintf(`Data : %d`, i)
	sort.Sort(Orders(datas))
	fmt.Println(datas, i)

}

func (orders Orders) Len() int { return len(orders) }

func (orders Orders) Less(i, j int) bool { return orders[i].Value1.LT(orders[j].Value1) }

func (orders Orders) Swap(i, j int) { orders[i], orders[j] = orders[j], orders[i] }
