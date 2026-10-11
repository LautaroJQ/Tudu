package main

import (
	"ToDo/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Something went wrong...")
	}

	server := http.NewServeMux()

	routes.RegisterTaskRoutes(server)

	http.ListenAndServe(":8080", server)
}
