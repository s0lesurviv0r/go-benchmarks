### Overview
This repo has both correctness tests and benchmarks for a variety of different thread/goroutine safe maps.

Tests the following types of concurrent maps:
* Unsharded map with `sync.Mutex`
* Unsharded map with `sync.RWMutex`
* Sharded map with each shard having it's own `sync.Mutex`
* Sharded map with each shard having it's own `sync.RWMutex`
* Sharded map with that has a off-to-the side array of `sync.RWMutex` for the shards

Tests the following external libraries:
* github.com/cornelk/hashmap (IN PROGRESS)
* github.com/dustinxie/lockfree
* github.com/fanliao/go-concurrentMap
* github.com/orcaman/concurrent-map
* github.com/tidwall/shardmap

Wherever possible, the number of shards for a map is set to 32.

### Results
```
go test --bench=. ./concurrent-map
goos: darwin
goarch: arm64
pkg: github.com/s0lesurviv0r/go-benchmarks/concurrent-map
cpu: Apple M4 Max
BenchmarkUnshardedSingleMutex/Get-14            	10007247	       119.7 ns/op
BenchmarkUnshardedSingleMutex/Set-14            	 4001626	       298.9 ns/op
BenchmarkUnshardedSingleMutex/Mix-14            	 5376544	       223.0 ns/op
BenchmarkUnshardedSingleRWMutex/Get-14          	 8871884	       135.9 ns/op
BenchmarkUnshardedSingleRWMutex/Set-14          	 3796230	       313.7 ns/op
BenchmarkUnshardedSingleRWMutex/Mix-14          	 8554413	       140.6 ns/op
BenchmarkShardedMultiMutexMap/Get-14            	15784929	        75.67 ns/op
BenchmarkShardedMultiMutexMap/Set-14            	10993722	       110.1 ns/op
BenchmarkShardedMultiMutexMap/Mix-14            	14174853	        85.39 ns/op
BenchmarkShardedMultiRWMutexMap/Get-14          	15683540	        76.61 ns/op
BenchmarkShardedMultiRWMutexMap/Set-14          	 8731974	       114.7 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-14          	15294814	        78.58 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-14         	15375448	        78.19 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-14         	 8343808	       123.4 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-14         	15486842	        77.53 ns/op
BenchmarkOrcamanLibrary/Get-14                           	15630129	        76.86 ns/op
BenchmarkOrcamanLibrary/Set-14                           	 8741568	       116.6 ns/op
BenchmarkOrcamanLibrary/Mix-14                           	15113881	        79.47 ns/op
BenchmarkFanLiaoLibrary/Get-14                           	16819285	        71.52 ns/op
BenchmarkFanLiaoLibrary/Set-14                           	 8913464	       134.5 ns/op
BenchmarkFanLiaoLibrary/Mix-14                           	15996764	        76.53 ns/op
BenchmarkTidwallLibrary/Get-14                           	15792711	        75.98 ns/op
BenchmarkTidwallLibrary/Set-14                           	11127793	        94.96 ns/op
BenchmarkTidwallLibrary/Mix-14                           	15331918	        78.25 ns/op
BenchmarkDustinxieLibrary/Get-14                         	 7230453	       165.8 ns/op
BenchmarkDustinxieLibrary/Set-14                         	 6600925	       181.3 ns/op
BenchmarkDustinxieLibrary/Mix-14                         	 7637224	       156.5 ns/op
BenchmarkSyncMap/Get-14                                  	20314407	        58.47 ns/op
BenchmarkSyncMap/Set-14                                  	11409482	        99.23 ns/op
BenchmarkSyncMap/Mix-14                                  	16014715	        74.79 ns/op
```
