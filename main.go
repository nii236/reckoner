package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	btc "github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/knakk/rdf"
	"github.com/pkg/errors"
)

func main() {
	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
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
	c, err := btc.New(config, nil)
	if err != nil {
		log.Fatal(errors.New("Could not create client:"), err)
	}
	for i := 100000; i < 100800; i++ {
		fmt.Println("Block:", i)
		hash, err := c.GetBlockHash(int64(i))
		if err != nil {
			log.Fatalln(err)
		}
		block, err := c.GetBlockVerbose(hash, true)
		if err != nil {
			log.Fatalln(err)
		}
		b := bytes.Buffer{}
		rdf.NewQuadEncoder(&b, rdf.NQuads)
		blockContainer := &blockStruct{
			block,
		}
		result := blockContainer.processBlock()
		for _, blockTriple := range result {
			fmt.Print(blockTriple.Serialize(rdf.NTriples))
		}
	}

	select {}
}

var subjmap map[string]string

func init() {
	subjmap = make(map[string]string)
	subjmap["block.hash"] = "block.hash.%s"
	subjmap["block.time"] = "block.time.%s"
}
