package Struct

type Block struct {
	Index     int
	TimeStamp string
	BPM       int
	Hash      string
	PrevHash  string
}

type Message struct {
	BPM int
}

var Blockchain []Block
