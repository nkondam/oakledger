package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte, prevIdx int) *Block {
	block := &Block{
		Index:     prevIdx + 1,
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		Hash:      []byte{},
		PrevHash:  prevHash,
	}
	block.SetHash()
	return block
}
