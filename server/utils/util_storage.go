package utils

import "sync"

type Storage struct {
	storage map[string]struct{}
	mutex   sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		storage: make(map[string]struct{}),
	}
}

func (s *Storage) Add(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.storage[key] = struct{}{}
}

func (s *Storage) Check(key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.storage[key]
	if ok {
		delete(s.storage, key)
	}
	return ok
}

func (s *Storage) Remove(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.storage, key)
}
