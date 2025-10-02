package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"app-futbol/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // variable global accesible desde cualquier lado
var onceDB sync.Once

// InitDatabase inicializa la conexión solo una vez
func InitDatabase() *gorm.DB {
	onceDB.Do(func() {
		cfg := config.GetConfig() // obtenemos config global
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

		// Configuración del pool de conexiones
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		DB = db
		log.Println("✅ Conexión a la base de datos establecida")
	})

	return DB
}

// GetDB retorna la conexión global
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("❌ Base de datos no inicializada. Llama primero a InitDatabase()")
	}
	return DB
}
