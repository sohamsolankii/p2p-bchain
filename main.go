package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain represents a chain of blocks
type BlockChain struct {
	blocks []*Block // Slice of pointers to Block
}

// Block represents each block in the blockchain
type Block struct {
	Hash     []byte // Hash of the current block
	Data     []byte // Data stored in the block
	PrevHash []byte // Hash of the previous block
}

// DeriveHash calculates and sets the hash for the block
func (b *Block) DeriveHash() {
	// Combine the block's data and previous hash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Compute the SHA-256 hash
	hash := sha256.Sum256(info)
	// Set the block's hash
	b.Hash = hash[:]
}

// CreateBlock creates a new block with the given data and previous hash
func CreateBlock(data string, prevHash []byte) *Block {
	// Initialize a new block
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Derive the hash for the new block
	block.DeriveHash()
	return block
}

// AddBlock adds a new block to the blockchain with the given data
func (chain *BlockChain) AddBlock(data string) {
	// Get the last block in the chain
	prevBlock := chain.blocks[len(chain.blocks)-1]
	// Create a new block with the previous block's hash
	new := CreateBlock(data, prevBlock.Hash)
	// Append the new block to the chain
	chain.blocks = append(chain.blocks, new)
}

// Genesis creates the first block in the blockchain (the genesis block)
func Genesis() *Block {
	// The genesis block has no previous hash
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with the genesis block
func InitBlockChain() *BlockChain {
	// Create a blockchain with the genesis block as the first block
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	// Initialize the blockchain
	chain := InitBlockChain()

	// Add blocks to the blockchain
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// Iterate through the blocks and print their details
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash) // Print the previous block's hash
		fmt.Printf("Data in Block: %s\n", block.Data)    // Print the block's data
		fmt.Printf("Hash: %x\n", block.Hash)             // Print the block's hash
	}
}