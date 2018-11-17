package models

import (
	"database/sql"
)

type User struct {
	ID        int   `gorm:"primary_key" json:"id"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	DeletedAt sql.NullString `json:"deletedAt"`
}
