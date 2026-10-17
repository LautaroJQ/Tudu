package main

import (
	"ToDo/middleware"
	"ToDo/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Something went wrong...")
	}

	file, err := os.OpenFile("request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	mux := http.NewServeMux()

	routes.RegisterTaskRoutes(mux)

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	server := &http.Server{
		Addr:    port,
		Handler: middleware.Logger(logger, mux),
	}

	fmt.Printf("Listening on port %v", os.Getenv("PORT"))
	server.ListenAndServe()
}
