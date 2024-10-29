package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"url-shortener/routes"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("BACKEND")
	if port == "" {
		port = "8080"
	}

	server.GET("/url", routes.GetUrl)
	server.GET("/shorten", routes.ShortenUrl)

	err = server.Run(port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	log.Printf("Starting server on port %s...\n", port)
}
