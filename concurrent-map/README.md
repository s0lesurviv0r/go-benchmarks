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
Benchmarked with Go 1.24.0 on KDE Neon 24.04 with a 13th Gen Intel(R) Core(TM) i5-1340P
```
goos: linux
goarch: amd64
pkg: github.com/s0lesurviv0r/go-benchmarks/concurrent-map
cpu: 13th Gen Intel(R) Core(TM) i5-1340P
BenchmarkUnshardedSingleMutex/Get-16                     6318457               199.3 ns/op
BenchmarkUnshardedSingleMutex/Set-16                     1000000              1140 ns/op
BenchmarkUnshardedSingleMutex/Mix-16                     2662177               578.5 ns/op
BenchmarkUnshardedSingleRWMutex/Get-16                  16861977                63.09 ns/op
BenchmarkUnshardedSingleRWMutex/Set-16                   1000000              1169 ns/op
BenchmarkUnshardedSingleRWMutex/Mix-16                  15507824                76.85 ns/op
BenchmarkShardedMultiMutexMap/Get-16                    25105923                43.67 ns/op
BenchmarkShardedMultiMutexMap/Set-16                    11875528               103.7 ns/op
BenchmarkShardedMultiMutexMap/Mix-16                    20151253                58.74 ns/op
BenchmarkShardedMultiRWMutexMap/Get-16                  28990989                37.37 ns/op
BenchmarkShardedMultiRWMutexMap/Set-16                  11193744               108.8 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-16                  25577576                46.91 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-16                29119434                41.13 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-16                12092128               100.2 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-16                25431547                44.99 ns/op
BenchmarkOrcamanLibrary/Get-16                                  27079370                40.93 ns/op
BenchmarkOrcamanLibrary/Set-16                                  10308474               119.9 ns/op
BenchmarkOrcamanLibrary/Mix-16                                  22908312                49.76 ns/op
BenchmarkFanLiaoLibrary/Get-16                                  30281772                38.80 ns/op
BenchmarkFanLiaoLibrary/Set-16                                   7234028               156.9 ns/op
BenchmarkFanLiaoLibrary/Mix-16                                  21038586                48.08 ns/op
BenchmarkTidwallLibrary/Get-16                                  35683909                33.44 ns/op
BenchmarkTidwallLibrary/Set-16                                  13203945                87.66 ns/op
BenchmarkTidwallLibrary/Mix-16                                  28398628                42.42 ns/op
BenchmarkDustinxieLibrary/Get-16                                15378068                77.98 ns/op
BenchmarkDustinxieLibrary/Set-16                                 5670385               193.9 ns/op
BenchmarkDustinxieLibrary/Mix-16                                 7471184               158.7 ns/op
BenchmarkSyncMap/Get-16                                         35217488                29.23 ns/op
BenchmarkSyncMap/Set-16                                         10701004               119.3 ns/op
BenchmarkSyncMap/Mix-16                                         21130794                58.10 ns/op
```
