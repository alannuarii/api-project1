package db

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    dsn := "user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + name + " sslmode=" + sslmode

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    DB = db
}
