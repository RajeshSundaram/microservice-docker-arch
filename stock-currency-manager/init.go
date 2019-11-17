package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	data "scm/infra"

	"github.com/joho/godotenv"
)

const PROFILE_ENV_NAME = "RS_PROFILE"

func InitializeEnvironment() {
	profiles := make([]string, 0)
	if os.Getenv(PROFILE_ENV_NAME) != "" {
		envProfile := strings.Split(os.Getenv(PROFILE_ENV_NAME), ",")
		for _, profile := range envProfile {
			profiles = append(profiles, fmt.Sprintf("%s.env", profile))
		}
	} else {
		profiles = append(profiles, "core.env")
	}
	err := godotenv.Load(profiles...)
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
