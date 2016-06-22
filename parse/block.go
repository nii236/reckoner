package parse

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

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
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create subj IRI"+subj))
	}
	p, err := rdf.NewIRI(pred)
	if err != nil {
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create pred IRI"+pred))
	}
	o, err := rdf.NewLiteral(obj)
	if err != nil {
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create obj literal"+obj))
	}

	if subj == "" {
		log.Fatalln("Empty subject")
	}
	if pred == "" {
		log.Fatalln("Empty predicate")
	}
	if obj == "" {
		oBlank, err := rdf.NewBlank(subj)
		if err != nil {
			log.Fatalln("Error creating blank subject")
		}
		return rdf.Triple{
			Subj: s,
			Pred: p,
			Obj:  oBlank,
		}
	}

	return rdf.Triple{
		Subj: s,
		Pred: p,
		Obj:  o,
	}
}

func iriTripler(subj, pred, obj string) rdf.Triple {
	s, err := rdf.NewIRI(subj)
	if err != nil {
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create subj IRI"+subj))
	}
	p, err := rdf.NewIRI(pred)
	if err != nil {
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create pred IRI"+pred))
	}
	o, err := rdf.NewIRI(obj)
	if err != nil {
		log.Println("Subj:", subj)
		log.Println("Pred:", pred)
		log.Println("Obj:", obj)
		log.Fatalln(errors.Wrap(err, "Can not create obj IRI"+obj))
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
	blockSubj := b.Hash
	result := []rdf.Triple{
		iriTripler(fmt.Sprintf("block.%d", b.Height), "bitcoin.height.block", blockSubj),
		iriTripler(blockSubj, "bitcoin.block.nextblock", b.NextHash),
		iriTripler(blockSubj, "bitcoin.block.prevblock", b.PreviousHash),
		literalTripler(blockSubj, "type", "type.bitcoin.block"),
		literalTripler(blockSubj, "type.bitcoin.block.hash", b.Hash),
		// literalTripler(blockSubj, "type.bitcoin.block.size", strconv.Itoa(int(b.Size))),
		// literalTripler(blockSubj, "type.bitcoin.block.version", strconv.Itoa(int(b.Version))),
		// literalTripler(blockSubj, "type.bitcoin.block.merkleroot", b.MerkleRoot),
		// literalTripler(blockSubj, "type.bitcoin.block.height", strconv.FormatInt(b.Height, 10)),
		// literalTripler(blockSubj, "type.bitcoin.block.time", strconv.FormatInt(b.Time, 10)),
		// literalTripler(blockSubj, "type.bitcoin.block.nonce", strconv.FormatUint(uint64(b.Nonce), 10)),
		literalTripler(blockSubj, "type.bitcoin.block.nexthash", b.NextHash),
		// literalTripler(blockSubj, "type.bitcoin.block.bits", b.Bits),
		// literalTripler(blockSubj, "type.bitcoin.block.difficulty", strconv.FormatFloat(b.Difficulty, 'f', -1, 64)),
		literalTripler(blockSubj, "type.bitcoin.block.previoushash", b.PreviousHash),
	}
	for i, tx := range b.RawTx {
		result = append(result, iriTripler(blockSubj, "bitcoin.block.tx", tx.Txid))
		result = append(result, b.processTX(i)...)
	}
	return result
}

func (b *blockStruct) processTX(txIndex int) []rdf.Triple {
	tx := b.RawTx[txIndex]
	txSubj := tx.Txid
	result := []rdf.Triple{
		iriTripler(txSubj, "bitcoin.tx.block", b.Hash),
		literalTripler(txSubj, "type", "type.bitcoin.tx"),
		// literalTripler(txSubj, "type.bitcoin.tx.hex", tx.Hex),
		// literalTripler(txSubj, "type.bitcoin.tx.txid", tx.Txid),
		// literalTripler(txSubj, "type.bitcoin.tx.version", strconv.Itoa(int(tx.Version))),
		// literalTripler(txSubj, "type.bitcoin.tx.locktime", strconv.FormatUint(uint64(tx.LockTime), 10)),
		// literalTripler(txSubj, "type.bitcoin.tx.blockhash", tx.BlockHash),
		// literalTripler(txSubj, "type.bitcoin.tx.time", strconv.FormatUint(uint64(tx.Time), 10)),
		// literalTripler(txSubj, "type.bitcoin.tx.blocktime", strconv.FormatUint(uint64(tx.Blocktime), 10)),
	}

	for vinIndex := range tx.Vin {
		result = append(result, iriTripler(txSubj, "bitcoin.tx.vin", fmt.Sprintf("%s.vin.%d", tx.Txid, vinIndex)))
		result = append(result, b.processVin(txIndex, vinIndex)...)
	}
	for voutIndex := range tx.Vout {
		result = append(result, iriTripler(txSubj, "bitcoin.tx.vout", fmt.Sprintf("%s.vout.%d", tx.Txid, voutIndex)))
		result = append(result, b.processVout(txIndex, voutIndex)...)
	}
	return result
}

func (b *blockStruct) processVin(txIndex, vinIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	// vin := tx.Vin[vinIndex]
	vinSubj := fmt.Sprintf("%s.vin.%d", tx.Txid, vinIndex)
	result = append(result, []rdf.Triple{
		iriTripler(vinSubj, "bitcoin.vin.tx", tx.Txid),
		literalTripler(vinSubj, "type", "type.bitcoin.vin"),
		// literalTripler(vinSubj, "type.bitcoin.vin.coinbase", vin.Coinbase),
		// literalTripler(vinSubj, "type.bitcoin.vin.sequence", strconv.FormatUint(uint64(vin.Sequence), 10)),
		// literalTripler(vinSubj, "type.bitcoin.vin.txid", vin.Txid),
		// literalTripler(vinSubj, "type.bitcoin.vin.vout", strconv.FormatUint(uint64(vin.Vout), 10)),
		// iriTripler(vinSubj, "bitcoin.vin.vout", fmt.Sprintf("%s.vout.%d", vin.Txid, vin.Vout)),
		// iriTripler(vinSubj, "bitcoin.vin.prevtx", vin.Txid),
	}...)
	return result
}

func (b *blockStruct) processVout(txIndex, voutIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	voutSubj := fmt.Sprintf("%s.vout.%d", tx.Txid, voutIndex)
	// vout := tx.Vout[voutIndex]
	result = append(result, []rdf.Triple{
		iriTripler(voutSubj, "bitcoin.vout.tx", tx.Txid),
		literalTripler(voutSubj, "type", "type.bitcoin.vout"),
		// literalTripler(voutSubj, "type.bitcoin.vout.n", strconv.FormatUint(uint64(vout.N), 10)),
		// literalTripler(voutSubj, "type.bitcoin.vout.value", strconv.FormatFloat(vout.Value, 'f', -1, 64)),
	}...)
	result = append(result, iriTripler(voutSubj, "bitcoin.vout.scriptpubkey", fmt.Sprintf("%s.vout.%d.scriptpubkey", tx.Txid, voutIndex)))
	result = append(result, b.processScriptPubkey(txIndex, voutIndex)...)
	return result
}

func (b *blockStruct) processScriptPubkey(txIndex, voutIndex int) []rdf.Triple {
	result := []rdf.Triple{}
	tx := b.RawTx[txIndex]
	spk := b.RawTx[txIndex].Vout[voutIndex].ScriptPubKey
	scriptPubkeySubj := fmt.Sprintf("%s.vout.%d.scriptpubkey", tx.Txid, voutIndex)
	result = append(result, []rdf.Triple{
		iriTripler(scriptPubkeySubj, "bitcoin.scriptpubkey.vout", fmt.Sprintf("%s.vout.%d", tx.Txid, voutIndex)),
		literalTripler(scriptPubkeySubj, "type", "type.bitcoin.scriptpubkey"),
		// literalTripler(scriptPubkeySubj, "type.bitcoin.scriptpubkey.asm", spk.Asm),
		// literalTripler(scriptPubkeySubj, "type.bitcoin.scriptpubkey.hex", spk.Hex),
		// literalTripler(scriptPubkeySubj, "type.bitcoin.scriptpubkey.reqsigs", strconv.Itoa(int(spk.ReqSigs))),
		// literalTripler(scriptPubkeySubj, "type.bitcoin.scriptpubkey.type", spk.Type),
	}...)
	for _, addr := range spk.Addresses {
		result = append(result, []rdf.Triple{
			iriTripler(scriptPubkeySubj, "bitcoin.scriptpubkey.address", addr),
			iriTripler(addr, "bitcoin.address.scriptpubkey", scriptPubkeySubj),
			literalTripler(addr, "type", "type.bitcoin.address"),
			literalTripler(addr, "type.bitcoin.scriptpubkey.address", addr),
		}...)
	}
	return result
}
