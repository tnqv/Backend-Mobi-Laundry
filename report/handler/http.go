package handler

import (
	"d2d-backend/report"
	"github.com/gin-gonic/gin"
	"strconv"
	"d2d-backend/common"
	"net/http"
	"d2d-backend/models"
	"strings"
	"errors"
)

type HttpReportHandler struct {
	reportService report.ReportService
}

func NewHttpReportHandler(e *gin.RouterGroup,
	reportServiceParam report.ReportService) (*HttpReportHandler){
	handler := &HttpReportHandler{
		reportService: reportServiceParam,
	}


	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpReportHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpReportHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/unresolved", s.GetAllUnResolvedReport)
	e.POST("/", s.CreateNewReport)
	e.PUT("/resolve/:id",s.ResolveReport)
	//e.DELETE("/:id", s.DeletePlacedOrder)

}



func (s *HttpReportHandler) GetAllUnResolvedReport(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listReport, err := s.reportService.GetAllUnresolvedReports(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listReport)
}

func (s *HttpReportHandler) CreateNewReport(c *gin.Context){

	var newReport models.Report
	err := common.Bind(c, &newReport)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if newReport.Content == "" || strings.TrimSpace(newReport.Content) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Nội dung báo cáo trống")))
		return
	}
	if newReport.PlacedOrderId == 0 {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Mã đơn hàng không hợp lệ")))
		return
	}
	createdReport, err := s.reportService.CreateNewReport(&newReport)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, createdReport)
}

func (s *HttpReportHandler) ResolveReport(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã báo cáo không hợp lệ")))
		return
	}
	idNum, err := strconv.ParseInt(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	report,err := s.reportService.ResolveReport(int(idNum))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", err))
		return
	}
	c.JSON(http.StatusOK, report)

}