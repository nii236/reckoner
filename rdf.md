This is an example of the RDF data that we expect will be parsed from this app.

```xml
<currblock123> <type.bitcoin.block.hash> "currblock123"
<currblock123> <type.bitcoin.block.tx> "currtx123"
<currblock123> <type.bitcoin.block.tx> "currtx456"
<currblock123> <bitcoin.block.tx> <currtx123>
<currblock123> <bitcoin.block.tx> <currtx456>

<currtx123> <bitcoin.tx.block> <currblock123>
<currtx123> <type.bitcoin.tx.blockhash> "currblock123"
<currtx123> <type.bitcoin.tx.txid> "currtx123"
<currtx123> <bitcoin.tx.vin> <tx.0.currtx123.vin.0>
<currtx123> <bitcoin.tx.vin> <tx.0.currtx123.vin.1>

<currtx123.vin.0> <type.bitcoin.vin.coinbase> "true"
<currtx123.vin.0> <type.bitcoin.vin.sequence> "1234567890"

<currtx123.vin.1> <vin.txid> "prevtx123"
<currtx123.vin.1> <vin.n> "0"
<currtx123.vin.1> <vin.vout> <prevtx123.vout.0>

<currtx123.vin.2> <vin.txid> "prevprevtx123"

<currtx123.vout.0> <vout.n> "0"
<currtx123.vout.0> <vout.value> "50"
<currtx123.vout.0> <vout.scriptpubkey> <currtx123.vout.0.scriptpubkey>

<currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1abc"
<currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1def"
<currtx123.vout.0.scriptpubkey> <scriptpubkey.addresses> "1ghi"
```
