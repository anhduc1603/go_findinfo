package main

import (
	"LeakInfo/config"
	"LeakInfo/controller"
	"LeakInfo/db"
	"log"
)

func main() {

	cfg := config.LoadConfig()
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal("❌ Failed to connect DB:", err)
	}

	log.Println("Connected:", db)

	controller.Controller(db, cfg)

}
