package bean

import (
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Username  string     `json:"username" gorm:"column:username;"`
	Password  string     `json:"-" gorm:"column:password;"`
	Role      string     `json:"role" gorm:"column:role;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (User) TableName() string {
	return "users"
}
