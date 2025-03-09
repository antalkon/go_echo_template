package main

import (
	"backend/internal/app"
	"log"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application initialization failed: %v", err)
	}
	application.Run()
	application.RunServer()
}
