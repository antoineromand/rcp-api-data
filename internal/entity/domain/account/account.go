package account

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ActivityMessage *string   `json:"activityMessage"`
	Address         *string   `json:"address"`
	City            *string   `json:"city"`
	Country         *string   `json:"country"`
	PostalCode      *string   `json:"postalCode"`
	PhoneNumber     *string   `json:"phoneNumber"`
	FirstName       *string   `json:"firstName"`
	LastName        *string   `json:"lastName"`
	IsNew           *bool     `json:"isNew" gorm:"default:true"`
	Username        *string   `json:"username" gorm:"not null"`
	Email           *string   `json:"email" gorm:"not null"`
	UserUUID        uuid.UUID `json:"-" gorm:"type:uuid;uniqueIndex;not null"`
}

func NewAccount(activityMessage, address, city, country, postalCode, phoneNumber, firstName, lastName *string, isNew *bool, userUUID uuid.UUID, username, email *string) *Account {
	return &Account{
		ActivityMessage: activityMessage,
		Address:         address,
		City:            city,
		Country:         country,
		PostalCode:      postalCode,
		PhoneNumber:     phoneNumber,
		FirstName:       firstName,
		LastName:        lastName,
		IsNew:           isNew,
		Username:        username,
		Email:           email,
		UserUUID:        userUUID,
	}
}
