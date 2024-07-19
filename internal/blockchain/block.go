package blockchain

import (
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Timestamp string
	Transaction transaction.Transaction
	PrevHash string
	Hash string
	Nonce int
}

type Blockchain struct {
	Chain []Block
}

func (bc *Blockchain) LastBlock() *Block{
	return &bc.Chain[len(bc.Chain) - 1]
}

func (block *Block) computeHash() string {
	record := strconv.Itoa(block.Nonce) + block.Timestamp + block.Transaction.ToString() + block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	hash := hex.EncodeToString(hashed)
	block.Hash = hash

	return hash
}

func generateBlock(prevBlock Block, t transaction.Transaction) *Block {
	var newBlock = &Block{}

	newBlock.Nonce = 0
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = t
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = newBlock.computeHash()
	return newBlock
}

func (bc *Blockchain) AddBlock(block *Block) {
	bc.Chain = append(bc.Chain, *block)
}

func (bc *Blockchain) AddTransaction(t transaction.Transaction, publicKey *rsa.PublicKey, signature string) {
	if encryption.VerifySignature(publicKey, []byte(t.ToString()), []byte(signature)) == nil {
		block := generateBlock(*bc.LastBlock(), t)
		bc.Chain = append(bc.Chain, *block)
	}
}

func NewGenesisBlock() Block {
	timestamp := time.Now().String()
	prevHash := "0"

	payerPublicKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	payeePublicKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	genesisTransaction := transaction.Transaction{
		Amount: 100,
		Payer:  &payerPublicKey.PublicKey,
		Payee:  &payeePublicKey.PublicKey,
	}

	genesisBlock := Block{
		Nonce:       0,
		Timestamp:   timestamp,
		Transaction: genesisTransaction,
		PrevHash:    prevHash,
	}

	genesisBlock.computeHash()

	return genesisBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain: []Block{
			NewGenesisBlock(),
		},
	}
}
