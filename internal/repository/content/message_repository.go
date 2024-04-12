package repository

import (
	entity_content "rcp-api-data/internal/entity/domain/content"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	InsertMessage(content string) error
	UpdateMessage(id uint, content string) error
	DeleteMessage(id uint) error
	GetMessage(id uint) (entity_content.Message, error)
}

type MessageRepository struct {
	IMessageRepository
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (r *MessageRepository) InsertMessage(content string) error {
	message := entity_content.Message{
		Content: content,
	}
	result := r.DB.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MessageRepository) UpdateMessage(id uint, content string) error {
	message := entity_content.Message{
		Content: content,
	}
	result := r.DB.Model(&message).Where("id = ?", id).Updates(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(id uint) error {
	message := entity_content.Message{}
	result := r.DB.Where("id = ?", id).Delete(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MessageRepository) GetMessage(id uint) (entity_content.Message, error) {
	message := entity_content.Message{}
	result := r.DB.Where("id = ?", id).First(&message)
	if result.Error != nil {
		return entity_content.Message{}, result.Error
	}
	return message, nil
}
