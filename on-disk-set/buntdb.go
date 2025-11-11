package set

import (
	"strconv"

	"github.com/tidwall/buntdb"
)

// BuntDB uses the DB to back a set
type BuntDB struct {
	path string
	db   *buntdb.DB
}

func NewBuntDB(path string) Set {
	return &BuntDB{
		path: path,
	}
}

func (s *BuntDB) Open() error {
	db, err := buntdb.Open(s.path)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *BuntDB) Close() error {
	return s.db.Close()
}

func (s *BuntDB) Add(v uint64) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		key := strconv.FormatUint(v, 10)
		_, _, err := tx.Set(key, "", nil)
		return err
	})
}

func (s *BuntDB) Delete(v uint64) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		key := strconv.FormatUint(v, 10)
		_, err := tx.Delete(key)
		if err == buntdb.ErrNotFound {
			return nil
		}
		return err
	})
}

func (s *BuntDB) Exists(v uint64) (bool, error) {
	var exists bool
	err := s.db.View(func(tx *buntdb.Tx) error {
		key := strconv.FormatUint(v, 10)
		_, err := tx.Get(key)
		if err != nil {
			if err == buntdb.ErrNotFound {
				exists = false
				return nil
			}
			return err
		}
		exists = true
		return nil
	})
	return exists, err
}
