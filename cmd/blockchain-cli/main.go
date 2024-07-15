package main

import (
	"fmt"
	"blockchain/internal/blockchain"
)

func main() {
	// Initialize a new blockchain
	blockchain := blockchain.Blockchain{}
	blockchain.AddBlock("Block 1 Data")
	blockchain.AddBlock("Block 2 Data")
	blockchain.AddBlock("Block 3 Data")

	// Print out the blockchain
	for _, block := range blockchain.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("--------------------")
	}
}
