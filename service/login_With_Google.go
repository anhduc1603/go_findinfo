package service

import (
	"LeakInfo/bean"
	"LeakInfo/constant"
	"LeakInfo/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

// Cấu hình OAuth2 cho Google
var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/callback",
	ClientID:     constant.ClientID,
	ClientSecret: constant.ClientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthStateString = "random-state"

func HomeHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html")
	tmpl.Execute(c.Writer, nil)
}

// Xử lý đăng nhập
func LoginHandler(c *gin.Context) {
	url := googleOAuthConfig.AuthCodeURL(
		oauthStateString,
		oauth2.SetAuthURLParam("prompt", "select_account"))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func LogoutHandler(c *gin.Context) {
	// Xóa cookie chứa JWT token
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	// Có thể thêm logic xóa session (nếu có) tại đây

	// Redirect về trang login (hoặc trả JSON nếu API)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Xử lý callback từ Google
func CallbackHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("state") != oauthStateString {
			c.String(http.StatusBadRequest, "Invalid OAuth state")
			return
		}

		code := c.Query("code")
		token, err := googleOAuthConfig.Exchange(context.Background(), code)
		if err != nil {
			c.String(http.StatusInternalServerError, "Token exchange failed")
			return
		}

		client := googleOAuthConfig.Client(context.Background(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user info")
			return
		}
		defer resp.Body.Close()

		var googleUser struct {
			Id      string `json:"id"`
			Email   string `json:"email"`
			Picture string `json:"picture"`
			Name    string `json:"name"`
		}
		json.NewDecoder(resp.Body).Decode(&googleUser)

		// Check user_providers table
		var provider bean.UserProvider
		err = db.Where("provider = ? AND provider_id = ?", "google", googleUser.Id).First(&provider).Error
		var user bean.User
		if err == gorm.ErrRecordNotFound {
			// Create new user
			user = bean.User{
				Username: googleUser.Email,
				Email:    googleUser.Email,
				Role:     "user",
				Status:   1,
			}
			db.Create(&user)

			// Create user_providers
			provider = bean.UserProvider{
				UserId:     user.ID,
				Provider:   "google",
				ProviderId: googleUser.Id,
				Avatar:     googleUser.Picture,
			}
			db.Create(&provider)
		} else {
			// Get existing user
			db.First(&user, provider.UserId)
		}

		// Generate JWT token
		jwtToken, err := utils.GenerateJWTToken(user.ID, user.Username, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		//c.JSON(http.StatusOK, gin.H{
		//	"message":      "Login successful",
		//	"token":        jwtToken,
		//	"access_token": token.AccessToken, // Google AccessToken
		//	"user":         user,
		//})

		redirectURL := fmt.Sprintf("http://localhost:3000/oauth-success?token=%s", jwtToken)
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)

	}
}
