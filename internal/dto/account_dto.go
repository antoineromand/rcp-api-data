package dto

type AccountDTO struct {
	ActivityMessage *string `json:"activityMessage"`
	Address         *string `json:"address"`
	City            *string `json:"city"`
	Country         *string `json:"country"`
	PostalCode      *string `json:"postalCode"`
	PhoneNumber     *string `json:"phoneNumber"`
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Username        *string `json:"username"`
	Email           *string `json:"email"`
	Pasword         *string `json:"password"`
	IsNew           *bool   `json:"isNew"`
}
