package persistence

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)
func (notification UserNotification) ConvertToGrpc() *protopb.Notification {
	return &protopb.Notification{
		Id:        notification.NotificationId,
		Text:      notification.Text,
		CreatorId: notification.CreatorId,
		UserId:    notification.UserId,
		Type:      notification.Type,
		ContentId: notification.ContentId,
		IsRead: notification.IsRead,
		CreatedAt: timestamppb.New(notification.CreatedAt),
	}
}

