package main

import (
	"encoding/hex"
	"fmt"
	
	"github.com/ethereum/go-ethereum/common"
	amino "github.com/tendermint/go-amino"
	// "golang.org/x/crypto/sha3"
)

type Data struct {
	Order string `json:"order"`
	// Add   string  `json:"add"`
}

const AddressLength = 40

type Address [AddressLength]byte

func main() {
	cdc := amino.NewCodec()
	fmt.Println(cdc)
	// address := common.HexToAddress("0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	bytes := common.FromHex("0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	b1 := BytesToAddress(bytes)
	// fmt.Println("address", address.String(), "bytes", string(bytes))
	fmt.Println(b1.Hex())
	order := Data{
		Order: "details",
	}
	bz, err := cdc.MarshalJSON(order)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", string(bz))
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

// Hex returns an EIP55-compliant hex string representation of the address.
func (a Address) Hex() string {
	unchecksummed := hex.EncodeToString(a[:])
	
	return "0x" + unchecksummed
}

// String implements fmt.Stringer.
func (a Address) String() string {
	return a.Hex()
}
