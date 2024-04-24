package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index)
	bc.blocks = append(bc.blocks, newBlock)
	return newBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, 0)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
