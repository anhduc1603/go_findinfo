package bean

import "time"

type GoogleToken struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID             int       `json:"userid" gorm:"column:userid;"`
	GoogleAccessToken  string    `gorm:"type:text;not null" json:"google_access_token"`
	GoogleRefreshToken string    `gorm:"type:text" json:"google_refresh_token"`
	TokenExpiry        time.Time `json:"token_expiry"`
	UserCookie         string    `gorm:"type:varchar(2000)" json:"user_cookie"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (GoogleToken) TableName() string {
	return "google_tokens"
}
