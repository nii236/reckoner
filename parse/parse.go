package parse

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

// Fetch will fetch blocks
func Fetch(start, end int) {
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
	defer c.Disconnect()
	for i := start; i < end; i++ {
		hash, err := c.GetBlockHash(int64(i))
		if err != nil {
			log.Fatalln(err)
		}
		block, err := c.GetBlockVerbose(hash, true)
		if err != nil {
			log.Fatalln(err)
		}
		b := bytes.Buffer{}
		rdf.NewTripleEncoder(&b, rdf.NTriples)
		blockContainer := &blockStruct{
			block,
		}
		result := blockContainer.processBlock()
		for _, blockTriple := range result {
			fmt.Print(blockTriple.Serialize(rdf.NTriples))
		}
	}
}
