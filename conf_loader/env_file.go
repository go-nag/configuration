package conf_loader

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadDefaultEnvFile() {
	if err := godotenv.Load(); err != nil {
		_ = fmt.Errorf("unable to load .env file. Error: %s", err)
	}
}

func LoadEnvFile(envFilePath string) {
	if err := godotenv.Load(envFilePath); err != nil {
		_ = fmt.Errorf("unable to load .env file. Path: %s Error: %s", envFilePath, err)
	}
}
