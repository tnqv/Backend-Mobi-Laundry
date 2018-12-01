package repository

import (
	"d2d-backend/common"
	"d2d-backend/notification"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}


func NewMysqlNotificationRepository() notification.NotificationRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.Notification, error) {
	var notificationModel models.Notification
	err := r.Conn.First(&notificationModel, id).Error
	if err != nil {
		return nil, err
	}
	return &notificationModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var notifications []*models.Notification
	paginator := pagination.Pagging(&pagination.Param{
		DB:      r.Conn,
		Page:    page,
		Limit:   limit,
		ShowSQL: true,
	}, &notifications)

	return paginator, nil
}

func (r *repo) Create(notification *models.Notification) (*models.Notification, error) {
	notification.DeletedAt = nil
	err := r.Conn.Create(notification).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *repo) Update(updateNotification *models.Notification) (*models.Notification, error) {
	var tempNotification models.Notification
	err := r.Conn.First(&tempNotification, updateNotification.ID).Error
	if err != nil {
		return nil, err
	}
	updateNotification.DeletedAt = nil
	err = r.Conn.Save(updateNotification).Error
	if err != nil {
		return nil, err
	}
	return updateNotification, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempNotification models.Notification
	err := r.Conn.First(&tempNotification, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempNotification).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repo) FindByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	var notifications []*models.Notification
	db := r.Conn
	db = db.Where("user_id = ?", id)
	paginator := pagination.Pagging(&pagination.Param{
		DB: db,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &notifications)
	return paginator,nil
}

func (r *repo) GetUnreadNotificationCount(userId int)(int,error){
	var count int
	err := r.Conn.Model(&models.Notification{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count,nil
}