package notification

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type NotificationRepository interface {
	Find(id int) (*models.Notification, error)
	FindByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(notification *models.Notification) (*models.Notification, error)
	Update(notification *models.Notification) (*models.Notification, error)
	Delete(id int) (bool, error)
	GetUnreadNotificationCount(userId int)(int,error)
}
