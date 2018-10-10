package repository

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/review"
	"d2d-backend/common"

	"errors"
	"github.com/biezhi/gorm-paginator/pagination"
	"strconv"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlReviewRepository() review.ReviewRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id string) (*review.Review, error) {
	if id == ""{
		return nil, errors.New("Invalid id")
	}

	var review review.Review
	idNum,err := strconv.Atoi(id)
	if err != nil {
		return nil,err
	}
	r.Conn.First(&review,idNum)

	return &review, nil
}

func (r *repo) FindReviewByRate(rate int) (*review.Review, error){
	return nil,nil
}


func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error){
	var reviews []review.Review

	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &reviews)

	return paginator,nil
}

func (r *repo)Create(review *review.Review) (*review.Review,error){
	err := r.Conn.Create(review).Error
	if err != nil {
		return nil,err
	}

	return review,nil
}

func (r *repo)Update(updatedReview *review.Review) error {

	var reviewTemp review.Review

	err := r.Conn.First(&reviewTemp,updatedReview.ID).Error

	if err != nil{
		return err
	}

	err = r.Conn.Save(updatedReview).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(id int) (bool,error){

	var reviewTemp review.Review

	err := r.Conn.First(&reviewTemp,id).Error

	if err != nil {
		return false,err
	}

	err = r.Conn.Delete(&reviewTemp).Error

	if err != nil {
		return false,err
	}

	return true,nil
}