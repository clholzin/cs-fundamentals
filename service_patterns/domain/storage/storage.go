package storage

import "errors"

type Storage struct {
	dictionary map[string]string
}

func NewStorage() *Storage {
	store := new(Storage)
	store.dictionary = make(map[string]string)
	return store
}

func (s *Storage) Add(key, value string) error {
	if _, ok := s.dictionary[key]; ok {
		return errors.New("key already exists")
	}
	s.dictionary[key] = value
	return nil
}

func (s *Storage) Search(key string) (result string) {
	return result
}
