package response

type DisplayItemResponse struct {
	Info     string `json:"info" gorm:"column:info"`
	Content  string `json:"content" gorm:"column:content"`
	UserId   string `json:"userid" gorm:"column:userid"`
	UserName string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
}
