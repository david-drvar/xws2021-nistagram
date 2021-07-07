package domain

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type LoginRequest struct {
	Email    string
	Password string
}

type Order struct {
	Id          string `gorm:"primaryKey"`
	UserId      string
	ProductId   string
	Quantity    int
	DateCreated time.Time
	TotalPrice  float32
	Referral    int
	Username    string
	ProductName string
}

func (o Order) ConvertToGrpc() *protopb.Order {
	return &protopb.Order{
		Id:          o.Id,
		UserId:      o.UserId,
		ProductId:   o.ProductId,
		Quantity:    int32(o.Quantity),
		Referral:    int32(o.Referral),
		DateCreated: timestamppb.New(o.DateCreated),
		TotalPrice:  o.TotalPrice,
		Username:    o.Username,
		ProductName: o.ProductName,
	}
}
