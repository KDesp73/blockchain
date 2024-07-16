package blockchain

import (
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rsa"
	"fmt"
	"strings"
)

const difficulty = 4

func (bc *Blockchain) MineBlock(t transaction.Transaction, publicKey *rsa.PublicKey, signature string) {
	if encryption.VerifySignature(publicKey, []byte(t.ToString()), []byte(signature)) != nil {
		fmt.Println("Could not verify the signature: ", signature)
		return
	}

	newBlock := generateBlock(*bc.LastBlock(), t)
	for {
		newBlock.computeHash()
		
		if newBlock.Hash[:difficulty] == strings.Repeat("0", difficulty) {
			break
		}
		newBlock.Index++
	}

	bc.AddBlock(&newBlock)
}
