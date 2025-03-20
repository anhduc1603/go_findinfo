package bean

import "time"

type UserLog struct {
	ID                 int       `json:"id" gorm:"column:id;"`
	UserID             int       `json:"userid" gorm:"column:userid;"`
	IPPublic           string    `json:"ip_public" gorm:"column:ip_public;"`
	UserAgent          string    `json:"user_agent" gorm:"column:user_agent;"`
	Action             string    `json:"action" gorm:"column:action;"`
	StatusDownloadFile int       `json:"status_download_file" gorm:"column:status_download_file;"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (UserLog) TableName() string {
	return "user_logs"
}
