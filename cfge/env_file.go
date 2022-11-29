package cfge

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadDefaultEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println(fmt.Errorf("unable to load .env file. Error: %s", err))
	}
}

func LoadEnvFile(envFilePath string) {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Println(fmt.Errorf("unable to load .env file. Path: %s Error: %s", envFilePath, err))
	}
}
