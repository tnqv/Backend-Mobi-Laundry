package user

import "github.com/biezhi/gorm-paginator/pagination"

type UserService interface {
	GetUserDetailByName(name string) (*User,error)
	CreateNewUser(newUser *User)(*User,error)
	GetUser(limit int, page int)(*pagination.Paginator,error)
	GetUserById(id int)(*User,error)
	UpdateUser(updateUser *User)(*User,error)
	DeleteUser(id int)(bool,error)
}

