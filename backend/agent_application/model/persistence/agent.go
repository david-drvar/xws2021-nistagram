package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/agent_application/model"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Product struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Price    float32
	IsActive bool
	Quantity int
	Photo    string
	AgentId  string
}

type PaymentType string

//const (
//	Cash      PaymentType = "Cash"
//	PayPal                = "PayPal"
//	Cripto                = "Cripto"
//	DebitCard             = "DebitCard"
//)

type Order struct {
	Id          string `gorm:"primaryKey"`
	UserId      string
	ProductId   string
	Quantity    int
	DateCreated time.Time
	TotalPrice  float32
	Referral    int
}

type OrderProducts struct {
	OrderId   string `gorm:"primaryKey"`
	ProductId string `gorm:"primaryKey"`
}

type CommonReport struct {
	Id           string `gorm:"primaryKey"`
	CampaignId   string
	PlacementNum int
	LinkClickNum int
	Revenue      float32
}

type PostReport struct {
	CommonReport
	LikesNum    int
	DislikesNum int
	CommentsNum int
}

type StoryReport struct {
	CommonReport
	ViewsNum int
}

type User struct {
	Id           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Username     string `gorm:"unique"`
	Role         model.UserRole
	BirthDate    time.Time
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	Address      string
	Password     string
	IsActive     bool
}

func (u *User) ConvertFromGrpc(user *protopb.UserAgentApp) *User {
	newUser := &User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Username:     user.Username,
		Password:     user.Password,
		Role:         model.GetUserRole(user.Role),
		BirthDate:    user.Birthdate.AsTime(),
		ProfilePhoto: user.ProfilePhoto,
		PhoneNumber:  user.PhoneNumber,
		Sex:          user.Sex,
		IsActive:     user.IsActive,
		Address:      user.Address,
	}

	return newUser
}

func (u User) ConvertToGrpc() *protopb.UserAgentApp {
	return &protopb.UserAgentApp{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Username:     u.Username,
		Password:     "",
		Role:         string(u.Role),
		Birthdate:    timestamppb.New(u.BirthDate),
		ProfilePhoto: u.ProfilePhoto,
		PhoneNumber:  u.PhoneNumber,
		Sex:          u.Sex,
		IsActive:     u.IsActive,
		Address:      u.Address,
	}
}

func (u *Product) ConvertFromGrpc(product *protopb.Product) *Product {
	newProduct := &Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    float32(product.Price),
		IsActive: product.IsActive,
		Quantity: int(product.Quantity),
		Photo:    product.Photo,
		AgentId:  product.AgentId,
	}

	return newProduct
}

func (p Product) ConvertToGrpc() *protopb.Product {
	return &protopb.Product{
		Id:       p.Id,
		Name:     p.Name,
		Price:    float64(p.Price),
		IsActive: p.IsActive,
		Quantity: int32(p.Quantity),
		Photo:    p.Photo,
		AgentId:  p.AgentId,
	}
}

func (o *Order) ConvertFromGrpc(order *protopb.Order) *Order {
	newOrder := &Order{
		Id:          order.Id,
		UserId:      order.UserId,
		ProductId:   order.ProductId,
		Quantity:    int(order.Quantity),
		DateCreated: order.DateCreated.AsTime(),
		Referral:    int(order.Referral),
		TotalPrice:  order.TotalPrice,
	}

	return newOrder
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
	}
}
