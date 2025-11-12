### Overview
This is a benchmark for sets that are stored on disk. `InMemoryMap` is used as a reference

Tests the following external libraries:
* github.com/outcaste-io/badger/v3
* github.com/tidwall/buntd
* github.com/mattn/go-sqlite3

### Results
```
go test --bench=. -benchmem ./on-disk-set
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/on-disk-set
cpu: Apple M4 Max
BenchmarkInMemoryMap/Exists-14         	48155188	        24.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkInMemoryMap/Add-14            	15589772	       151.9 ns/op	      72 B/op	       0 allocs/op
BenchmarkInMemoryMap/Delete-14         	21023139	        55.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkBadgerDB/Exists-14            	  590360	      1738 ns/op	     622 B/op	      13 allocs/op
BenchmarkBadgerDB/Add-14               	  234094	      5185 ns/op	    1428 B/op	      32 allocs/op
BenchmarkBadgerDB/Delete-14            	  216046	      5792 ns/op	    1428 B/op	      32 allocs/op
BenchmarkBuntDB/Exists-14              	  983827	      1132 ns/op	      96 B/op	       3 allocs/op
BenchmarkBuntDB/Add-14                 	  396726	      2804 ns/op	     720 B/op	       9 allocs/op
BenchmarkBuntDB/Delete-14              	  873738	      1248 ns/op	     304 B/op	       7 allocs/op
BenchmarkSQLite3/Exists-14             	  421768	      2900 ns/op	     536 B/op	      16 allocs/op
BenchmarkSQLite3/Add-14                	   16352	     72994 ns/op	     286 B/op	      10 allocs/op
BenchmarkSQLite3/Delete-14             	  384096	      3139 ns/op	     276 B/op	      10 allocs/op
```
