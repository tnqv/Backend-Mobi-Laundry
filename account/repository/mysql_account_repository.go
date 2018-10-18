package repository

import (
	"d2d-backend/account"
	"d2d-backend/common"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"errors"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlAccounteRepository() account.AccountRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*account.Account, error) {
	var account account.Account
	err := r.Conn.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var account []*account.Account
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &account)
	return paginator,nil
}

func (r *repo) Create(accountModel *account.Account) (*account.Account, error) {

	var temp account.Account

	if err := r.Conn.Where("email = ?",accountModel.Email).First(&temp).Error; err == nil {
		return nil, errors.New("Email đã có người đăng ký")
	}

	if err := r.Conn.Create(accountModel).Error; err != nil{
		return nil,err
	}

	return accountModel,nil
}

func (r *repo) Update(updateAccount *account.Account) (*account.Account, error) {
	var tempAccount account.Account
	err := r.Conn.First(&tempAccount,updateAccount.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateAccount).Error
	if err != nil {
		return nil, err
	}
	return updateAccount, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempAccount account.Account
	err := r.Conn.First(&tempAccount, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempAccount).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repo) FindOneAccount(condition interface{})(account.Account,error){
	var model account.Account
	err := r.Conn.Where(condition).First(&model).Error
	return model, err
}

