package service

import (
	"d2d-backend/user"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type userService struct {
	userRepos	user.UserRepository
}

func NewUserService(UserRepository user.UserRepository) user.UserService {
	return &userService{UserRepository}
}

func (userService *userService) GetUserDetailByName(name string) (*models.User, error) {
	panic("implement me")
}

func (userService *userService) CreateNewUser(newUser *models.User) (*models.User, error) {
	_,err := userService.userRepos.Create(newUser)

	if err != nil {
		return nil,err
	}
	return newUser,nil
}

func (userService *userService) GetUser(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := userService.userRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (userService *userService) GetUserById(id int) (*models.User, error) {
	userModel,err := userService.userRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return userModel,nil
}

func (userService *userService) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	userModel,err := userService.userRepos.FindUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil,err
	}

	return userModel,nil
}

func (userService *userService) UpdateUser(updateUser *models.User) (*models.User, error) {
	updateUser,err := userService.userRepos.Update(updateUser)
	if err != nil {
		return nil,err
	}
	return updateUser,nil
}

func (userService *userService) DeleteUser(id int) (bool, error) {
	isDeletedSuccess,err := userService.userRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess,err
	}
	return isDeletedSuccess,nil
}

