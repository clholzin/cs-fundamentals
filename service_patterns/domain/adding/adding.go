package adding

import "cs-fundamentals/service_patterns/domain/dictionary"

type Service struct {
	db dictionary.Repository
}

func NewService(repo dictionary.Repository) *Service {
	s := &Service{}
	s.db = repo
	return s
}

func (s *Service) Add(word dictionary.Word) error {
	return s.db.Add(word.Key, word.Value)
}
