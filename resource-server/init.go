package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const PROFILE_ENV_NAME = "RS_PROFILE"

func InitializeEnvironment() {
	var envFileName string
	if os.Getenv(PROFILE_ENV_NAME) != "" {
		envFileName = fmt.Sprintf(".env.%s", os.Getenv(PROFILE_ENV_NAME))
	} else {
		envFileName = ".env"
	}
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatal("Environment Variables required")
	}
}
