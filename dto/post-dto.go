package dto

// PostUpdateDTO is a model that client use when updating a post
type PostUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" `
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	PostTypeID  uint64 `json:"post_type_id,omitempty" form:"post_type_id,omitempty"`
}

// PostCreateDTO is a model that client use when create a new post
type PostCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	PostTypeID  uint64 `json:"post_type_id"  form:"post_type_id" binding:"required"`
}
