### Overview
This is a benchmark for serializing and deserializing structs

Tests the following external libraries:
* github.com/fxamacker/cbor/v2
* github.com/pquerna/ffjson/fflib/v1
* google.golang.org/protobuf/proto

### Results
```
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/serde
cpu: Apple M4 Max
BenchmarkJsonMarshal-14        	 8861052	       117.5 ns/op	     112 B/op	       2 allocs/op
BenchmarkJsonUnmarshal-14      	 6554515	       169.5 ns/op	     232 B/op	       5 allocs/op
BenchmarkFfjsonMarshal-14      	 9966855	       118.9 ns/op	     112 B/op	       2 allocs/op
BenchmarkFfjsonUnmarshal-14    	 7060164	       168.4 ns/op	     232 B/op	       5 allocs/op
BenchmarkCborMarshal-14        	11909743	        97.50 ns/op	      96 B/op	       2 allocs/op
BenchmarkCborUnmarshal-14      	29174127	        40.85 ns/op	     112 B/op	       3 allocs/op
BenchmarkProtobufMarshal-14    	12745878	        97.56 ns/op	     168 B/op	       3 allocs/op
BenchmarkProtobufUnmarshal-14  	11065078	       106.7 ns/op	     152 B/op	       4 allocs/op
PASS
```
