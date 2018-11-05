package common

import (
	"google.golang.org/grpc"
)

type Notification struct {
	cc *grpc.ClientConn
}

const (
	MESSAGE_PATTERN_STATUS_1 = `Bạn vừa tạo đơn hàng #%s`
	MESSAGE_PATTERN_STATUS_2 = `Đơn hàng #%s của bạn đã được cửa hàng %s xác nhận `
	MESSAGE_PATTERN_STATUS_3 = `Shipper %s đã tiếp nhận đơn hàng`
	MESSAGE_PATTERN_STATUS_4 = `Đơn hàng #%s đã được xác nhận dịch vụ giặt`
	MESSAGE_PATTERN_STATUS_5 = `Đơn hàng #%s đã được vận chuyển tới cửa hàng`
	MESSAGE_PATTERN_STATUS_6 = `Đơn hàng #%s đang trong quá trình giặt`
	MESSAGE_PATTERN_STATUS_7 = `Đơn hàng #%s của bạn đã hoàn thành`
	MESSAGE_PATTERN_STATUS_8 = `Đơn hàng #%scủa bạn đang được đem trả`
	MESSAGE_PATTERN_STATUS_9 = `Hoàn thành đơn hàng`
	//MESSAGE_PATTERN_STATUS_10
)

const (
	TITLE_MESSAGE_PATTERN = `Cập nhật trạng thái đơn hàng`
)

var Con *grpc.ClientConn

func InitNotificationConnection(notificationAddr string) *grpc.ClientConn {
	conn, err := grpc.Dial(notificationAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//c := proto.NewGorushClient(conn)
	Con = conn
	return Con
}

// Using this function to get a connection, you can create your connection pool here.
func GetNotificationConnection() *grpc.ClientConn {
	return Con
}
