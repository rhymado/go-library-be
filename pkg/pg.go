package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func InitDB() (*sqlx.DB, error) {
	var options []any
	options = append(options, os.Getenv("DB_HOST"))
	options = append(options, os.Getenv("DB_USER"))
	options = append(options, os.Getenv("DB_PWD"))
	options = append(options, os.Getenv("DB_NAME"))

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", options...)
	return sqlx.Connect("postgres", config)
}
