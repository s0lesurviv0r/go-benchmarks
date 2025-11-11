package set

import (
	"encoding/binary"
	"errors"

	cutedb "github.com/naqvijafar91/cuteDB"
)

// CuteDB uses a B-Tree on disk to back a set
type CuteDB struct {
	path string
	db   *cutedb.DB
}

func NewCuteDB(path string) Set {
	return &CuteDB{
		path: path,
	}
}

func (s *CuteDB) Open() error {
	db, err := cutedb.Open(s.path)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *CuteDB) Close() error {
	return nil
}

func (s *CuteDB) Add(v uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return s.db.Put(string(b), "")
}

func (s *CuteDB) Delete(v uint64) error {
	// CuteDB library does not support deletion operations
	// The underlying B-tree implementation would need to be extended
	// to support this functionality
	return errors.ErrUnsupported
}

func (s *CuteDB) Exists(v uint64) (bool, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	_, exists, err := s.db.Get(string(b))
	return exists, err
}
