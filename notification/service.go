package notification

import "github.com/biezhi/gorm-paginator/pagination"

type NotificationService interface {
	GetNotification(name string) (*Notification, error)
	CreateNewNotification(newNotification *Notification) (*Notification, error)
	GetNotifications(limit int, page int) (*pagination.Paginator, error)
	GetNotificationById(id int) (*Notification, error)
	UpdateNotification(updateNotification *Notification) (*Notification, error)
	DeleteNotification(id int) (bool, error)
	GetNotificationByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	GetTotalUnreadNotification(userId int)(int,error)
}
