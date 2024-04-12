package usecase

import (
	repository "rcp-api-data/internal/repository/content"

	"gorm.io/gorm"
)

type DeleteMessageUseCase struct {
	DB *gorm.DB
}

func NewDeleteMessageUseCase(db *gorm.DB) *GetMessageUseCase {
	return &GetMessageUseCase{
		DB: db,
	}
}

func (e *GetMessageUseCase) DeleteMessage(id uint) error {
	err := repository.NewMessageRepository(e.DB).DeleteMessage(id)
	if err != nil {
		return err
	}
	return nil
}
