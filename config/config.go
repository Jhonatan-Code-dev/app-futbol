package config

import (
	"log"
	"sync"

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

var (
	cfg     *Config
	onceCfg sync.Once
)

// GetConfig devuelve la configuración global (cargada solo una vez)
func GetConfig() *Config {
	onceCfg.Do(func() {
		_ = godotenv.Load() // cargar .env si existe

		configInstance := &Config{}
		if err := envconfig.Process("", configInstance); err != nil {
			log.Fatalf("❌ Error al cargar la configuración: %v", err)
		}

		log.Println("✅ Variables de entorno cargadas correctamente")
		cfg = configInstance
	})

	return cfg
}
