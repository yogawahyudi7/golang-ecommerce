package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   struct{ Port string }
	Database struct{ DSN string }
	RabbitMQ struct{ URL string }
	JWT      struct {
		Secret string
		TTL    time.Duration
	}
	Redis struct{ Addr string }
}

var C Config

func Init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
	C.Server.Port = viper.GetString("PORT")
	C.Database.DSN = viper.GetString("DB_DSN")
	C.RabbitMQ.URL = viper.GetString("RABBITMQ_URL")
	C.JWT.Secret = viper.GetString("JWT_SECRET")
	C.JWT.TTL = viper.GetDuration("JWT_TTL")
	C.Redis.Addr = viper.GetString("REDIS_ADDR")
}
