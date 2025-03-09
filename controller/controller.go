package controller

import (
	"LeakInfo/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
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
		v1.POST("/items", service.CreateItem(db))           // create item
		v1.GET("/items", service.GetListOfItems(db))        // list items
		v1.GET("/items/:id", service.ReadItemById(db))      // get an item by ID
		v1.PUT("/items/:id", service.EditItemById(db))      // edit an item by ID
		v1.DELETE("/items/:id", service.DeleteItemById(db)) // delete an item by ID
	}

	router.Run()
}

func Login() {

	http.HandleFunc("/", service.HomeHandler)
	http.HandleFunc("/auth/login", service.LoginHandler)
	http.HandleFunc("/auth/callback", service.CallbackHandler)

	port := "8080"
	fmt.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
