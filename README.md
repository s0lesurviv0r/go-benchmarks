 # Go Benchmarks

### Overview
This repo started out as just a benchmark for various Go concurrent map implementations. Today this new repo aims to test various algorithms and data structures:
* Concurrent maps
* On disk sets
* Serde

### Running
Run functional tests: `make test`

Run benchmarks: `make bench-[map|set|serde]`
