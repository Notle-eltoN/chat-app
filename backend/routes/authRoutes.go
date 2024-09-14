package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your_project/controllers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/ws", controllers.HandleWebSocket).Methods("GET")
}
