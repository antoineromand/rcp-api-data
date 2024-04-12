package usecase

import (
	repository "rcp-api-data/internal/repository/content"

	"gorm.io/gorm"
)

type UpdateMessageUseCase struct {
	DB *gorm.DB
}

func NewUpdateMessageUseCase(db *gorm.DB) *UpdateMessageUseCase {
	return &UpdateMessageUseCase{
		DB: db,
	}
}

func (e *UpdateMessageUseCase) UpdateMessage(id uint, content string) error {
	err := repository.NewMessageRepository(e.DB).UpdateMessage(id, content)
	if err != nil {
		return err
	}
	return nil
}
