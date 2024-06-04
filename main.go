package main

import (
	"fmt"

	"github.com/honeyjain0097/golang-labs/lib/readcontract"
)

func main() {
	fmt.Println("---- MAIN ----")
	// readcontract.ReadIsWhiteList()
	readcontract.ReadIsBridgeIsTxnFilled("https://rpc.escscan.com", "0x96a9d989fE46220c1124f12384571aB6fdEb0B1E")
	// readcontract.ReadBridgeOwner()
	// readcontract.ReadWhitelistOwner()
	// readcontract.ReadOwner("https://rpc-testnet.escscan.com", "0x37Cf8Bbb0358AeFa99580DC904Aea7863a6EE9B6")
	// readcontract.ReadOwner("https://rpc.escscan.com", "0xc929Be511347Cd4133fC22d4eA405fc3161c12A5")
}
