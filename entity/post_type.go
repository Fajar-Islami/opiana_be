package entity

//PostType struct represents post_types table in database
type PostType struct {
	// ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Jenis string `gorm:"type:varchar(255)" json:"jenis"`
	Type  string `gorm:"type:varchar(255)" json:"type"`
}
