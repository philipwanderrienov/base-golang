package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost", // Change as needed
		Port:     5432,
		User:     "postgres",  // Replace with your DB username
		Password: "qwerty123", // Replace with your DB password
		DBName:   "church",    // Replace with your DB name
		SSLMode:  "disable",   // or "require" depending on your setup
	}
}

func (c *DBConfig) Connect() (*sql.DB, error) {
	var connStr string
	// Omit the password from the connection string if it's not provided.
	if c.Password != "" {
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
	} else {
		connStr = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.DBName, c.SSLMode)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database")
	return db, nil
}
