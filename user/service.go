package user

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type UserService interface {
	GetUserDetailByName(name string) (*models.User,error)
	CreateNewUser(newUser *models.User)(*models.User,error)
	GetUser(limit int, page int)(*pagination.Paginator,error)
	GetUserById(id int)(*models.User,error)
	UpdateUser(updateUser *models.User)(*models.User,error)
	DeleteUser(id int)(bool,error)
	GetUserByPhoneNumber(phoneNumber string) (*models.User, error)
	GetUserByAccountId(accountId uint)(*models.User,error)
	SaveNewShippingLocation(shippingLocation *models.UserShippingLocation)(*models.UserShippingLocation,error)
	UpdateUserLocation(shippingLocation *models.UserShippingLocation)(*models.UserShippingLocation,error)
	DeleteUserLocation(id uint) (bool,error)
}

