build:
	go build ./...

test:
	go test ./...

bench-map:
	go test --bench=. ./concurrent-map

bench-set:
	go test --bench=. -benchmem ./on-disk-set

bench-serde:
	go test --bench=. -benchmem ./serde

generate:
	ffjson serde/models/object.go
