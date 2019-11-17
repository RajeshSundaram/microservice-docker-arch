package infra

import (
	"fmt"
	"os"
	"strconv"
)

func GetDbConnectionURL() string {
	PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	connectionURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), PORT, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	return connectionURL
}

// docker run --name wallet-postgres -e POSTGRES_PASSWORD=1234 -e POSTGRES_USER=prod -e POSTGRES_DB=wallet -p 5432:5432 -d postgres
