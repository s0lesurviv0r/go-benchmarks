package set

type Set interface {
	Open() error
	Close() error
	Add(uint64) error
	Delete(uint64) error
	Exists(uint64) (bool, error)
}
