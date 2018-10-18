package review

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ReviewRepository interface{
	Find(id string) (*models.Review, error)
	FindReviewByRate(rate int) (*models.Review, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(review *models.Review) (*models.Review,error)
	Update(review *models.Review) error
	Delete(id int) (bool,error)
}