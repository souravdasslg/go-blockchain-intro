package blockchain

import (
	"blockchain-intro/Struct"
	blockchainlib "blockchain-intro/lib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func HandleGetBlockChain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Struct.Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

/* HandleWriteBlock prints all the blocks */
func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Length Of Blockchain", len(Struct.Blockchain)-1)
	spew.Dump(Struct.Blockchain)
	var m Struct.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	Struct.Mutex.Lock()
	newBlock, err := blockchainlib.GenerateBlock(Struct.Blockchain[len(Struct.Blockchain)-1], m.BPM)
	Struct.Mutex.Unlock()

	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
	}
	if blockchainlib.IsBlockValid(newBlock, Struct.Blockchain[len(Struct.Blockchain)-1]) {
		newBlockChain := append(Struct.Blockchain, newBlock)
		blockchainlib.ReplaceChain(newBlockChain)
		spew.Dump(newBlockChain)
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Http : 500. Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
