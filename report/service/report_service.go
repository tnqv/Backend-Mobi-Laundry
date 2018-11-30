package service

import (
	"d2d-backend/report"
	"d2d-backend/models"
	"d2d-backend/common"
)

type reportService struct {
	reportRepos report.ReportRepository

}

func NewReportService(reportRepository report.ReportRepository) report.ReportService{
	return &reportService{reportRepository}
}

func (s *reportService) CreateNewReport(newReport *models.Report)(*models.Report, error){
	newReport,err := s.reportRepos.Create(newReport)

	if err != nil {
		return nil,err
	}

	return newReport,nil
}

func (s *reportService) GetAllUnresolvedReports(limit int, page int)(*common.Paginator, error){
	results,err := s.reportRepos.FindUnresolvedReports(limit,page)
	if err != nil {
		return nil,err
	}
	return results,nil
}

func (s *reportService) ResolveReport(id int)(*models.Report,error){
	report,err := s.reportRepos.ResolveReport(id)
	if err != nil {
		return nil,err
	}
	return report,nil
}