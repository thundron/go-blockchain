package main

import (
	"fmt"
)

func main() {
	blockchain := CreateBlockchain() // init the blockchain
	blockchain.AddBlock("first transaction") // first block that represents a transaction
	blockchain.AddBlock("second transaction") // second block that represents a transaction
	for i, block := range blockchain.Blocks {
		fmt.Prinf("Block ID: %d\n", i)
		fmt.Prinf("Timestamp: %d\n", block.Timestamp+int64(i)) // differentiate the block timestamps with the ID
		fmt.Printf("Hash of the %d block: %x\n", i, block.CurrentBlockHash)
		fmt.Printf("Hash of the previous %d block: %x\n", i-1, block.PreviousBlockHash)
		fmt.Printf("All transactions: %s\n", block.AllData)
	}
}
