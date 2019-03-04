package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"blockchain-intro/Struct"
	Controllers "blockchain-intro/controllers"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func run() error {
	mux := Controllers.MakeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Server Has Been Running On: " + httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Struct.Block{1, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		fmt.Println("Length Of Blockchain Prev", len(Struct.Blockchain))
		Struct.Blockchain = append(Struct.Blockchain, genesisBlock)
	}()
	log.Fatal(run())
}
