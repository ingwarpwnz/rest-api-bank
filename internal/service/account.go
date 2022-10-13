package service

import (
	"fmt"
	"sync"

	"github.com/ingwarpwnz/rest-api-bank/internal/repository/entity"
)

type AccountRepositoryInterface interface {
	Create(balance float64) (*entity.Account, error)
	FindById(id string) *entity.Account
	UpdateBalance(id string, balance float64) *entity.Account
}

type AccountService struct {
	repo AccountRepositoryInterface
}

func NewAccountService(repo AccountRepositoryInterface) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

// Create an account with a balance
func (s *AccountService) Create(balance float64) (*entity.Account, error) {
	return s.repo.Create(balance)
}

// FindById get account balance by its id
func (s *AccountService) FindById(id string) *entity.Account {
	return s.repo.FindById(id)
}

// Transfer funds from one account to another
func (s *AccountService) Transfer(from, to string, amount float64) (accountFrom *entity.Account, accountTo *entity.Account, err error) {
	// transaction simulation
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	f := s.repo.FindById(from)
	if f == nil {
		return nil, nil, fmt.Errorf("sender account not found")
	}

	if f.Balance-amount < 0 {
		return nil, nil, fmt.Errorf("the sender has insufficient funds")
	}

	t := s.repo.FindById(to)
	if t == nil {
		return nil, nil, fmt.Errorf("recipient account not found")
	}

	accountFrom = s.repo.UpdateBalance(f.Id, f.Balance-amount)
	accountTo = s.repo.UpdateBalance(t.Id, t.Balance+amount)

	return accountFrom, accountTo, nil
}
