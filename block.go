package main

import (
	"bytes" // byte conversion lib to send over network
	"crypto/sha256" // crypto lib to hash data
	"strconv" // string conversion lib
	"time" // time management lib
)

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10)) // get time and convert into a unique series
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concat block data
	hash := sha256.Sum256(headers) // create block hash
	block.CurrentBlockHash = hash[:] // set curr block hash
}

func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // block received
	block.SetHash() // hash the block
	return block
}

// genesis block = first block of the blockchain
func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

