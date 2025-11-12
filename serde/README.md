### Overview
This is a benchmark for serializing and deserializing structs

Tests the following implementations:
* Custom binary format
* encoding/gob
* encoding/json

Tests the following external libraries:
* github.com/fxamacker/cbor/v2
* github.com/pquerna/ffjson/fflib/v1
* google.golang.org/protobuf/proto

### Results
```
go test --bench=. -benchmem ./serde
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/serde
cpu: Apple M4 Max
BenchmarkJsonMarshal-14              	 8938593	       119.1 ns/op	     112 B/op	       2 allocs/op
BenchmarkJsonUnmarshal-14            	 6985383	       175.1 ns/op	     232 B/op	       5 allocs/op
BenchmarkFfjsonMarshal-14            	 9406219	       121.4 ns/op	     112 B/op	       2 allocs/op
BenchmarkFfjsonUnmarshal-14          	 7055457	       175.1 ns/op	     232 B/op	       5 allocs/op
BenchmarkCborMarshal-14              	12039806	        95.98 ns/op	      96 B/op	       2 allocs/op
BenchmarkCborUnmarshal-14            	28485519	        41.01 ns/op	     112 B/op	       3 allocs/op
BenchmarkProtobufMarshal-14          	12478984	        96.75 ns/op	     168 B/op	       3 allocs/op
BenchmarkProtobufUnmarshal-14        	11295020	       105.7 ns/op	     152 B/op	       4 allocs/op
BenchmarkCustomBinaryMarshal-14      	13385648	        90.85 ns/op	     160 B/op	       8 allocs/op
BenchmarkCustomBinaryUnmarshal-14    	12282355	        93.67 ns/op	      96 B/op	       9 allocs/op
BenchmarkGobMarshal-14               	 1000000	      1038 ns/op	    1392 B/op	      19 allocs/op
BenchmarkGobUnmarshal-14             	  205678	      5775 ns/op	    7912 B/op	     182 allocs/op
```
