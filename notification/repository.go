package notification

import "github.com/biezhi/gorm-paginator/pagination"

type NotificationRepository interface {
	Find(id int) (*Notification, error)
	FindByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(notification *Notification) (*Notification, error)
	Update(notification *Notification) (*Notification, error)
	Delete(id int) (bool, error)
	GetUnreadNotificationCount(userId int)(int,error)
}
