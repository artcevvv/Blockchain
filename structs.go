package main

type Block struct {
	Time          int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type Blockchain struct {
	blocks []*Block
}
