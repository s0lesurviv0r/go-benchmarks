package set

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir string

func init() {
	dir := os.TempDir()
	tmpDir = filepath.Join(dir, "on-disk-set")
}

func getRand() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

// testAddExists tests basic Add and Exists operations
func testAddExists(t *testing.T, s Set) {
	err := s.Open()
	assert.NoError(t, err)
	defer s.Close()

	testValues := []uint64{
		12345,
		67890,
		111111,
		999999,
	}

	for _, value := range testValues {
		// Ensure value doesn't exist
		exists, err := s.Exists(value)
		assert.NoError(t, err)
		assert.False(t, exists)

		// Add value
		err = s.Add(value)
		assert.NoError(t, err)

		// Ensure the value we just added exists
		exists, err = s.Exists(value)
		assert.NoError(t, err)
		assert.True(t, exists)
	}
}

// testDelete tests Delete operation
func testDelete(t *testing.T, s Set) {
	err := s.Open()
	assert.NoError(t, err)
	defer s.Close()

	testValue := uint64(54321)

	// Add value
	err = s.Add(testValue)
	assert.NoError(t, err)

	// Ensure value exists
	exists, err := s.Exists(testValue)
	assert.NoError(t, err)
	assert.True(t, exists)

	// Delete value
	err = s.Delete(testValue)
	// Note: CuteDB doesn't support delete, so we allow unsupported error
	if err != nil && err.Error() != "unsupported operation" {
		assert.NoError(t, err)
	}

	// If delete succeeded, verify it's gone
	if err == nil {
		exists, err = s.Exists(testValue)
		assert.NoError(t, err)
		assert.False(t, exists)
	}
}

// testMultipleOperations tests a sequence of operations
func testMultipleOperations(t *testing.T, s Set) {
	err := s.Open()
	assert.NoError(t, err)
	defer s.Close()

	// Add multiple values
	values := []uint64{100, 200, 300, 400, 500}
	for _, v := range values {
		err = s.Add(v)
		assert.NoError(t, err)
	}

	// Verify all values exist
	for _, v := range values {
		exists, err := s.Exists(v)
		assert.NoError(t, err)
		assert.True(t, exists)
	}

	// Verify non-existent value returns false
	exists, err := s.Exists(999)
	assert.NoError(t, err)
	assert.False(t, exists)
}

func test(t *testing.T, createSet func() Set) {
	// Each test gets its own instance
	t.Run("AddExists", func(t *testing.T) {
		testAddExists(t, createSet())
	})
	t.Run("Delete", func(t *testing.T) {
		testDelete(t, createSet())
	})
	t.Run("MultipleOperations", func(t *testing.T) {
		testMultipleOperations(t, createSet())
	})
}

func TestInMemoryMap(t *testing.T) {
	test(t, func() Set { return NewInMemoryMap() })
}

func TestBadgerDB(t *testing.T) {
	// Clean up any existing test database
	os.RemoveAll(tmpDir + "/badger-test")
	test(t, func() Set { return NewBadgerDB(tmpDir + "/badger-test") })
}

func TestBuntDB(t *testing.T) {
	// Clean up any existing test database
	os.Remove(tmpDir + "/buntdb-test.db")
	test(t, func() Set { return NewBuntDB(tmpDir + "/buntdb-test.db") })
}

func TestSQLite3(t *testing.T) {
	// Clean up any existing test database
	os.Remove(tmpDir + "/sqlite3-test.db")
	test(t, func() Set { return NewSQLite3(tmpDir + "/sqlite3-test.db") })
}

func TestRocksDB(t *testing.T) {
	// Clean up any existing test database
	os.RemoveAll(tmpDir + "/rocksdb-test")
	test(t, func() Set { return NewRocksDB(tmpDir + "/rocksdb-test") })
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
			_, _ = s.Exists(getRand())
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

func BenchmarkBuntDB(b *testing.B) {
	benchmark(b, NewBuntDB(tmpDir+"/buntdb.db"))
}

func BenchmarkSQLite3(b *testing.B) {
	benchmark(b, NewSQLite3(tmpDir+"/sqlite3.db"))
}

func BenchmarkRocksDB(b *testing.B) {
	benchmark(b, NewRocksDB(tmpDir+"/rocksdb"))
}
