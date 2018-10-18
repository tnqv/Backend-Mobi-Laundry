package notification

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type NotificationService interface {
	GetNotification(name string) (*models.Notification, error)
	CreateNewNotification(newNotification *models.Notification) (*models.Notification, error)
	GetNotifications(limit int, page int) (*pagination.Paginator, error)
	GetNotificationById(id int) (*models.Notification, error)
	UpdateNotification(updateNotification *models.Notification) (*models.Notification, error)
	DeleteNotification(id int) (bool, error)
	GetNotificationByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	GetTotalUnreadNotification(userId int)(int,error)
}
