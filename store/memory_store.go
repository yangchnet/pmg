package store

import "pmg/pass"

var _ Store = (*MemoryStore)(nil)

// MemoryStore store all token in memory
type MemoryStore struct {
	store map[string]*pass.Token
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]*pass.Token),
	}
}

// Set save token by key, error when failed
func (s *MemoryStore) Set(key, passwd string, values ...any) error {
	t, _ := s.store[key]
	s.store[key] = t.SetToken(passwd, values...)
	return nil
}

// Get get token by key
func (s *MemoryStore) Get(key string) (*pass.Token, error) {
	if t, ok := s.store[key]; ok {
		return t, nil
	}

	return nil, ErrKeyNotFound
}

// Persistence save all data somewhere permanently
func (s *MemoryStore) Persistence() error {
	return nil
}
