### Overview
This is a benchmark for serializing and deserializing structs

Tests the following external libraries:
* github.com/fxamacker/cbor/v2
* github.com/pquerna/ffjson/fflib/v1

### Results
```
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/serde
cpu: Apple M4 Max
BenchmarkJsonMarshal-14        	10189734	       116.1 ns/op	     112 B/op	       2 allocs/op
BenchmarkJsonUnmarshal-14      	 7199274	       165.2 ns/op	     232 B/op	       5 allocs/op
BenchmarkFfjsonMarshal-14      	10154517	       122.2 ns/op	     112 B/op	       2 allocs/op
BenchmarkFfjsonUnmarshal-14    	 6596562	       174.3 ns/op	     232 B/op	       5 allocs/op
BenchmarkCborMarshal-14        	12534366	        96.36 ns/op	      96 B/op	       2 allocs/op
BenchmarkCborUnmarshal-14      	29101696	        42.31 ns/op	     112 B/op	       3 allocs/op
PASS
```
