package main

import (
	"LeakInfo/controller"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	//controller.LoginWithGoogle()
	dsn := "root:@tcp(127.0.0.1:3306)/info_find?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Bật log SQL
	})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	//controller.Controller(db)
	controller.Login(db)

}
