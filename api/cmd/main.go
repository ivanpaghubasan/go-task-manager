package main

import (
	"fmt"
	"go-task-manager-api/internal/app"
	"go-task-manager-api/internal/config"
	"go-task-manager-api/internal/router"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cfg.DBUrl)
	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	application, err := app.New(db, cfg)
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	r := router.SetupRouter(application)
	_ = r
}
