package service

import (
	"LeakInfo/bean"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

// Đăng ký người dùng
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user bean.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
			return
		}

		// Mã hóa mật khẩu
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi mã hóa mật khẩu"})
			return
		}
		user.Password = string(hashedPassword)

		// Lưu vào DB
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo tài khoản"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
	}
}

// Đăng nhập và tạo token
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input bean.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
			return
		}

		var user bean.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Tài khoản không tồn tại"})
			return
		}

		// Kiểm tra mật khẩu
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai mật khẩu"})
			return
		}

		// Tạo token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"role":     user.Role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
