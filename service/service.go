package service

import (
	"LeakInfo/bean/response"
	"bytes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem response.ResponseHistoryInfo

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dataItem.Status = 1 // set to default
		dataItem.UserId = generateShortID()
		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}

func ReadItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem response.ResponseHistoryInfo

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}

func ReadItemByUserId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []response.ResponseHistoryInfo

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		log.Println("📥 Request Body:", string(body))

		// Reset lại body để Gin có thể đọc tiếp (do ReadAll() làm mất dữ liệu)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// 📌 Đọc trực tiếp JSON từ request body
		var dataItem response.ResponseHistoryInfo
		if err := c.ShouldBindJSON(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}
		// 📌 Kiểm tra `userID` có tồn tại không
		if dataItem.UserId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
			return
		}

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Where("userid = ?", dataItem.UserId).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []response.ResponseHistoryInfo

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Đọc toàn bộ request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		log.Println("📥 Request Body:", string(body))

		// Reset lại body để Gin có thể đọc tiếp (do ReadAll() làm mất dữ liệu)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		var dataItem response.ResponseHistoryInfo

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//db.Updates(&dataItem)
		if err := db.Model(&dataItem).
			Where("id = ? AND userid = ?", id, dataItem.UserId).
			Update("status", dataItem.Status).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func EditItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Đọc toàn bộ request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		log.Println("📥 Request Body:", string(body))

		// Reset lại body để Gin có thể đọc tiếp (do ReadAll() làm mất dữ liệu)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		var dataItem response.ResponseHistoryInfo

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//db.Updates(&dataItem)

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

//func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
//			Where("id = ?", id).
//			Delete(nil).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{"data": true})
//	}
//}

// Gen userId
func generateShortID() int {
	now := time.Now()
	milli := now.Nanosecond() / 1e6    // Lấy mili giây (0-999)
	randomPart := rand.Intn(900) + 100 // Sinh số ngẫu nhiên từ 100-999

	// Ghép mili giây và số ngẫu nhiên (lấy 3 số cuối của mili giây)
	return (milli*1000 + randomPart) % 1000000 // Giới hạn 6 chữ số
}
