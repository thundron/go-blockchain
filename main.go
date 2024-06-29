package main

import (
	"fmt"
)

func main() {
	blockchain := CreateBlockchain() // init the blockchain
	blockchain.AddBlock("first transaction") // first block that represents a transaction
	blockchain.AddBlock("second transaction") // second block that represents a transaction
	for i, block := range blockchain.Blocks {
		fmt.Printf("Hash of the %d block: %x", i, block.CurrentBlockHash)
		fmt.Printf("Hash of the previous %d block: %x", i-1, block.PreviousBlockHash)
		fmt.Printf("All transactions: %s\n", block.AllData)
	}
}
