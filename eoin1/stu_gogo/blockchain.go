package main

/*
用go进行区块链开发
 */
type Block struct {
	Timestame     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
