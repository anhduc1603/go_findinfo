package bean

import "time"

type FacebookToken struct {
	ID                   int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID               int       `json:"userid" gorm:"column:userid;"`
	FacebookAccessToken  string    `gorm:"type:text;not null" json:"facebook_access_token"`
	FacebookRefreshToken string    `gorm:"type:text" json:"facebook_refresh_token"`
	TokenExpiry          time.Time `json:"token_expiry"`
	UserCookie           string    `gorm:"type:varchar(2000)" json:"user_cookie"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (FacebookToken) TableName() string {
	return "facebook_tokens"
}
