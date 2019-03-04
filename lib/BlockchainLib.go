package blockchainlib

import (
	"blockchain-intro/Struct"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func CalculateHash(block Struct.Block) string {
	record := string(block.Index) + string(block.TimeStamp) + string(block.BPM) + string(block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Struct.Block, BPM int) (Struct.Block, error) {
	var newBlock Struct.Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock Struct.Block) bool {
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []Struct.Block) {
	if len(newBlocks) > len(Struct.Blockchain) {
		Struct.Blockchain = newBlocks
	}
}
