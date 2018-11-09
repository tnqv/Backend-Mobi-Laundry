package models

import "d2d-backend/models"

type NotificationMessage struct {
	Order models.PlacedOrder
	Tokens []string
}