package set

import (
	"encoding/binary"

	"github.com/linxGnu/grocksdb"
)

// RocksDB uses RocksDB to back a set
type RocksDB struct {
	path string
	db   *grocksdb.DB
	ro   *grocksdb.ReadOptions
	wo   *grocksdb.WriteOptions
}

func NewRocksDB(path string) Set {
	return &RocksDB{
		path: path,
	}
}

func (s *RocksDB) Open() error {
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	opts.SetErrorIfExists(false)

	db, err := grocksdb.OpenDb(opts, s.path)
	if err != nil {
		return err
	}

	s.db = db
	s.ro = grocksdb.NewDefaultReadOptions()
	s.wo = grocksdb.NewDefaultWriteOptions()

	return nil
}

func (s *RocksDB) Close() error {
	if s.wo != nil {
		s.wo.Destroy()
	}
	if s.ro != nil {
		s.ro.Destroy()
	}
	if s.db != nil {
		s.db.Close()
	}
	return nil
}

func (s *RocksDB) Add(v uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)

	err := s.db.Put(s.wo, b, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func (s *RocksDB) Delete(v uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)

	err := s.db.Delete(s.wo, b)
	if err != nil {
		return err
	}

	return nil
}

func (s *RocksDB) Exists(v uint64) (bool, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)

	slice, err := s.db.Get(s.ro, b)
	if err != nil {
		return false, err
	}
	defer slice.Free()

	if !slice.Exists() {
		return false, nil
	}

	return true, nil
}
