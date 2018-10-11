package account

import "github.com/biezhi/gorm-paginator/pagination"

type AccountService interface {
	CreateNewAccount(newAccount *Account)(*Account, error)
	GetAccounts(limit int, page int)(*pagination.Paginator, error)
	GetAccountById(id int)(*Account, error)
	UpdateAccount(updateAccount *Account)(*Account, error)
	DeleteAccount(id int)(bool, error)
}
