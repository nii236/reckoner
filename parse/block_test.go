package parse

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/knakk/rdf"
)

func TestTripler(t *testing.T) {
	b := &btcjson.GetBlockVerboseResult{
		Hash: "abc",
	}
	triples := []rdf.Triple{}
	triples = append(triples, literalTripler("block.hash.block123", "block.hash", b.Hash))
	triples = append(triples, iriTripler("block.hash.block123", "block.nextblock", "block123"))

	var bs bytes.Buffer
	enc := rdf.NewTripleEncoder(&bs, rdf.NTriples)
	enc.EncodeAll(triples)
	enc.Close()
	t.Log("\n" + bs.String())

}
