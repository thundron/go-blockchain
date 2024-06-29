package main

func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // every time we add a block besides the GenesisBlock, we need to provide the previous block
	newBlock := CreateBlock(data, PreviousBlock.CurrentBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock) // append the new block to the blockchain
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}} // create a new blockchain and add the genesis block to it
}

