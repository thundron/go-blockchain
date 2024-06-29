package main

type Block struct {
	Timestamp int64 // creation time of block
	PreviousBlockHash []byte // hash of prev block
	CurrentBlockHash []byte // hash of the current block
	AllData []byte // transaction data (body info)
}

type Blockchain struct {
	Blocks []*Block // blockchain = series of blocks
}
