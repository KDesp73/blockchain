package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int
	Timestamp string
	Data string
	PrevHash string
	Hash string
}

type Blockchain struct {
	Chain []Block
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func (bc *Blockchain) AddBlock(data string) {
	var prevBlock Block
	if len(bc.Chain) == 0 {
		prevBlock = Block{
			Index: -1, Data: "", Timestamp: "", PrevHash: "",
		}
	} else {
		prevBlock = bc.Chain[len(bc.Chain) - 1]
	}
	newBlock := generateBlock(prevBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}


