package user

import "github.com/biezhi/gorm-paginator/pagination"

type UserRepository interface{
	Find(id int) (*User, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *User) (*User,error)
	Update(category *User) (*User, error)
	Delete(id int) (bool,error)
}
