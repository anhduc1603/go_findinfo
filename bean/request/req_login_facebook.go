package request

type FacebookLoginRequest struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userID"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}
