package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UserRegisterInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required,oneof=user admin"`
}

type UserLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
