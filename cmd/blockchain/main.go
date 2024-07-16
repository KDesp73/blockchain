package main

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/wallet"
	"fmt"
)

func printBlockchain(bc *blockchain.Blockchain) {
	println("")
	for _, block := range bc.Chain {
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Transaction: [Amount: %d, From: ..., To: ...]\n", block.Transaction.Amount)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("--------------------")
	}
}

func main() {
	bc := blockchain.NewBlockchain();

	kdesp73 := wallet.NewWallet("KDesp73");
	thanasisgeorg := wallet.NewWallet("ThanasisGeorg");
	creatorkostas := wallet.NewWallet("creatorkostas");

	kdesp73.SendMoney(bc, 10, thanasisgeorg.PublicKey)
	thanasisgeorg.SendMoney(bc, 20, creatorkostas.PublicKey)
	creatorkostas.SendMoney(bc, 30, kdesp73.PublicKey)

	printBlockchain(bc)
}
