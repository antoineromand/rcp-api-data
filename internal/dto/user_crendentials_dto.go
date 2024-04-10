package dto

type UserCredentialsDTO struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

func NewUserCredentialsDTO(username, password, email *string) *UserCredentialsDTO {
	return &UserCredentialsDTO{
		Username: username,
		Password: password,
		Email:    email,
	}
}
