package main

import (
	"ablufus/config"
	"ablufus/internal/database"
	"ablufus/internal/http/router"
	"log"
	"os"
)

func main() {
	var err error
	envs, err := config.LoadEnvVars()

	if err != nil {
		log.Fatalln("Failed loading env", err)
	}

	db, err := database.New(envs.DSN)
	if err != nil {
		log.Fatalln("can't connect with err", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatal("Error while trying to migrate", err)
		os.Exit(1)
	}

	e := router.Handlers(db)

	if err := e.Start(envs.ApiPort); err != nil {
		log.Fatal("Error running api.", err)
		os.Exit(1)
	}
}
