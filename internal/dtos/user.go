package dtos

type LoginInput struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}

type RegisterInput struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}
