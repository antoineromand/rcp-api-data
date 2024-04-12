package usecase

import (
	repository "rcp-api-data/internal/repository/content"

	"gorm.io/gorm"
)

type InsertMessageUseCase struct {
	DB *gorm.DB
}

func NewInsertMessageUseCase(db *gorm.DB) *InsertMessageUseCase {
	return &InsertMessageUseCase{
		DB: db,
	}
}

func (e *InsertMessageUseCase) InsertMessage(content string) error {
	err := repository.NewMessageRepository(e.DB).InsertMessage(content)
	if err != nil {
		return err
	}
	return nil
}
