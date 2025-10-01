// package database
package database

import (
	"fmt"
	"log"
	"time" // Necesario para configurar el pool de conexiones

	"app-futbol/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabaseConnection es el provider de fx.
// Recibe la configuración (*config.Config) y el ciclo de vida (fx.Lifecycle) inyectados por fx.
func NewDatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Error al abrir la conexión: %v", err)
		return nil, err
	}

	// 1. Obtener la conexión SQL subyacente para el pooling y el cierre.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configuración del pool de conexiones (Buenas prácticas para un servidor)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conexión a la base de datos establecida")

	return db, nil
}
