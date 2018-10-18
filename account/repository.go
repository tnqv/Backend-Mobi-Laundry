package account

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type AccountRepository interface{
	Find(id int) (*models.Account, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(account *models.Account) (*models.Account,error)
	Update(account *models.Account) (*models.Account, error)
	Delete(id int) (bool,error)
	FindOneAccount(condition interface{})(models.Account,error)
}
