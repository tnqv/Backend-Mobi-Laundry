package review

import "github.com/biezhi/gorm-paginator/pagination"

type ReviewRepository interface{
	Find(id string) (*Review, error)
	FindReviewByRate(rate int) (*Review, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(review *Review) (*Review,error)
	Update(review *Review) error
	Delete(id int) (bool,error)
}