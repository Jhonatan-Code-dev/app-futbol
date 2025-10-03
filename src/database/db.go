package database

import (
	"app-futbol/src/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error al abrir la conexión: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Error obteniendo SQL DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conexión a la base de datos establecida")
	return db
}
