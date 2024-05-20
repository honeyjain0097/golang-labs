package readcontract

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadWhitelistOwner() {
	// Connect to an Ethereum node
	const providerUrl = "https://rpc.escscan.com"
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Load the ABI of the smart contract
	contractAddress := common.HexToAddress("0x9D0b3C95dfc5D6913679E9D3B11Cc3A294a35a6B")
	data, err := os.ReadFile("./lib/readcontract/whitelist_abi.json")
	if err != nil {
		fmt.Printf("Error in read contract abi ")
		fmt.Print(err)
	}
	abiInstance, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("Error in abiInstance")
		log.Fatal(err)
	}

	// Call the balanceOf function
	data, err = abiInstance.Pack("owner")
	if err != nil {
		fmt.Printf("Error in pack")
		log.Fatal(err)
	}
	data, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		fmt.Printf("Error in call contract")
		log.Fatal(err)
	}

	// Decode the result
	unPackData, err := abiInstance.Unpack("owner", data)
	if err != nil {
		fmt.Printf("Error in Unpack")
		log.Fatal(err)
	}

	fmt.Printf("whitelist owner : %t\n", unPackData)
}
func ReadIsWhiteList() {
	// Connect to an Ethereum node
	const providerUrl = "https://rpc.escscan.com"
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Load the ABI of the smart contract
	contractAddress := common.HexToAddress("0x9D0b3C95dfc5D6913679E9D3B11Cc3A294a35a6B")
	data, err := os.ReadFile("./lib/readcontract/whitelist_abi.json")
	if err != nil {
		fmt.Printf("Error in read contract abi ")
		fmt.Print(err)
	}
	abiInstance, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("Error in abiInstance")
		log.Fatal(err)
	}

	// Call the balanceOf function
	userAddress := common.HexToAddress("0x764D3842e10f2D9Eda9e6b90c08D002d7D26E7c5")
	data, err = abiInstance.Pack("isWhitelisted", userAddress)
	if err != nil {
		fmt.Printf("Error in pack")
		log.Fatal(err)
	}
	data, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		fmt.Printf("Error in call contract")
		log.Fatal(err)
	}

	// Decode the result
	unPackData, err := abiInstance.Unpack("isWhitelisted", data)
	if err != nil {
		fmt.Printf("Error in Unpack")
		log.Fatal(err)
	}

	fmt.Printf("Is whitelisted: %t\n", unPackData)
}
func hexStringToBytes(hexString string) ([]byte, error) {
	// Remove any non-hex characters (like '0x' prefix)
	hexString = strings.TrimPrefix(hexString, "0x")
	hexString = strings.Replace(hexString, " ", "", -1) // Remove spaces if any

	// Decode the hexadecimal string
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
func ReadIsBridgeIsTxnFilled() {
	// Connect to an Ethereum node
	const providerUrl = "https://rpc.escscan.com"
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Load the ABI of the smart contract
	contractAddress := common.HexToAddress("0x96a9d989fE46220c1124f12384571aB6fdEb0B1E")
	data, err := os.ReadFile("./lib/readcontract/bridge_abi.json")
	if err != nil {
		fmt.Printf("Error in read contract abi ")
		fmt.Print(err)
	}
	abiInstance, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("Error in abiInstance")
		log.Fatal(err)
	}

	// Call the balanceOf function
	txn_hash := "0xcd66897bf618ee0d522ec4f915abe250151db465122feeefec06cac255456293"
	// param := common.HexToHash(txn_hash)

	txnHashBytes, err := hexStringToBytes(txn_hash)
	if err != nil {
		fmt.Println("Error decoding transaction hash:", err)
		return
	}
	var txnHashArray [32]byte
	copy(txnHashArray[:], txnHashBytes)

	data, err = abiInstance.Pack("isTxnFilled", txnHashArray)
	if err != nil {
		fmt.Printf("Error in pack")
		log.Fatal(err)
	}
	data, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		fmt.Printf("Error in call contract")
		log.Fatal(err)
	}
	// Check if result data is empty
	if len(data) == 0 {
		fmt.Println("Empty result returned from contract function")
	}

	fmt.Printf("data :: %b", data)

	// Decode the result
	unPackData, err := abiInstance.Unpack("isTxnFilled", data)
	if err != nil {
		fmt.Printf("Error in Unpack")
		log.Fatal(err)
	}

	fmt.Printf("Is isTxnFilled: %t\n", unPackData)
}

func ReadBridgeOwner() {
	// Connect to an Ethereum node
	const providerUrl = "https://rpc.escscan.com"
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Load the ABI of the smart contract
	contractAddress := common.HexToAddress("0x96a9d989fE46220c1124f12384571aB6fdEb0B1E")
	data, err := os.ReadFile("./lib/readcontract/bridge_abi.json")
	if err != nil {
		fmt.Printf("Error in read contract abi ")
		fmt.Print(err)
	}
	abiInstance, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("Error in abiInstance")
		log.Fatal(err)
	}

	// Call the balanceOf function
	data, err = abiInstance.Pack("owner")
	if err != nil {
		fmt.Printf("Error in pack")
		log.Fatal(err)
	}
	data, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		fmt.Printf("Error in call contract")
		log.Fatal(err)
	}

	// Decode the result
	unPackData, err := abiInstance.Unpack("owner", data)
	if err != nil {
		fmt.Printf("Error in Unpack")
		log.Fatal(err)
	}

	fmt.Printf("whitelist owner : %t\n", unPackData)
}
