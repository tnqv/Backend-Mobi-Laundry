package report

import (
	"d2d-backend/common"
	"d2d-backend/models"
)

type ReportService interface {
	CreateNewReport(newPlacedOrder *models.Report)(*models.Report, error)
	GetAllUnresolvedReports(limit int, page int)(*common.Paginator, error)
	ResolveReport(id int)(*models.Report,error)
}
