package service

import (
	"LeakInfo/bean"
	"LeakInfo/config"
	"LeakInfo/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"time"
)

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
func LoginHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cấu hình OAuth2 cho Google
		var googleOAuthConfig = config.NewGoogleOAuthConfig(cfg)

		url := googleOAuthConfig.AuthCodeURL(
			oauthStateString,
			oauth2.SetAuthURLParam("prompt", "select_account"))
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func LogoutHandler(c *gin.Context) {
	// Xóa cookie chứa JWT token
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	// Có thể thêm logic xóa session (nếu có) tại đây

	// Redirect về trang login (hoặc trả JSON nếu API)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Xử lý callback từ Google
func CallbackHandler(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("state") != oauthStateString {
			c.String(http.StatusBadRequest, "Invalid OAuth state")
			return
		}

		code := c.Query("code")
		var googleOAuthConfig = config.NewGoogleOAuthConfig(cfg)
		token, err := googleOAuthConfig.Exchange(context.Background(), code)
		if err != nil {
			c.String(http.StatusInternalServerError, "Token exchange failed")
			return
		}

		client := googleOAuthConfig.Client(context.Background(), token)
		resp, err := client.Get(cfg.GoogleApiOauth)
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

			// ✅ Lưu token: AccessToken, RefreshToken, Expiry
			accessToken := token.AccessToken
			refreshToken := token.RefreshToken
			expiry := token.Expiry
			err = SaveOrUpdateGoogleToken(user.ID, accessToken, "", refreshToken, expiry, db)
			if err != nil {
				log.Println("error:", err.Error())
			}
		} else {
			// Get existing user
			db.First(&user, provider.UserId)

			accessToken := token.AccessToken
			refreshToken := token.RefreshToken
			expiry := token.Expiry
			err = SaveOrUpdateGoogleToken(user.ID, accessToken, "", refreshToken, expiry, db)
			if err != nil {
				log.Println("error save or update google token:", err.Error())
			}
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

		redirectURL := fmt.Sprintf("%s?token=%s", cfg.FrontendRedirectURL, jwtToken)
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)

	}
}

func SaveOrUpdateGoogleToken(userID int, accessToken, refreshToken, userCookie string, expiry time.Time,
	database *gorm.DB) error {
	var token bean.GoogleToken

	// Kiểm tra xem đã tồn tại token cho user này chưa
	result := database.Where("userid = ?", userID).First(&token)

	if result.Error != nil && result.RowsAffected == 0 {
		// Chưa tồn tại → Insert mới
		newToken := bean.GoogleToken{
			UserID:             userID,
			GoogleAccessToken:  accessToken,
			GoogleRefreshToken: refreshToken,
			TokenExpiry:        expiry,
			UserCookie:         userCookie,
		}
		if err := database.Create(&newToken).Error; err != nil {
			return err
		}
	} else {
		// Đã tồn tại → Update
		token.GoogleAccessToken = accessToken
		token.GoogleRefreshToken = refreshToken
		token.TokenExpiry = expiry
		token.UserCookie = userCookie
		if err := database.Save(&token).Error; err != nil {
			return err
		}
	}

	return nil
}
