package service

import (
	"d2d-backend/user"
	"github.com/biezhi/gorm-paginator/pagination"
)

type userService struct {
	userRepos	user.UserRepository
}

func NewUserService(UserRepository user.UserRepository) user.UserService {
	return &userService{UserRepository}
}

func (userService *userService) GetUserDetailByName(name string) (*user.User, error) {
	panic("implement me")
}

func (userService *userService) CreateNewUser(newUser *user.User) (*user.User, error) {
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

func (userService *userService) GetUserById(id int) (*user.User, error) {
	user,err := userService.userRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return user,nil
}

func (userService *userService) UpdateUser(updateUser *user.User) (*user.User, error) {
	updateUser,err := userService.userRepos.Update(updateUser)
	if err != nil {
		return nil,err
	}
	return updateUser,nil
}

func (userService *userService) DeleteUser(id int) (bool, error) {
	bool,err := userService.userRepos.Delete(id)
	if err != nil {
		return bool,err
	}
	return bool,nil
}

