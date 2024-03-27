package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"to-do-movie_list/internal/database"
	"to-do-movie_list/internal/handler"
	"to-do-movie_list/internal/service"
	"to-do-movie_list/pkg/postgres_database"
	"to-do-movie_list/server"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	db, err := postgres_database.InitDB()

	if err != nil {
		log.Fatal(err)
	}
	databases := database.NewDB(db)
	services := service.NewService(databases)
	handlers := handler.NewHandler(services)

	log.Println("SERVER IS STARTED")

	srv := new(server.Server)

	if err := srv.RunServer(":8080", handlers.InitHandler()); err != nil {
		log.Fatal(err)
	}
}
