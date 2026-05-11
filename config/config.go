package config

import "gorm.io/gorm"

var (
	DB        *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort       string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
	JWTExpiredMin string
	JWTRefresh    string
	JWTExpire     string
}

func 