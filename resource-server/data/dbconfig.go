package data

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
