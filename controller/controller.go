package controller

import (
	"LeakInfo/config"
	"LeakInfo/middleware"
	"LeakInfo/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"time"
)

func Controller(db *gorm.DB, cfg *config.Config) {
	router := gin.Default()

	// Bật CORS cho tất cả các request
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.UrlFe}, // Cho phép React frontend gọi API
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request trong 12h
	}))

	v1 := router.Group("/v1")
	{
		v1.POST("/items", service.CreateItem(db))
		v1.GET("/items", service.GetListOfItems(db))                      // list items by admin
		v1.GET("/items/search/:info", service.GetListOfItemsWithInfo(db)) //Search with SDT/Mail/TKCK
		v1.GET("/list/:userid", service.ReadItemByUserId(db))             // list by userId
		v1.GET("/items/:id", service.ReadItemById(db))                    // get an item by ID
		v1.POST("/items/update/:id", service.EditItemById(db))            // edit an item by ID
		v1.POST("/update/all", service.DeleteItemByListId(db))            // delete an item by list ID
		v1.POST("/delete", service.DeleteItems(db))                       // delete an item by ID
		v1.GET("/item/display/:id", service.GetDisplayItems(db))          // Get display item by ID
		v1.GET("/dasboard/:userId", service.GetInfoDashboardByUserId(db)) //Get thông tin dasborad byUser
	}

	router.POST("/userlogs", service.CreateUserLogs(db))

	// Đăng ký và đăng nhập
	router.POST("/register", service.Register(db))
	router.POST("/login", service.Login(db))

	//Login with google
	router.GET("/", service.HomeHandler)
	router.GET("/auth/login", service.LoginHandler(cfg))
	router.GET("/auth/callback", service.CallbackHandler(db, cfg))
	router.GET("/auth/logout", service.LogoutHandler)

	//Login with facebook
	router.POST("/auth/facebook", service.HandleFacebookLogin(db, cfg))

	// Route yêu cầu quyền admin
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Chào mừng admin!"})
		})

		admin.GET("/items", service.GetListOfItemsByAdmin(db))                      //Get all items with all status
		admin.POST("/items/upload", service.UploadFileContent)                      // Get display item by ID
		admin.GET("/items/search/:info", service.GetListOfItemsByAdminWithInfo(db)) //Search with SDT/Mail/TKCK
		admin.GET("/dasboard", service.GetInfoDashboard(db))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

	//router.Run()
}
