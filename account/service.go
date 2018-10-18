package account

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type AccountService interface {
	CreateNewAccount(newAccount *models.Account)(*models.Account, error)
	GetAccounts(limit int, page int)(*pagination.Paginator, error)
	GetAccountById(id int)(*models.Account, error)
	UpdateAccount(updateAccount *models.Account)(*models.Account, error)
	DeleteAccount(id int)(bool, error)
	FindOneAccount(condition interface{})(*models.Account,error)
	UpdateAccountFcmToken(accountID int,fcmToken string) error
}
