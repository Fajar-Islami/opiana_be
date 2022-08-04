package entity

//User represents users table in database
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string  `gorm:"type:varchar(255)" json:"fullname"`
	Phone    string  `gorm:"type:varchar(255)" json:"phone"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"type:varchar(255)" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Posts    *[]Post `json:"post,omitempty"`
}
