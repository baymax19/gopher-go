package main

import (
	"fmt"
)

type Negotiation interface {
	GetData() int64
	SetData(int64) error
	String() string
}
type BaseNegotiation struct {
	Data int64 `json:"data"`
}

var _ Negotiation = &BaseNegotiation{}

func NewBaseNegotitaion(data int64) Negotiation {
	return &BaseNegotiation{
		Data: data,
	}
}

func (b BaseNegotiation) GetData() int64 {
	return b.Data
}

func (b *BaseNegotiation) SetData(d int64) error {
	b.Data = d
	return nil
}

func (b BaseNegotiation) String() string {
	return fmt.Sprintf(`
Data: %d
`, b.Data)
}

type Negotiations []Negotiation

func NewNegotiations(ns ...Negotiation) Negotiations {
	if len(ns) == 0 {
		return Negotiations{}
	}
	return Negotiations(ns)
}

func (ns Negotiations) Add(nsB Negotiations) Negotiations {
	// fmt.Println("nsB", nsB)
	ns = append(ns, nsB...)
	// fmt.Println("totalNS", ns)
	return ns
}

func (ns Negotiations) String() string {

	out := ""
	for _, n := range ns {
		fmt.Println("n is", len(ns), n)
		out += fmt.Sprintf(n.String())
	}
	return out[:]
}

type Collection struct {
	Hash         int64        `json:"hash"`
	Negotiations Negotiations `json:"negotiations"`
}

func main() {
	bn1 := NewBaseNegotitaion(1)
	bn2 := NewBaseNegotitaion(2)

	nns1 := NewNegotiations(bn1)
	nns2 := NewNegotiations(bn2)
	nn := nns1.Add(nns2)
	fmt.Println("nns1", nn)

	// bn3 := NewBaseNegotitaion(3)
	// nns2 := NewNegotiations(bn3)
	// fmt.Println("nns2", nns2.String())

}
