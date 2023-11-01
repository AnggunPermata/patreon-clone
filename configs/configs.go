package configs

import (
	"os"

	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretJWT string

type Config struct {
	ServerPort  string
	AppVersion  string
	Mode        string
	LoggerLevel string
	SecretJWT   string
	DB          *gorm.DB
	BaseURL string
}

func LoadConfig() (config *Config) {
	_ = godotenv.Load()

	serverPort := os.Getenv("PORT")
	mode := os.Getenv("MODE")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	appVersion := os.Getenv("APP_VERSION")
	secretJWT := os.Getenv("SECRET_JWT")
	baseUrl := os.Getenv("BASE_URL")
	db, _ := InitDB()
	config = &Config{
		ServerPort:  serverPort,
		AppVersion:  appVersion,
		Mode:        mode,
		LoggerLevel: loggerLevel,
		SecretJWT:   secretJWT,
		DB:          db,
		BaseURL: baseUrl,
	}

	SecretJWT = secretJWT

	return
}

func LoadEnv(key string) (value string) {
	return os.Getenv(key)
}

func InitDB() (*gorm.DB, error) {
	databaseUrl := LoadEnv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(DB)

	return DB, err
}

func InitMigrate(DB *gorm.DB) {
	if LoadEnv("AUTO_MIGRATION_STATUS") == "ON" {
		DB.AutoMigrate(&models.User{})
		DB.AutoMigrate(&models.Followers{})
		DB.AutoMigrate(&models.Post{})
	}
}
