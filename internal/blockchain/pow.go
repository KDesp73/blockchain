package blockchain

import (
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rsa"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const Difficulty = 4

func (bc *Blockchain) MineBlock(t transaction.Transaction, publicKey *rsa.PublicKey, signature string) {
	if encryption.VerifySignature(publicKey, []byte(t.ToString()), []byte(signature)) != nil {
		fmt.Println("Could not verify the signature: ", signature)
		return
	}


	fmt.Println("⛏️ Mining...")
	newBlock := generateBlock(*bc.LastBlock(), t)

	start := time.Now()
	for {
		hash := encryption.Hash(newBlock.computeHash() + strconv.Itoa(newBlock.Nonce))
		
		if hash[:Difficulty] == strings.Repeat("0", Difficulty) {
			fmt.Printf("Hash found: %s\n", hash)
			break
		}
		newBlock.Nonce++
	}
	end := time.Now()

	elapsed := end.Sub(start)
	fmt.Println("Time elapsed: ", elapsed)

	bc.AddBlock(newBlock)
}
