package account

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	ActivityMessage  string    `json:"activityMessage"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	Country          string    `json:"country"`
	PostalCode       string    `json:"postalCode"`
	PhoneNumber      string    `json:"phoneNumber"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	IsNew            bool      `json:"isNew" gorm:"default:false"`
	UserUUID         uuid.UUID `json:"user_uuid" gorm:"type:uuid;uniqueIndex;not null"`
}

func NewAccount(username, email, activityMessage, address, city, country, postalCode, phoneNumber, firstName, lastName string, isNew bool, userUUID uuid.UUID) *Account {
	return &Account{
		Username:     username,
		Email:        email,
		ActivityMessage: activityMessage,
		Address:      address,
		City:         city,
		Country:      country,
		PostalCode:   postalCode,
		PhoneNumber:  phoneNumber,
		FirstName:    firstName,
		LastName:     lastName,
		IsNew:        isNew,
		UserUUID: userUUID,
	}
}