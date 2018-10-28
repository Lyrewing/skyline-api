package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}
