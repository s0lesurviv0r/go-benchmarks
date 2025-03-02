package set

type InMemoryMap struct {
	m map[uint64]struct{}
}

func NewInMemoryMap() Set {
	return &InMemoryMap{
		m: make(map[uint64]struct{}),
	}
}

func (s *InMemoryMap) Open() error {
	return nil
}

func (s *InMemoryMap) Close() error {
	s.m = nil
	return nil
}

func (s *InMemoryMap) Add(v uint64) error {
	s.m[v] = struct{}{}
	return nil
}

func (s *InMemoryMap) Delete(v uint64) error {
	delete(s.m, v)
	return nil
}

func (s *InMemoryMap) Exists(v uint64) (bool, error) {
	_, exists := s.m[v]
	return exists, nil
}
