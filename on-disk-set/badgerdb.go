package set

import (
	"encoding/binary"

	"github.com/outcaste-io/badger/v3"
)

// BadgerDB uses the DB to back a set
type BadgerDB struct {
	path string
	db   *badger.DB
}

func NewBadgerDB(path string) Set {
	return &BadgerDB{
		path: path,
	}
}

func (s *BadgerDB) Open() error {
	opts := badger.DefaultOptions(s.path).
		WithLoggingLevel(badger.ERROR)

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *BadgerDB) Close() error {
	return s.db.Close()
}

func (s *BadgerDB) Add(v uint64) error {
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)

	err := txn.SetEntry(badger.NewEntry(b, []byte{}))
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *BadgerDB) Delete(v uint64) error {
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)

	err := txn.Delete(b)
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *BadgerDB) Exists(v uint64) (bool, error) {
	err := s.db.View(func(txn *badger.Txn) error {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, v)

		_, err := txn.Get(b)
		return err
	})

	if err != nil {
		if err == badger.ErrKeyNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
