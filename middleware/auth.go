package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Vui lòng cung cấp token"})
			c.Abort()
			return
		}

		// Lấy phần token sau "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		// Kiểm tra role
		userRole, ok := claims["role"].(string)
		if !ok || userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Không có quyền truy cập"})
			c.Abort()
			return
		}

		c.Next()
	}
}
