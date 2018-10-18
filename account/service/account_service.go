package service

import (
	"d2d-backend/account"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type accountService struct {
	accountRepos account.AccountRepository
}

func NewAccountService(accountRepository account.AccountRepository) account.AccountService {
	return &accountService{accountRepository}
}

func (accountService *accountService) CreateNewAccount(newAccount *models.Account) (*models.Account, error) {
	_, err := accountService.accountRepos.Create(newAccount)
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func (accountService *accountService) DeleteAccount(id int) (bool, error) {
	isDeletedSuccess, err := accountService.accountRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess, err
	}
	return isDeletedSuccess, nil
}

func (accountService *accountService) GetAccountById(id int) (*models.Account, error) {
	role, err := accountService.accountRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (accountService *accountService) GetAccounts(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := accountService.accountRepos.FindAll(limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (accountService *accountService) UpdateAccount(updateAccount *models.Account) (*models.Account, error) {
	updateAccount, err := accountService.accountRepos.Update(updateAccount)
	if err != nil {
		return nil, err
	}
	return updateAccount, nil
}


func (accountService *accountService) FindOneAccount(condition interface{})(*models.Account,error){
	accountModel,err := accountService.accountRepos.FindOneAccount(condition)
	if err != nil {
		return nil,err
	}
	return &accountModel,nil
}

func (accountService *accountService) UpdateAccountFcmToken(accountID int,fcmToken string) error {
	accountModel,err := accountService.accountRepos.Find(accountID)
	if err != nil {
		return err
	}
	accountModel.FcmToken = fcmToken
	_,err = accountService.accountRepos.Update(accountModel)
	if err != nil {
		return err
	}

	return nil

}

