package notification

import "github.com/biezhi/gorm-paginator/pagination"

type NotificationRepository interface {
	Find(id int) (*Notification, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(notification *Notification) (*Notification, error)
	Update(notification *Notification) (*Notification, error)
	Delete(id int) (bool, error)
}
