package service

import (
	"d2d-backend/account"
	"github.com/biezhi/gorm-paginator/pagination"
)

type accountService struct {
	accountRepos account.AccountRepository
}

func NewAccountService(accountRepository account.AccountRepository) account.AccountService {
	return &accountService{accountRepository}
}

func (accountService *accountService) CreateNewAccount(newAccount *account.Account) (*account.Account, error) {
	_, err := accountService.accountRepos.Create(newAccount)
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func (accountService *accountService) DeleteAccount(id int) (bool, error) {
	bool, err := accountService.accountRepos.Delete(id)
	if err != nil {
		return bool, err
	}
	return bool, nil
}

func (accountService *accountService) GetAccountById(id int) (*account.Account, error) {
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

func (accountService *accountService) UpdateAccount(updateAccount *account.Account) (*account.Account, error) {
	updateAccount, err := accountService.accountRepos.Update(updateAccount)
	if err != nil {
		return nil, err
	}
	return updateAccount, nil
}



