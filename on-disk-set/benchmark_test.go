package set

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

var tmpDir string

func init() {
	dir := os.TempDir()
	tmpDir = filepath.Join(dir, "on-disk-set")
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

func BenchmarkInMemoryMap(b *testing.B) {
	benchmark(b, NewInMemoryMap())
}

func BenchmarkBadgerDB(b *testing.B) {
	benchmark(b, NewBadgerDB(tmpDir+"/badger"))
}

func BenchmarkCuteDB(b *testing.B) {
	benchmark(b, NewCuteDB(tmpDir+"/cutedb"))
}

func BenchmarkBuntDB(b *testing.B) {
	benchmark(b, NewBuntDB(tmpDir+"/buntdb.db"))
}
