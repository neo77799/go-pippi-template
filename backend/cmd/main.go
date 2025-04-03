package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shibu/omikuji-app/config"
	"github.com/shibu/omikuji-app/di"
)

func main() {
	config.InitLogger()

	db := config.InitDB()
	defer config.CloseDB(db)

	container := di.NewContainer(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := container.Router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
