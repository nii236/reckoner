package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/knakk/rdf"
)

type blockStruct struct {
	*btcjson.GetBlockVerboseResult
}

func literalTripler(subj, pred, obj string) rdf.Triple {
	var err error
	s, err := rdf.NewIRI(subj)
	if err != nil {
		log.Fatalln(err)
	}
	p, err := rdf.NewIRI(pred)
	if err != nil {
		log.Fatalln(err)
	}
	o, err := rdf.NewLiteral(obj)
	if err != nil {
		log.Fatalln(err)
	}
	if s.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty subject")
	}
	if p.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty predicate")
	}
	if o.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty object")
	}
	t := rdf.Triple{
		Subj: s,
		Pred: p,
		Obj:  o,
	}
	return t
}

func iriTripler(subj, pred, obj string) rdf.Triple {
	s, err := rdf.NewIRI(subj)
	if err != nil {
		log.Fatalln(err)
	}
	p, err := rdf.NewIRI(pred)
	if err != nil {
		log.Fatalln(err)
	}
	o, err := rdf.NewIRI(obj)
	if err != nil {
		log.Fatalln(err)
	}

	if s.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty subject")
	}
	if p.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty predicate")
	}
	if o.Serialize(rdf.NTriples) == "" {
		log.Println("Subj:", s)
		log.Println("Pred:", p)
		log.Println("Obj:", o)
		log.Fatalln("Empty object")
	}
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
		literalTripler(blockSubj, "block.size", strconv.Itoa(int(b.Size))),
		literalTripler(blockSubj, "block.version", strconv.Itoa(int(b.Version))),
		literalTripler(blockSubj, "block.merkleroot", b.MerkleRoot),
		literalTripler(blockSubj, "block.height", strconv.FormatInt(b.Height, 10)),
		literalTripler(blockSubj, "block.time", strconv.FormatInt(b.Time, 10)),
		literalTripler(blockSubj, "block.nonce", strconv.FormatUint(uint64(b.Nonce), 10)),
		literalTripler(blockSubj, "block.nexthash", b.NextHash),
		literalTripler(blockSubj, "block.bits", b.Bits),
		literalTripler(blockSubj, "block.difficulty", strconv.FormatFloat(b.Difficulty, 'f', -1, 64)),
		literalTripler(blockSubj, "block.previoushash", b.PreviousHash),
		literalTripler(blockSubj, "block.nexthash", b.NextHash),
	}
	for i, tx := range b.RawTx {
		result = append(result, iriTripler(blockSubj, "block.rawtx", fmt.Sprintf("tx.%d.%s", i, tx.Txid)))
		result = append(result, b.processTX(i)...)
	}
	return result
}

func (b *blockStruct) processTX(txIndex int) []rdf.Triple {
	tx := b.RawTx[txIndex]
	txSubj := fmt.Sprintf("tx.%d.%s", txIndex, tx.Txid)
	result := []rdf.Triple{
		literalTripler(txSubj, "tx.hex", tx.Hex),
		literalTripler(txSubj, "tx.txid", tx.Txid),
		literalTripler(txSubj, "tx.version", strconv.Itoa(int(tx.Version))),
		literalTripler(txSubj, "tx.locktime", strconv.FormatUint(uint64(tx.LockTime), 10)),
		literalTripler(txSubj, "tx.blockhash", tx.BlockHash),
		literalTripler(txSubj, "tx.time", strconv.FormatUint(uint64(tx.Time), 10)),
		literalTripler(txSubj, "tx.blocktime", strconv.FormatUint(uint64(tx.Blocktime), 10)),
	}

	for vinIndex := range tx.Vin {
		result = append(result, iriTripler(txSubj, "tx.vin", fmt.Sprintf("tx.%d.%s.vin.%d", txIndex, tx.Txid, vinIndex)))
		result = append(result, b.processVin(txIndex, vinIndex)...)
	}
	for voutIndex := range tx.Vout {
		result = append(result, iriTripler(txSubj, "tx.vout", fmt.Sprintf("tx.%d.%s.vout.%d", txIndex, tx.Txid, voutIndex)))
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
		literalTripler(voutSubj, "vout.n", strconv.FormatUint(uint64(vout.N), 10)),
		literalTripler(voutSubj, "vout.value", strconv.FormatFloat(vout.Value, 'f', -1, 64)),
	}...)
	result = append(result, iriTripler(voutSubj, "vout.scriptpubkey", fmt.Sprintf("tx.%d.%s.vout.%d.scriptpubkey", txIndex, tx.Txid, voutIndex)))
	result = append(result, b.processScriptPubkey(txIndex, voutIndex)...)
	return result
}

func (b *blockStruct) processScriptPubkey(txIndex, voutIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	spk := b.RawTx[txIndex].Vout[voutIndex].ScriptPubKey
	scriptPubkeySubj := fmt.Sprintf("tx.%d.%s.vout.%d.scriptpubkey", txIndex, tx.Txid, voutIndex)
	result = append(result, []rdf.Triple{
		literalTripler(scriptPubkeySubj, "scriptpubkey.asm", spk.Asm),
		literalTripler(scriptPubkeySubj, "scriptpubkey.hex", spk.Hex),
		literalTripler(scriptPubkeySubj, "scriptpubkey.reqsigs", strconv.Itoa(int(spk.ReqSigs))),
		literalTripler(scriptPubkeySubj, "scriptpubkey.type", spk.Type),
	}...)
	for _, addr := range spk.Addresses {
		literalTripler(scriptPubkeySubj, "scriptpubkey.addresses", addr)
	}
	return result
}
