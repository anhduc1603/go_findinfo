package controller

import (
	"LeakInfo/middleware"
	"LeakInfo/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func Controller(db *gorm.DB) {
	router := gin.Default()

	// Bật CORS cho tất cả các request
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Cho phép React frontend gọi API
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request trong 12h
	}))

	v1 := router.Group("/v1")
	{
		v1.POST("/items", service.CreateItem(db))
		v1.GET("/items", service.GetListOfItems(db))           // list items
		v1.GET("/items/userid", service.ReadItemByUserId(db))  // list by userId
		v1.GET("/items/:id", service.ReadItemById(db))         // get an item by ID
		v1.POST("/items/:id", service.EditItemById(db))        // edit an item by ID
		v1.POST("/update/all", service.DeleteItemByListId(db)) // delete an item by list ID
		v1.POST("/delete", service.DeleteItems(db))            // delete an item by ID
	}

	// Đăng ký và đăng nhập
	router.POST("/register", service.Register(db))
	router.POST("/login", service.Login(db))

	// Route yêu cầu quyền admin
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Chào mừng admin!"})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

	//router.Run()
}

func LoginWithGoogle() {

	http.HandleFunc("/", service.HomeHandler)
	http.HandleFunc("/auth/login", service.LoginHandler)
	http.HandleFunc("/auth/callback", service.CallbackHandler)

	port := "8080"
	fmt.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//
//func Login(db *gorm.DB) {
//	r := gin.Default()
//
//	r.Use(cors.New(cors.Config{
//		AllowOrigins:     []string{"http://localhost:3000"}, // Cho phép React frontend gọi API
//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
//		ExposeHeaders:    []string{"Content-Length"},
//		AllowCredentials: true,
//		MaxAge:           12 * time.Hour, // Cache preflight request trong 12h
//	}))
//
//	// Đăng ký và đăng nhập
//	r.POST("/register", service.Register(db))
//	r.POST("/login", service.Login(db))
//
//	// Route yêu cầu quyền admin
//	admin := r.Group("/admin")
//	admin.Use(middleware.AuthMiddleware("admin"))
//	{
//		admin.GET("/dashboard", func(c *gin.Context) {
//			c.JSON(200, gin.H{"message": "Chào mừng admin!"})
//		})
//	}
//
//	port := os.Getenv("PORT")
//	if port == "" {
//		port = "8080"
//	}
//
//	r.Run(":" + port)
//}
