package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your_project/websocket"
	"github.com/your_project/routes"
	"github.com/your_project/db"
)

func main() {
	// Initialize database
	db.ConnectDB()

	// Create a new router
	router := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(router)

	// Start WebSocket hub
	go websocket.StartHub()

	// Start the server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
