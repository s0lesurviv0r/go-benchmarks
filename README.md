 # Go Benchmarks

### Overview
This repo started out as just a benchmark for various Go concurrent map implementations. Today this new repo aims to test various algorithms and data structures:
* Concurrent maps (COMPLETE)
* On disk sets (IN PROGRESS)
* Serde (IN PROGRESS)

### Running
`make bench-[map|set|serde]`

### TODO

#### On Disk Set
- [ ] Try skiplist on BadgerDB
- [ ] Try bbolt
- [ ] Try https://github.com/tidwall/buntdb

#### Serde
- [ ] Try Protobuf
- [ ] Try custom binary format
