package database_postgres

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
)

type InfoDB struct {
	Port       string `envconfig:"DB_PORT"`
	DriverName string `envconfig:"DB_DRIVERNAME"`
	UserName   string `envconfig:"DB_USERNAME"`
	DBname     string `envconfig:"DB_DNAME"`
	Sslmode    string `envconfig:"DB_SSLMODE"`
	Password   string `envconfig:"DB_PASSWORD"`
}

func InitDB() (*sql.DB, error) {
	var dbConfig InfoDB
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	if err := envconfig.Process("", &dbConfig); err != nil {
		return nil, err
	}

	format := fmt.Sprintf("port=%s user=%s dbname=%s sslmode=%s password=%s", dbConfig.Port, dbConfig.UserName, dbConfig.DBname, dbConfig.Sslmode, dbConfig.Password)

	db, err := sql.Open(
		dbConfig.DriverName,
		format)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
