package db

import (
	"sync"

	"github.com/ingwarpwnz/rest-api-bank/internal/repository/entity"
)

type MemoryStorage struct {
	mu   sync.RWMutex
	data map[string]*entity.Account
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]*entity.Account, 100),
	}
}

func (s *MemoryStorage) Insert(acc *entity.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[acc.Id] = acc
	return nil
}

func (s *MemoryStorage) GetById(id string) *entity.Account {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if acc, ok := s.data[id]; ok {
		return acc
	}
	return nil
}

// UpdateBalance fake sql update request
func (s *MemoryStorage) UpdateBalance(id string, amount float64) *entity.Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	if acc, ok := s.data[id]; ok {
		acc.Balance = amount
		return acc
	}

	return nil
}
