package content

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content string `json:"content"`
}

func (Message) TableName() string {
	return "message"
}
