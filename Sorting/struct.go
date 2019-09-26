package main

import (
	"fmt"
)

type Two struct {
	Name2 string
}
type One struct {
	Two
	Name string
}

func main() {
	var data1 One
	data1.Name2 = "cjecl0"
	data1.Two.Name2 = "saa"
	fmt.Println(data1)
}
