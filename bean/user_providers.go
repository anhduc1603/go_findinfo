package bean

import (
	"time"
)

type UserProvider struct {
	ID         int        `json:"id" gorm:"column:id;"`
	UserId     int        `json:"userid" gorm:"column:userid;"`
	Provider   string     `json:"provider" gorm:"column:provider;"`
	ProviderId string     `json:"provider_id" gorm:"column:provider_id;"`
	Name       string     `json:"name" gorm:"column:name;"`
	Avatar     string     `json:"avatar" gorm:"column:avatar;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (UserProvider) TableName() string {
	return "user_providers"
}
