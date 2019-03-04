package Controllers

import (
	BlockchainController "blockchain-intro/controllers/blockchain"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", BlockchainController.HandleGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", BlockchainController.HandleWriteBlock).Methods("POST")
	return muxRouter

}
