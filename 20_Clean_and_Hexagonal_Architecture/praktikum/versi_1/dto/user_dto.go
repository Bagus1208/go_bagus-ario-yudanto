package dto

type UserDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginModel struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
