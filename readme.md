# Reckoner
This is a continuation of my personal work from [here](http://blog.thefrontiergroup.com.au/2015/05/blockchain-analytics-with-cayley-db/). This project will parse the entire blockchain into a graph database, in this case [dgraph](dgraph.io).

## Installation

### Dependencies

- btcd
- dgraph
- Go 1.6 or above

## Setup

### btcd

```
git clone https://github.com/btcsuite/btcd $GOPATH/src/github.com/btcsuite/btcd
cd $GOPATH/src/github.com/btcsuite/btcd
git pull && glide install
go install . ./cmd/...
```

### reckoner

```
glide install
go build
./reckoner parse > triples
wc -l triples
tail triples
gzip triples
cayley init -db bolt
cayley load --quads ./triples.gz -db bolt -format nquad -ignoremissing -ignoredup
```

## Usage

```
# Pull blocks 0 to 1000 from the BTCD daemon
$ btcer parse -s 0 -e 1000 | gzip > blocktriples
# Assign UIDs to subjects
$ dgraphassigner --rdfgzips blocktriples.gzip --uids ~/dgraph/uids/
# Load into dgraph
$ dgraphloader --rdfgzips blocktriples.gzip --uids ~/dgraph/uasync.final/ --postings ~/dgraph/p0/
# Run dgraph
$ dgraph --mutations ~/dgraph/m --postings ~/dgraph/p --uids ~/dgraph/u
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
