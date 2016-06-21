package main

import (
	"fmt"
	"strconv"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/knakk/rdf"
)

type blockStruct struct {
	*btcjson.GetBlockVerboseResult
}

func literalTripler(subj, pred, obj string) rdf.Triple {
	s, _ := rdf.NewIRI(subj)
	p, _ := rdf.NewIRI(pred)
	o, _ := rdf.NewLiteral(obj)
	t := rdf.Triple{
		Subj: s,
		Pred: p,
		Obj:  o,
	}
	return t
}

func iriTripler(subj, pred, obj string) rdf.Triple {
	s, _ := rdf.NewIRI(subj)
	p, _ := rdf.NewIRI(pred)
	o, _ := rdf.NewIRI(obj)
	t := rdf.Triple{
		Subj: s,
		Pred: p,
		Obj:  o,
	}
	return t
}

func (b *blockStruct) processBlock() []rdf.Triple {
	blockSubj := fmt.Sprintf("block.hash.%s", b.Hash)
	result := []rdf.Triple{
		literalTripler(blockSubj, "block.hash", b.Hash),
		literalTripler(blockSubj, "block.time", strconv.FormatInt(b.Time, 10)),
		literalTripler(blockSubj, "block.nexthash", b.NextHash),
		literalTripler(blockSubj, "block.height", strconv.FormatInt(b.Height, 10)),
	}
	for i := range b.RawTx {
		result = append(result, b.processTX(i)...)
	}
	return result
}

func (b *blockStruct) processTX(txIndex int) []rdf.Triple {
	tx := b.RawTx[txIndex]
	txSubj := fmt.Sprintf("tx.%d.%s", txIndex, tx.Txid)
	result := []rdf.Triple{
		literalTripler(txSubj, "tx.txid", tx.Txid),
	}

	for vinIndex := range tx.Vin {
		result = append(result, b.processVin(txIndex, vinIndex)...)
	}
	for voutIndex := range tx.Vout {
		result = append(result, b.processVout(txIndex, voutIndex)...)
	}
	return result
}

func (b *blockStruct) processVin(txIndex, vinIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	vin := tx.Vin[vinIndex]
	vinSubj := fmt.Sprintf("tx.%d.%s.vin.%d", txIndex, tx.Txid, vinIndex)
	result = append(result, []rdf.Triple{
		literalTripler(vinSubj, "vin.coinbase", vin.Coinbase),
		literalTripler(vinSubj, "vin.sequence", strconv.FormatUint(uint64(vin.Sequence), 10)),
		literalTripler(vinSubj, "vin.txid", vin.Txid),
		literalTripler(vinSubj, "vin.vout", strconv.FormatUint(uint64(vin.Vout), 10)),
	}...)
	return result
}

func (b *blockStruct) processVout(txIndex, voutIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	voutSubj := fmt.Sprintf("tx.%d.%s.vout.%d", txIndex, tx.Txid, voutIndex)
	vout := tx.Vout[voutIndex]
	result = append(result, []rdf.Triple{
		literalTripler(voutSubj, "vout.value", strconv.FormatFloat(vout.Value, 'f', -1, 64)),
	}...)
	for addrIndex := range vout.ScriptPubKey.Addresses {
		result = append(result, b.processAddress(txIndex, voutIndex, addrIndex))
	}
	return result
}

func (b *blockStruct) processAddress(txIndex, voutIndex, addrIndex int) rdf.Triple {
	addressSubj := fmt.Sprintf("tx.%d.vout.%d.scriptpubkey", txIndex, voutIndex)
	address := b.RawTx[txIndex].Vout[voutIndex].ScriptPubKey.Addresses[addrIndex]
	return literalTripler(addressSubj, "scriptpubkey.addresses", address)
}
