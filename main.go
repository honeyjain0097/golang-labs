package main

import (
	"fmt"

	"github.com/honeyjain0097/golang-labs/lib/readcontract"
)

func main() {
	fmt.Println("---- MAIN ----")
	readcontract.ReadIsWhiteList()
	readcontract.ReadIsBridgeIsTxnFilled()
}
