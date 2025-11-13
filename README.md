 # Go Benchmarks

### Overview
This repo started out as just a benchmark for various Go concurrent map implementations. Today this new repo aims to test various algorithms and data structures:
* Concurrent maps
* On disk sets
* Serde

### Prerequisites
For RocksDB support, you need to install the RocksDB C++ library:

**macOS (Homebrew):**
```bash
brew install rocksdb
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install librocksdb-dev
```

**Linux (CentOS/RHEL):**
```bash
sudo yum install rocksdb-devel
```

### Running
Run functional tests: `make test`

Run benchmarks: `make bench-[map|set|serde]`
