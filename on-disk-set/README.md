### Overview
This is a benchmark for sets that are stored on disk. `InMemoryMap` is used as a reference

Tests the following external libraries:
* github.com/outcaste-io/badger/v3
* github.com/tidwall/buntd
* github.com/naqvijafar91/cuteDB
* github.com/mattn/go-sqlite3

### Results

```
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/on-disk-set
cpu: Apple M4 Max
BenchmarkInMemoryMap/Exists-14         	46468292	        25.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkInMemoryMap/Add-14            	17366178	       154.7 ns/op	      65 B/op	       0 allocs/op
BenchmarkInMemoryMap/Delete-14         	21620500	        56.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkBadgerDB/Exists-14            	  711852	      1463 ns/op	     452 B/op	      11 allocs/op
BenchmarkBadgerDB/Add-14               	  217690	      5457 ns/op	    1428 B/op	      32 allocs/op
BenchmarkBadgerDB/Delete-14            	  211894	      5605 ns/op	    1428 B/op	      32 allocs/op
BenchmarkCuteDB/Exists-14              	  182904	      6087 ns/op	    7542 B/op	     186 allocs/op
BenchmarkCuteDB/Add-14                 	  132280	      8474 ns/op	   15929 B/op	     218 allocs/op
BenchmarkCuteDB/Delete-14              	129270724	         9.229 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuntDB/Exists-14              	 1785706	       669.1 ns/op	      96 B/op	       3 allocs/op
BenchmarkBuntDB/Add-14                 	  453660	      2214 ns/op	     721 B/op	       9 allocs/op
BenchmarkBuntDB/Delete-14              	 1226276	       965.3 ns/op	     304 B/op	       7 allocs/op
```
