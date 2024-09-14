package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content   string `json:"content"`
	Username  string `json:"username"`
	RoomID    string `json:"room_id"`
}
