package blockchain

import "strings"

const difficulty = 1

func (bc *Blockchain) mineBlock(newBlock *Block) {
	for {
		newBlock.Hash = calculateHash(*newBlock)
		
		if newBlock.Hash[:difficulty] == strings.Repeat("0", difficulty) {
			break
		}
		newBlock.Index++
	}
}
