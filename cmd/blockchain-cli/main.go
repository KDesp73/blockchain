package main

import (
	"fmt"
	"blockchain/internal/blockchain"
	"time"
)

func printBlockchain(bc *blockchain.Blockchain) {
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("--------------------")
	}

}

func main() {
	bc := &blockchain.Blockchain{
		Chain: []blockchain.Block{
			{
				Index:     0,
				Timestamp: time.Now().String(),
				Data:      "Genesis Block",
				PrevHash:  "",
				Hash:      "",
			},
		},
	}

	bc.AddBlock("Block 1 Data")
	bc.AddBlock("Block 2 Data")
	bc.AddBlock("Block 3 Data")

	printBlockchain(bc)
}
