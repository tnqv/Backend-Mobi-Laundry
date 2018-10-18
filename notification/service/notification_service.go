package service

import (
	"d2d-backend/notification"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type notificationService struct {
	notificationRepos	notification.NotificationRepository
}

func NewNotificationService(notificationRepository notification.NotificationRepository) notification.NotificationService {
	return &notificationService{notificationRepository}
}

func (notificationService *notificationService) GetNotification(name string) (*models.Notification, error) {
	panic("implement me")
}

func (notificationService *notificationService) CreateNewNotification(newNotification *models.Notification) (*models.Notification, error) {
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

func (notificationService *notificationService) GetNotificationById(id int) (*models.Notification, error) {
	notificationModel,err := notificationService.notificationRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return notificationModel,nil
}

func (notificationService *notificationService) UpdateNotification(updateNotification *models.Notification) (*models.Notification, error) {
	updateNotification,err := notificationService.notificationRepos.Update(updateNotification)
	if err != nil {
		return nil,err
	}
	return updateNotification,nil
}

func (notificationService *notificationService) DeleteNotification(id int) (bool, error) {
	isDeletedSuccess,err := notificationService.notificationRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess,err
	}
	return isDeletedSuccess,nil
}

func (notificationService *notificationService) GetNotificationByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	paginate,err := notificationService.notificationRepos.FindByUserId(limit, page, id)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}


func (notificationService *notificationService) GetTotalUnreadNotification(userId int)(int,error){
	count,err := notificationService.notificationRepos.GetUnreadNotificationCount(userId)
	if err != nil {
		return 0,err
	}
	return count,nil
}