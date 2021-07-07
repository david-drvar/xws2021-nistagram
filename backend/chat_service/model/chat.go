package model

import "time"

type Message struct{
	Id string `gorm:"primaryKey"`
	SenderId string
	ReceiverId string
	RoomId string
	DateCreated time.Time
	ContentType ContentType
	IsRead bool
	Content string
	IsMediaOpened bool
}

type ContentType string
const(
	Image ContentType = "Image"
	Video             = "Video"
	String             = "String"
	Link               = "Link"
)

type MessageRequest struct{
	SenderId string `gorm:"primaryKey"`
	ReceiverId string `gorm:"primaryKey"`
	IsAccepted bool
}

type ChatRoom struct {
	Id string `gorm:"primaryKey"`
	Person1 string
	Person2 string
}