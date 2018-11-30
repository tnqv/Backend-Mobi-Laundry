package report

import (
	"d2d-backend/common"
	"d2d-backend/models"
)

type ReportRepository interface{
	Find(id int) (*models.Report, error)
	FindAll(limit int,page int) (*common.Paginator, error)
	Create(report *models.Report) (*models.Report,error)
	Update(report *models.Report) (*models.Report, error)
	Delete(id int) (bool,error)
	FindReportByPlacedOrderId(page int,limit int, orderId uint)(*common.Paginator,error)
	FindUnresolvedReports(limit int,page int) (*common.Paginator, error)
	ResolveReport(id int)(*models.Report, error)
}
