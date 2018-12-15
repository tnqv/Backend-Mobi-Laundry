package handler

import (
	"github.com/gin-gonic/gin"
	"d2d-backend/review"
	"strconv"
	"d2d-backend/common"
	"net/http"
	"strings"
	"errors"
	"d2d-backend/models"
	"d2d-backend/placedOrder"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpReviewHandler struct {
	reviewService review.ReviewService
	placedOrderService placedOrder.PlacedOrderService
}

func NewReviewHttpHandler(e *gin.RouterGroup, service review.ReviewService,placedOrderService placedOrder.PlacedOrderService)(*HttpReviewHandler) {
	handler := &HttpReviewHandler{
		reviewService: service,
		placedOrderService: placedOrderService,
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
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty content",errors.New("Chưa nhập nội dung")))
		return
	}

	rate := c.PostForm("rate")

	if rate == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty rate",errors.New("Chưa nhập điểm ")))
		return
	}

	rateNum,err := strconv.Atoi(rate)

	if err != nil {
		c.JSON(http.StatusNotAcceptable,common.NewError("Rate is not valid", errors.New("Điểm không hợp lệ")))
		return
	}

	userId, err := strconv.Atoi(c.PostForm("user_id"))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("Mã người dùng không hợp lệ", err))
		return
	}

	placedOrderId,err := strconv.Atoi(c.PostForm("placed_order_id"))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("Mã hoá đơn không hợp lệ", err))
		return
	}

	placedOrder,err := s.placedOrderService.GetPlacedOrderById(placedOrderId)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("Mã hoá đơn không tồn tại", err))
		return
	}

	if placedOrder.ReviewID != 0 {
		c.JSON(http.StatusNotAcceptable, common.NewError("Hoá đơn không thể đánh giá", err))
		return
	}


	var newReview models.Review

	newReview.UserID = userId
	newReview.Content = content
	newReview.Rate = rateNum

	_,err = s.reviewService.CreateNewReview(&newReview)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Lỗi khi tạo đánh giá", err))
		return
	}

	if newReview.ID != 0 {
		placedOrder.ReviewID = newReview.ID
		_,err := s.placedOrderService.UpdatePlacedOrder(placedOrder)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("Lỗi cập nhật đánh giá cho đơn hàng", err))
			return
		}
	}

	c.JSON(http.StatusOK,newReview)
}

func  (s *HttpReviewHandler) UpdateReview(c *gin.Context){

}

func (s *HttpReviewHandler) DeleteReview(c *gin.Context){

}