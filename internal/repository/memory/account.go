package memory

import (
	"github.com/google/uuid"
	"github.com/ingwarpwnz/rest-api-bank/internal/repository/db"
	"github.com/ingwarpwnz/rest-api-bank/internal/repository/entity"
	"github.com/ingwarpwnz/rest-api-bank/internal/service"
)

type AccountRepository struct {
	storage *db.MemoryStorage
}

func NewAccountRepository(storage *db.MemoryStorage) service.AccountRepositoryInterface {
	return &AccountRepository{
		storage: storage,
	}
}

func (r *AccountRepository) Create(balance float64) (*entity.Account, error) {
	id := uuid.New()
	acc := entity.NewAccount(id.String(), balance)
	if err := r.storage.Insert(acc); err != nil {
		return nil, err
	}
	return acc, nil
}

func (r *AccountRepository) FindById(id string) *entity.Account {
	return r.storage.GetById(id)
}

func (r *AccountRepository) UpdateBalance(id string, balance float64) *entity.Account {
	return r.storage.UpdateBalance(id, balance)
}
