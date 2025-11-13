# RocksDB CGO flags (adjust paths for your OS)
# macOS (Homebrew): /opt/homebrew
# Linux: /usr
ROCKSDB_PREFIX ?= /opt/homebrew
CGO_CFLAGS_ROCKSDB = -I$(ROCKSDB_PREFIX)/include
CGO_LDFLAGS_ROCKSDB = -L$(ROCKSDB_PREFIX)/lib -lrocksdb

build:
	CGO_CFLAGS="$(CGO_CFLAGS_ROCKSDB)" CGO_LDFLAGS="$(CGO_LDFLAGS_ROCKSDB)" go build ./...

test:
	CGO_CFLAGS="$(CGO_CFLAGS_ROCKSDB)" CGO_LDFLAGS="$(CGO_LDFLAGS_ROCKSDB)" go test ./...

bench-map:
	go test --bench=. ./concurrent-map

bench-set:
	CGO_CFLAGS="$(CGO_CFLAGS_ROCKSDB)" CGO_LDFLAGS="$(CGO_LDFLAGS_ROCKSDB)" go test --bench=. -benchmem ./on-disk-set

bench-serde:
	go test --bench=. -benchmem ./serde

deps:
	@echo "Installing protoc-gen-go..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "Checking for protoc installation..."
	@which protoc > /dev/null || (echo "Error: protoc not found. Please install it via: brew install protobuf" && exit 1)
	@echo "Checking for RocksDB installation..."
	@test -f $(ROCKSDB_PREFIX)/include/rocksdb/c.h || (echo "Error: RocksDB not found. Please install it via: brew install rocksdb (macOS) or apt-get install librocksdb-dev (Linux)" && exit 1)
	@echo "Generating protobuf Go files..."
	@cd serde/models && protoc --go_out=. --go_opt=paths=source_relative object.proto
	@echo "Done! Protobuf files generated."

generate:
	ffjson serde/models/object.go
	cd serde/models && protoc --go_out=. --go_opt=paths=source_relative object.proto
