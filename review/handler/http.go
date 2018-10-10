package handler

import (
	"github.com/gin-gonic/gin"
	"d2d-backend/review"
	"strconv"
	"d2d-backend/common"
	"net/http"
	"strings"
	"errors"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpReviewHandler struct {
	reviewService review.ReviewService
}

func NewReviewHttpHandler(e *gin.RouterGroup, service review.ReviewService)(*HttpReviewHandler) {
	handler := &HttpReviewHandler{
		reviewService: service,
	}

	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpReviewHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllReviews)

}

func (s *HttpReviewHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.POST("/", s.CreateReview)
	e.GET("/:id", s.GetReviewsById)
	e.PUT("/:id",s.UpdateReview)
	e.DELETE("/:id", s.DeleteReview)
}

func (s *HttpReviewHandler) GetAllReviews(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listReview, err := s.reviewService.GetReviews(page,limit)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK,listReview)
}

func  (s *HttpReviewHandler) GetReviewsById(c *gin.Context){

}

func  (s *HttpReviewHandler) CreateReview(c *gin.Context){

	content := c.PostForm("content")

	if content == "" || strings.TrimSpace(content) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty content",errors.New("Content not valid")))
		return
	}

	rate := c.PostForm("rate")

	if rate == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty rate",errors.New("Rate is empty")))
		return
	}

	rateNum,err := strconv.Atoi(rate)

	if err != nil {
		c.JSON(http.StatusNotAcceptable,common.NewError("Rate is not valid", errors.New("Rate is not valid")))
		return
	}

	userId, err := strconv.Atoi(c.PostForm("user_id"))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("Invalid user_id", err))
		return
	}

	var newReview review.Review

	newReview.UserID = userId
	newReview.Content = content
	newReview.Rate = rateNum

	_,err = s.reviewService.CreateNewReview(&newReview)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Invalid user", err))
		return
	}

	c.JSON(http.StatusOK,newReview)
}

func  (s *HttpReviewHandler) UpdateReview(c *gin.Context){

}

func (s *HttpReviewHandler) DeleteReview(c *gin.Context){

}