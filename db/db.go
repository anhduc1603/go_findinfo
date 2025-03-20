package db

import (
	"LeakInfo/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB // Global DB instance (tuỳ bạn, hoặc có thể return mà không dùng global)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	// Build DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Optional: Ping database
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to Database successfully")
	return db, nil
}
