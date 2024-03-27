package postgres_database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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
		return nil, err
	}

	if err := envconfig.Process("", &dbConfig); err != nil {
		log.Fatalf("Error processing env variable to db configs %s", err)
		return nil, err
	}
	format := fmt.Sprintf("port=%s user=%s dbname=%s sslmode=%s password=%s", dbConfig.Port, dbConfig.UserName, dbConfig.DBname, dbConfig.Sslmode, dbConfig.Password)
	fmt.Println("check:", format)
	db, err := sql.Open(
		"postgres",
		format)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error opening db %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging db %s", err)
		return nil, err
	}
	query := "CREATE TABLE IF NOT EXISTS movies(" +
		"id SERIAL NOT NULL UNIQUE," +
		"title VARCHAR(255) NOT NULL," +
		"director VARCHAR(255) NOT NULL," +
		"year INT NOT NULL," +
		"country VARCHAR(255) NOT NULL)"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to db successfully")
	return db, nil
}
