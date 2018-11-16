package models

type User struct {
	ID        int   `gorm:"primary_key" json:"id"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	DeletedAt string `json:"deletedAt"`
}
