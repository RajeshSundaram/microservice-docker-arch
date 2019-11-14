package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"./data"
	_ "github.com/lib/pq"

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

func InitializeDatabase() {
	db, err := sql.Open("postgres", data.GetDbConnectionURL())
	if err != nil {
		log.Fatal("DATABASE Connection failed")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Print("DB connected successfully")
}
