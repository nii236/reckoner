package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	btc "github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
)

func main() {
	notificationHandlers := btc.NotificationHandlers{
		OnTxAccepted: func(hash *wire.ShaHash, amount btcutil.Amount) {
			log.Println("Accepted TX:", hash, amount)
		},
		OnRecvTx: func(transaction *btcutil.Tx, details *btcjson.BlockDetails) {
			log.Println("Received TX:", transaction, details)
		},
		OnBlockConnected: func(hash *wire.ShaHash, height int32, time time.Time) {
			log.Printf("Block connected: %v (%d) %v", hash, height, time)
		},
		OnBlockDisconnected: func(hash *wire.ShaHash, height int32, time time.Time) {
			log.Printf("Block disconnected: %v (%d) %v", hash, height, time)
		},
	}
	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	fmt.Println("BTCDHomeDir:", btcdHomeDir)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(errors.New("Could not get cert"), err)
	}
	config := &btc.ConnConfig{
		Host:         "127.0.0.1:8334",
		Endpoint:     "ws",
		Certificates: certs,
		User:         "user",
		Pass:         "pass",
	}
	c, err := btc.New(config, &notificationHandlers)
	if err != nil {
		log.Fatal(errors.New("Could not create client:"), err)
	}
	c.NotifyBlocks()
	log.Println("NotifyBlocks Done")
	blockCount, err := c.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	select {}
}
