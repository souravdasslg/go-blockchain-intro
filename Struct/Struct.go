package Struct

import "sync"

// Block
type Block struct {
	Index      int    `bson:"index" json:"index"`
	TimeStamp  string `bson:"timeStamp" json:"timeStamp"`
	BPM        int    `bson:"BPM" json:"BPM"`
	Hash       string `bson:"hash" json:"hash"`
	PrevHash   string `bson:"prevHash" json:"prevHash"`
	Difficulty int    `bson:"difficulty" json:"difficulty"`
	Nonce      string `bson:"nonce" json:"nonce"`
}

type Message struct {
	BPM int
}

var Blockchain []Block

var Mutex = &sync.Mutex{}

var bcServer = make(chan []Block)

const Difficulty = 5
