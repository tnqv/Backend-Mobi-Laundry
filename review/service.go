package review

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ReviewService interface {
	GetReviews(page int,limit int) (*pagination.Paginator,error)
	CreateNewReview(newReview *models.Review)(*models.Review,error)
}

