package entity

//Post struct represents Posts table in database
type Post struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	UserID      uint64    `gorm:"not null" json:"-"`
	User        *User     `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user,omitempty"`
	PostTypeID  uint64    `gorm:"not null" json:"-"`
	PostType    *PostType `gorm:"foreignkey:PostTypeID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"post_type,omitempty"`
}
