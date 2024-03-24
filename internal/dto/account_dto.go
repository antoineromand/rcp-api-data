package dto

type AccountDTO struct {
	Username        *string `json:"username"`
	ActivityMessage *string `json:"activityMessage"`
	Address         *string `json:"address"`
	City            *string `json:"city"`
	Country         *string `json:"country"`
	PostalCode      *string `json:"postalCode"`
	PhoneNumber     *string `json:"phoneNumber"`
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	IsNew           *bool   `json:"isNew"`
}