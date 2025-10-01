package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	JWTSecret  string `envconfig:"JWT_SECRET" required:"true"`
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Printf("❌ Error al cargar la configuración: %v", err)
		return nil, err
	}

	log.Printf("✅ Variables de entorno cargadas correctamente")
	return cfg, nil
}
