package entity

type Account struct {
	Id      string
	Balance float64
}

func NewAccount(id string, balance float64) *Account {
	return &Account{
		Id:      id,
		Balance: balance,
	}
}
