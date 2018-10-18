package account

import "github.com/biezhi/gorm-paginator/pagination"

type AccountRepository interface{
	Find(id int) (*Account, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(account *Account) (*Account,error)
	Update(account *Account) (*Account, error)
	Delete(id int) (bool,error)
	FindOneAccount(condition interface{})(Account,error)
}
