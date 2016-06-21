# Reckoner
This is a continuation of my personal work from [here](http://blog.thefrontiergroup.com.au/2015/05/blockchain-analytics-with-cayley-db/). This project will parse the entire blockchain into a graph database, in this case [dgraph](dgraph.io).

## Installation

### Dependencies

- btcd
- dgraph
- Go 1.6 or above

## Usage

```
# Pull blocks 0 to 1000 from the BTCD daemon
$ btcer parse -s 0 -e 1000 > blocktriples.rdf
# Assign UIDs to subjects
$ dgraphassigner --numInstances 1 --instanceIdx 0 --rdf blocktriples.rdf --uids ~/dgraph/uids/
# Load into dgraph
$ dgraphloader --numInstances 1 --instanceIdx 0 --rdf blocktriples.rdf--uids ~/dgraph/uasync.final --postings ~/dgraph/p0
# Run dgraph
$ dgraph --mutations ~/dgraph/m --postings ~/dgraph/p --uids ~/dgraph/u
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
