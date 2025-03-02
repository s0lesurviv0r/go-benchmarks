### Overview
**In Progress**

This is a benchmark for on disk sets. `InMemorySet` is used as a reference

### Results
Benchmarked with Go 1.24.0 on KDE Neon 24.04 with a 13th Gen Intel(R) Core(TM) i5-1340P
```
goos: linux
goarch: amd64
pkg: github.com/s0lesurviv0r/go-benchmarks/on-disk-set
cpu: 13th Gen Intel(R) Core(TM) i5-1340P
BenchmarkInMemoryMap/Exists-16          27587095                36.57 ns/op            0 B/op          0 allocs/op
BenchmarkInMemoryMap/Add-16              6870076               192.2 ns/op            76 B/op          0 allocs/op
BenchmarkInMemoryMap/Delete-16          12533425               119.5 ns/op             0 B/op          0 allocs/op
BenchmarkBadgerDB/Exists-16               703706              1683 ns/op             428 B/op         10 allocs/op
BenchmarkBadgerDB/Add-16                   83558             14761 ns/op            1432 B/op         32 allocs/op
BenchmarkBadgerDB/Delete-16                57964             17629 ns/op            1433 B/op         32 allocs/op
BenchmarkCuteDB/Exists-16                 158510              7893 ns/op            8048 B/op        198 allocs/op
BenchmarkCuteDB/Add-16                     96351             10697 ns/op           16276 B/op        226 allocs/op
BenchmarkCuteDB/Delete-16               80028921                14.52 ns/op            0 B/op          0 allocs/op
```