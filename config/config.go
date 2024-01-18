package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeConfig(files ...string) error {
	// Check Each Environment Variable is Exist
	if _, exist := os.LookupEnv("DB_HOST"); !exist {
		// Load .env file
		if err := godotenv.Load(files...); err != nil {
			return err
		}

		// Initialize Oauth
		InitializeOauth()
	}

	return nil
}
