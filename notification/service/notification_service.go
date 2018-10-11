package service

import (
	"d2d-backend/notification"
	"github.com/biezhi/gorm-paginator/pagination"
)

type notificationService struct {
	notificationRepos	notification.NotificationRepository
}

func NewNotificationService(notificationRepository notification.NotificationRepository) notification.NotificationService {
	return &notificationService{notificationRepository}
}

func (notificationService *notificationService) GetNotification(name string) (*notification.Notification, error) {
	panic("implement me")
}

func (notificationService *notificationService) CreateNewNotification(newNotification *notification.Notification) (*notification.Notification, error) {
	_,err := notificationService.notificationRepos.Create(newNotification)
	if err != nil {
		return nil,err
	}
	return newNotification,nil
}

func (notificationService *notificationService) GetNotifications(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := notificationService.notificationRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (notificationService *notificationService) GetNotificationById(id int) (*notification.Notification, error) {
	notification,err := notificationService.notificationRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return notification,nil
}

func (notificationService *notificationService) UpdateNotification(updateNotification *notification.Notification) (*notification.Notification, error) {
	updateNotification,err := notificationService.notificationRepos.Update(updateNotification)
	if err != nil {
		return nil,err
	}
	return updateNotification,nil
}

func (notificationService *notificationService) DeleteNotification(id int) (bool, error) {
	bool,err := notificationService.notificationRepos.Delete(id)
	if err != nil {
		return bool,err
	}
	return bool,nil
}
