package common

import (
	"github.com/adjust/rmq"
)

type LddsMessageQueue struct {
		rmq.Queue
}

var firebaseMQ rmq.Queue
var notificationMQ rmq.Queue

const (
	FIREBASE_QUEUE = "firebase"
	NOTIFICATION_QUEUE = "notification"
	REDIS_TAG = "produccer"
	QUEUE_NETWORK = "tcp"
)

func InitMessageQueueConnection(redisAddr string) {

	connection := rmq.OpenConnection(REDIS_TAG, QUEUE_NETWORK, redisAddr, 2)

	firebaseMessageQueue := connection.OpenQueue(FIREBASE_QUEUE)
	notificationMessageQueue := connection.OpenQueue(NOTIFICATION_QUEUE)
	firebaseMQ = firebaseMessageQueue
	notificationMQ = notificationMessageQueue

}

func GetNotificationMQ() rmq.Queue {
	return notificationMQ
}

func GetFirebaseMQ() rmq.Queue {
	return firebaseMQ
}