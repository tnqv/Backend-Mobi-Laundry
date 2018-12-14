package user

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type UserRepository interface{
	Find(id int) (*models.User, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *models.User) (*models.User,error)
	Update(category *models.User) (*models.User, error)
	Delete(id int) (bool,error)
	FindUserByPhoneNumber(phoneNumber string)(*models.User,error)
	FindUserByAccountId(accountId uint)(*models.User,error)
	SaveNewUserLocation(location *models.UserShippingLocation) (*models.UserShippingLocation,error)
	UpdateUserLocation(location *models.UserShippingLocation) (*models.UserShippingLocation,error)
	DeleteUserLocation(id uint) (bool,error)
}
