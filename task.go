package blockchain

import (
    "fmt"
)

// Block represents a block in the blockchain
type Block struct {
    Data string
}

// DisplayAllBlocks displays all blocks in the blockchain
func DisplayAllBlocks(blocks []*Block) {
    for i, block := range blocks {
        fmt.Printf("Block %d: %s\n", i+1, block.Data)
    }
}

// NewBlock creates a new block with the given data
func NewBlock(data string) *Block {
    return &Block{Data: data}
}

// ModifyBlock modifies the data of a given block
func ModifyBlock(block *Block, newData string) {
    block.Data = newData
}
