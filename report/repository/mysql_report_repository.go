package repository

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/report"
	"d2d-backend/common"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlReportRepository() report.ReportRepository{
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.Report, error){
	var reportModel models.Report
	err := r.Conn.First(&reportModel,id).Error
	if err != nil {
		return nil,err
	}
	return &reportModel, nil
}

func (r *repo) FindAll(limit int,page int) (*common.Paginator, error){
	var listReports []*models.Report
	paginator := common.Pagging(&common.Param{
		DB: r.Conn ,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &listReports)
	return paginator,nil
}

func (r *repo) FindUnresolvedReports(limit int,page int) (*common.Paginator, error){
	var listReports []*models.Report
	paginator := common.Pagging(&common.Param{
		DB: r.Conn.Where("is_resolved = ?",false) ,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &listReports)
	return paginator,nil
}

func (r *repo) Create(report *models.Report) (*models.Report,error){
	err := r.Conn.Create(&report).Error
	if err != nil {
		return nil,err
	}
	return report,nil
}

func (r *repo) Update(report *models.Report) (*models.Report, error){
	var tempCategory models.Report
	err := r.Conn.First(&tempCategory,report.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Model(&report).Update(&report).Error
	if err != nil {
		return nil, err
	}
	return report,nil
}

func (r *repo) ResolveReport(id int)(*models.Report, error){
	var tempReport models.Report
	err := r.Conn.First(&tempReport,id).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Model(&tempReport).Update("is_resolved",true).Error
	if err != nil {
		return nil, err
	}
	return &tempReport,nil
}

func (r *repo) Delete(id int) (bool,error){
	var tempReport models.Report
	err := r.Conn.First(&tempReport, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempReport).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repo) FindReportByPlacedOrderId(page int,limit int, orderId uint)(*common.Paginator,error){
	var listReports []*models.Report
	paginator := common.Pagging(&common.Param{
		DB: r.Conn.Where("placed_order_id = ?",orderId) ,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &listReports)
	return paginator,nil
}