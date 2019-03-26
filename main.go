package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"blockchain-intro/Struct"
	Controllers "blockchain-intro/controllers"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
}
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

func initiateDB(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/blockchain"))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to mongo db %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could Not Connect With DB in Background")
	}
	blockchainDB := client.Database("blockchain")
	return blockchainDB, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Struct.Block{1, t.String(), 0, "", " ", 0, ""}
		spew.Dump(genesisBlock)
		fmt.Println("Length Of Blockchain Prev", len(Struct.Blockchain))
		Struct.Mutex.Lock()
		Struct.Blockchain = append(Struct.Blockchain, genesisBlock)
		Struct.Mutex.Unlock()
		tcpserver, err := net.Listen("tcp", ":"+os.Getenv("TCPADDR"))
		if err != nil {
			log.Fatal(err)
		}
		defer tcpserver.Close()
		for {
			conn, err := tcpserver.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleConn(conn)
		}
	}()
	log.Fatal(run())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	db, err := initiateDB(ctx)
	if err != nil {
		log.Fatalf("Database configuration Crashed %v", err)
	}

}
