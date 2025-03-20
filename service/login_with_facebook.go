package service

import (
	"LeakInfo/bean"
	"LeakInfo/bean/request"
	"LeakInfo/config"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var jwtSecret = []byte("your_jwt_secret")

type FacebookUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Expiry  string `json:"expiry"`
}

func HandleFacebookLogin(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fbReq request.FacebookLoginRequest
		if err := c.ShouldBindJSON(&fbReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Verify Facebook Access Token
		fbUser, err := verifyFacebookToken(fbReq.AccessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Facebook token"})
			return
		}

		// Optional: Check or Create user in DB here (MySQL, Postgres, etc.)
		var provider bean.UserProvider
		err = db.Where("provider = ? AND provider_id = ?", "facebook", fbUser.ID).First(&provider).Error
		var user bean.User
		if err == gorm.ErrRecordNotFound {
			// Create new user
			user = bean.User{
				Username: fbUser.Email,
				Email:    fbUser.Email,
				Role:     "user",
				Status:   1,
			}
			db.Create(&user)
			// Create user_providers
			provider = bean.UserProvider{
				UserId:     user.ID,
				Provider:   "facebook",
				ProviderId: fbUser.ID,
				Avatar:     fbUser.Picture,
			}
			db.Create(&provider)

			// ✅ Lưu token: AccessToken, RefreshToken, Expiry
			accessToken := fbReq.AccessToken
			expiry := time.Now()
			err = SaveOrUpdateFacebookToken(user.ID, accessToken, "", expiry, db)
			if err != nil {
				log.Println("error:", err.Error())
			}
		} else {
			// Get existing user
			db.First(&user, provider.UserId)

			accessToken := fbReq.AccessToken
			expiry := time.Now()
			err = SaveOrUpdateFacebookToken(user.ID, accessToken, "", expiry, db)
			if err != nil {
				log.Println("error save or update google token:", err.Error())
			}
		}

		fmt.Printf("Facebook user: %+v\n", user)

		// Generate JWT Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate JWT"})
			return
		}

		//c.JSON(http.StatusOK, gin.H{"token": tokenString})
		redirectURL := fmt.Sprintf("%s?token=%s", cfg.FrontendRedirectURL, tokenString)
		log.Println("redirectURL:", redirectURL)
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}

// Verify token using Facebook Graph API
func verifyFacebookToken(accessToken string) (FacebookUser, error) {
	resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/me?fields=id,name,email&access_token=%s", accessToken))
	if err != nil {
		return FacebookUser{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var user FacebookUser
	if err := json.Unmarshal(body, &user); err != nil {
		return FacebookUser{}, err
	}

	return user, nil
}

func SaveOrUpdateFacebookToken(userID int, accessToken, userCookie string, expiry time.Time,
	database *gorm.DB) error {
	var token bean.FacebookToken

	// Kiểm tra xem đã tồn tại token cho user này chưa
	result := database.Where("userid = ?", userID).First(&token)

	if result.Error != nil && result.RowsAffected == 0 {
		// Chưa tồn tại → Insert mới
		newToken := bean.FacebookToken{
			UserID:              userID,
			FacebookAccessToken: accessToken,
			TokenExpiry:         expiry,
			UserCookie:          userCookie,
		}
		if err := database.Create(&newToken).Error; err != nil {
			return err
		}
	} else {
		// Đã tồn tại → Update
		token.FacebookRefreshToken = accessToken
		token.TokenExpiry = expiry
		token.UserCookie = userCookie
		if err := database.Save(&token).Error; err != nil {
			return err
		}
	}

	return nil
}
