package dto

// RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Fullname string `json:"fullname" form:"fullname" binding:"required,min=1"`
	Phone    string `json:"phone" form:"phone" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
