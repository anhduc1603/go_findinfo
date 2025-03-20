package service

import (
	"LeakInfo/bean"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateUserLogs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataUserLog bean.UserLog

		if err := c.ShouldBind(&dataUserLog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userAgent := c.Request.UserAgent()
		// Tạo log mới
		userLog := bean.UserLog{
			UserID:    dataUserLog.UserID,
			IPPublic:  dataUserLog.IPPublic,
			UserAgent: userAgent,
			Action:    dataUserLog.Action,
		}

		if err := db.Create(&userLog).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dataUserLog.ID})
	}
}
