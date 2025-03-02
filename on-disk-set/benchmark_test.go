package set

import (
	"io/ioutil"
	"math/rand"
	"testing"
)

var tmpDir string

func init() {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		panic(err)
	}
	tmpDir = dir
}

func getRand() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

func preloadSet(s Set, size int) {
	for i := 0; i < size; i++ {
		s.Add(getRand())
	}
}

func benchmark(b *testing.B, s Set) {
	s.Open()
	preloadSet(s, 100000)

	b.Run("Exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			exists, _ := s.Exists(getRand())
			if exists {
			}
		}
	})

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Add(getRand())
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Delete(getRand())
		}
	})
}

func BenchmarkMap(b *testing.B) {
	benchmark(b, NewMap())
}

func BenchmarkBadgerDB(b *testing.B) {
	benchmark(b, NewBadgerDB(tmpDir+"/badger"))
}

func BenchmarkCuteDB(b *testing.B) {
	benchmark(b, NewCuteDB(tmpDir+"/cutedb"))
}
