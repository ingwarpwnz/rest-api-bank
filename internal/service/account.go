package service

import (
	"fmt"

	"github.com/ingwarpwnz/rest-api-bank/internal/repository/entity"
)

type AccountRepositoryInterface interface {
	Create(balance float64) (*entity.Account, error)
	FindById(id string) *entity.Account
	UpdateBalance(id string, balance float64) *entity.Account
}

type AccountService struct {
	repo AccountRepositoryInterface
	tx   chan struct{}
}

func NewAccountService(repo AccountRepositoryInterface) *AccountService {
	return &AccountService{
		repo: repo,
		tx:   make(chan struct{}, 1),
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
	s.tx <- struct{}{}
	defer func() {
		<-s.tx
	}()

	sender := s.repo.FindById(from)
	if sender == nil {
		return nil, nil, fmt.Errorf("sender account not found")
	}

	senderNewBalance := sender.Balance - amount
	if senderNewBalance < 0 {
		return nil, nil, fmt.Errorf("the sender has insufficient funds")
	}

	recipient := s.repo.FindById(to)
	if recipient == nil {
		return nil, nil, fmt.Errorf("recipient account not found")
	}

	accountFrom = s.repo.UpdateBalance(sender.Id, senderNewBalance)
	accountTo = s.repo.UpdateBalance(recipient.Id, recipient.Balance+amount)

	return accountFrom, accountTo, nil
}
