package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mysql *mysqlConfig
	// WebService *webServiceConfig
	// Storage    *storageConfig
}

type mysqlConfig struct {
	Host          string
	User          string
	Password      string
	DatabaseName  string
	Port          string
	ContainerName string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Use local environment values.")
	} else {
		log.Print("Use .env file.")
	}
}

func NewConfig() *Config {
	return &Config{
		Mysql: &mysqlConfig{
			Host:         os.Getenv("MYSQL_HOST"),
			User:         os.Getenv("MYSQL_USER"),
			Password:     os.Getenv("MYSQL_PASSWORD"),
			DatabaseName: os.Getenv("MYSQL_DATABASE_NAME"),
			Port:         os.Getenv("MYSQL_PORT"),
		},
		// WebService: &webServiceConfig{
		// 	URL: os.Getenv("FRONT_URL"),
		// },
		// Storage: &storageConfig{
		// 	BucketName: os.Getenv("BUCKET_NAME"),
		// },
	}
}
