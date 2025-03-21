package service

import (
	"LeakInfo/bean/request"
	"LeakInfo/bean/response"
	"LeakInfo/constant"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem response.ResponseHistoryInfo

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dataItem.Status = 1 // set to default
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
		userid, err := strconv.Atoi(c.Param("userid"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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

		// 📌 Kiểm tra `userID` có tồn tại không
		if userid == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
			return
		}
		var statusList = []int{constant.StatusSuccess, constant.StatusProcess}

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Where("userid = ? and status in (?) ", userid, statusList).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id DESC").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
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
		var statusList = []int{constant.StatusSuccess, constant.StatusProcess}

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Where("status IN (?)", statusList).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id DESC").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
	}
}

func GetListOfItemsWithInfo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		info := c.Param("info")

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
		var statusList = []int{constant.StatusSuccess, constant.StatusProcess}

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Where("info LIKE ? AND status IN (?)", "%"+info+"%", statusList).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id DESC").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
	}
}

func GetListOfItemsByAdmin(db *gorm.DB) gin.HandlerFunc {
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
		var statusList = []int{constant.StatusSuccess, constant.StatusProcess, constant.StatusUserClickDisplay, constant.StatusUserClickDownload}

		if err := db.Table(response.ResponseHistoryInfo{}.TableName()).
			Where("status IN (?)", statusList).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id DESC").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
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
			log.Println("error", string(err.Error()))
			return
		}
		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func DeleteItemByListId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Đọc toàn bộ request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		log.Println("📥 Request Body:", string(body))

		// Reset lại body để Gin có thể đọc tiếp (do ReadAll() làm mất dữ liệu)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		var listId request.ReqUpdateAllId

		if err := c.ShouldBind(&listId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("📥List IDS", listId)

		if len(listId.IDs) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "list_ids cannot be empty"})
			return
		}

		result := db.Model(&response.ResponseHistoryInfo{}).
			Where("id IN (?)", listId.IDs).
			Update("status", constant.StatusClose)

		// Kiểm tra số dòng bị ảnh hưởng
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No records updated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": true})
	}
}

func DeleteItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Đọc toàn bộ request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}
		log.Println("📥 Request Body:", string(body))

		// Reset lại body để Gin có thể đọc tiếp (do ReadAll() làm mất dữ liệu)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		var listId request.ReqUpdateAllId

		if err := c.ShouldBind(&listId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("📥List IDS", listId)

		if len(listId.IDs) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "list_ids cannot be empty"})
			return
		}

		result := db.Model(&response.ResponseHistoryInfo{}).
			Where("id IN (?)", listId.IDs).
			Update("status", constant.StatusClose)

		// Kiểm tra số dòng bị ảnh hưởng
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No records updated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": true})
	}
}

func GetDisplayItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var result response.DisplayItemResponse

		if err := db.Table("response_history_info r ").
			Select("r.info, r.content,r.userid, u.username, u.email ").
			Joins("join users u on r.userid = u.id").
			Where("r.id = ?", id).
			Scan(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": result})
	}

}

func GetListOfItemsByAdminWithInfo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		info := c.Param("info")

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
			Where("info LIKE ? ", "%"+info+"%").
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id DESC").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
	}
}

func UploadFileContent(c *gin.Context) {
	// Lấy file từ form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lưu file vào thư mục local (ví dụ ./uploads/)
	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload file thành công!",
		"file":    file.Filename,
		"path":    filePath,
	})
}

func GetInfoDashboard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var totalRecords int64
		var completed int64
		var inProgress int64
		var userDisplayedData int64
		var userDownloaded int64

		db.Model(&response.ResponseHistoryInfo{}).Count(&totalRecords)
		db.Model(&response.ResponseHistoryInfo{}).Where("status = ?", 1).Count(&completed)
		db.Model(&response.ResponseHistoryInfo{}).Where("status = ?", 2).Count(&inProgress)
		db.Model(&response.ResponseHistoryInfo{}).Where("status = ?", 3).Count(&userDisplayedData)
		db.Model(&response.ResponseHistoryInfo{}).Where("status = ?", 4).Count(&userDownloaded)

		c.JSON(http.StatusOK, gin.H{
			"totalRecords":  totalRecords,
			"completed":     completed,
			"inProgress":    inProgress,
			"displayedData": userDisplayedData,
			"downloaded":    userDownloaded,
		})
	}
}

func GetInfoDashboardByUserId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		var totalRecords int64
		var statusCount = make(map[int]int64)

		// Tổng số bản ghi theo user_id
		db.Model(&response.ResponseHistoryInfo{}).
			Where("user_id = ?", userId).
			Count(&totalRecords)

		// Đếm theo từng status (1, 2, 3, 4)
		statusList := []int{1, 2, 3, 4}
		for _, status := range statusList {
			var count int64
			db.Model(&response.ResponseHistoryInfo{}).
				Where("userid = ? AND status = ?", userId, status).
				Count(&count)
			statusCount[status] = count
		}

		c.JSON(http.StatusOK, gin.H{
			"totalRecords":  totalRecords,
			"completed":     statusCount[1],
			"inProgress":    statusCount[2],
			"displayedData": statusCount[3],
			"downloaded":    statusCount[4],
		})
	}
}
