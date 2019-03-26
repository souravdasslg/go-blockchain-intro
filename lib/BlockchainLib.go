package blockchainlib

import (
	"blockchain-intro/Struct"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CalculateHash(block Struct.Block) string {
	record := strconv.Itoa(block.Index) + block.TimeStamp + strconv.Itoa(block.BPM) + block.PrevHash + block.Nonce
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
	newBlock.Difficulty = Struct.Difficulty
	// newBlock.Hash = CalculateHash(newBlock)
	// return newBlock, nil
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHasValid(CalculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(CalculateHash(newBlock), "need more work to done")
			//time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(CalculateHash(newBlock), "work done")
			newBlock.Hash = CalculateHash(newBlock)
			break
		}
	}
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

func isHasValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("00000", difficulty)
	return strings.HasPrefix(hash, prefix)
}
