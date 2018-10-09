package service

import (
	"d2d-backend/review"
	"github.com/biezhi/gorm-paginator/pagination"
)

type reviewService struct {
	reviewRepos   review.ReviewRepository
}

func NewReviewService(reviewRepository review.ReviewRepository) review.ReviewService {
	return &reviewService{reviewRepository}
}

func (reviewService *reviewService) GetReviews(page int,limit int) (*pagination.Paginator,error){
	var listReviewsPaginator *pagination.Paginator

	listReviewsPaginator,err := reviewService.reviewRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return listReviewsPaginator,nil
}

func (reviewService *reviewService) CreateNewReview(newReview *review.Review)(*review.Review,error){
	_,err := reviewService.reviewRepos.Create(newReview)

	if err != nil {
		return nil,err
	}

	return newReview,nil
}

