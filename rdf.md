This is an example of the RDF data that we expect will be parsed from this app.

```xml
<block.hash.currblock123> <block.hash> "currblock123"
<block.hash.currblock123> <block.rawtx> <tx.0.txid.currtx123>
<block.hash.currblock123> <block.rawtx> <tx.1.txid.currtx456>

<tx.0.txid.currtx123> <tx.block> <block.hash.currblock123>
<tx.0.txid.currtx123> <tx.blockhash> "currblock123"
<tx.0.txid.currtx123> <tx.txid> "currtx123"
<tx.0.txid.currtx123> <tx.vin> <tx.0.txid.currtx123.vin.0>
<tx.0.txid.currtx123> <tx.vin> <tx.0.txid.currtx123.vin.1>

<tx.0.txid.currtx123.vin.0> <vin.coinbase> "true"
<tx.0.txid.currtx123.vin.0> <vin.sequence> "1234567890"

<tx.0.txid.currtx123.vin.1> <vin.txid> "prevtx123"
<tx.0.txid.currtx123.vin.1> <vin.n> "0"
<tx.0.txid.currtx123.vin.1> <vin.vout> <tx.txid.prevtx123.vout.0>

<tx.0.txid.currtx123.vin.2> <vin.txid> "prevprevtx123"

<tx.0.txid.currtx123.vout.0> <vout.n> "0"
<tx.0.txid.currtx123.vout.0> <vout.value> "50"
<tx.0.txid.currtx123.vout.0> <vout.scriptpubkey> <tx.txid.currtx123.vout.0.scriptpubkey>

<tx.0.txid.currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1abc"
<tx.0.txid.currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1def"
<tx.0.txid.currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1ghi"
```
