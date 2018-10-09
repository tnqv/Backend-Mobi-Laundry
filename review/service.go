package review

import "github.com/biezhi/gorm-paginator/pagination"

type ReviewService interface {
	GetReviews(page int,limit int) (*pagination.Paginator,error)
	CreateNewReview(newReview *Review)(*Review,error)
}

