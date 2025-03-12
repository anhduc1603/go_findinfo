package response

import (
	"fmt"
	"time"
)

type ResponseHistoryInfo struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Info      string     `json:"info" gorm:"column:info;"`
	Content   string     `json:"content" gorm:"column:content;"`
	Status    int        `json:"status" gorm:"column:status;"`
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ResponseHistoryInfo) TableName() string {
	return "response_history_info"
}

func (res ResponseHistoryInfo) ToString() string {
	return fmt.Sprintf("id: %s\nInfo: %s\nContent: %s\nStatuus %s\n", res.Id, res.Info, res.Content, res.Status)
}
