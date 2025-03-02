package serde

import (
	"encoding/json"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/pquerna/ffjson/ffjson"

	"github.com/s0lesurviv0r/go-benchmarks/serde/models"
)

func getTestObject() models.Object {
	return models.Object{
		Id:   uint64(123),
		Name: "test",
		Nested: models.Nested{
			Id:   uint64(123),
			Name: "test",
		},
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(getTestObject())
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	buf, _ := json.Marshal(getTestObject())
	var obj models.Object
	for i := 0; i < b.N; i++ {
		_ = ffjson.Unmarshal(buf, obj)
	}
}

func BenchmarkFfjsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ffjson.Marshal(getTestObject())
	}
}

func BenchmarkFfjsonUnmarshal(b *testing.B) {
	buf, _ := ffjson.Marshal(getTestObject())
	var obj models.Object
	for i := 0; i < b.N; i++ {
		_ = ffjson.Unmarshal(buf, obj)
	}
}

func BenchmarkCborMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(getTestObject())
	}
}

func BenchmarkCborUnmarshal(b *testing.B) {
	buf, _ := ffjson.Marshal(getTestObject())
	var obj models.Object
	for i := 0; i < b.N; i++ {
		_ = cbor.Unmarshal(buf, obj)
	}
}
