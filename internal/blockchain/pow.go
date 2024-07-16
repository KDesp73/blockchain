package blockchain

import (
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const difficulty = 5

func (bc *Blockchain) MineBlock(t transaction.Transaction, publicKey *rsa.PublicKey, signature string) {
	if encryption.VerifySignature(publicKey, []byte(t.ToString()), []byte(signature)) != nil {
		fmt.Println("Could not verify the signature: ", signature)
		return
	}


	fmt.Println("⛏️ Mining...")
	newBlock := generateBlock(*bc.LastBlock(), t)

	solution := 1
	for {
		h := sha256.New()
		h.Write([]byte(strconv.Itoa(newBlock.Nonce + solution)))
		hashed := h.Sum(nil)
		hash := hex.EncodeToString(hashed)
		
		if hash[:difficulty] == strings.Repeat("0", difficulty) {
			fmt.Printf("Hash found: %s\n\n", hash)
			break
		}
		newBlock.Nonce++
	}

	bc.AddBlock(newBlock)
}
