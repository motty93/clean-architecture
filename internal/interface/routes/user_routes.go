package routes

import (
	"net/http"

	"github.com/motty93/clean-architecture/internal/interface/handler"
)

func RegisterUserRoutes(mux *http.ServeMux, userHandler *handler.UserHandler) {
	// mux.HandleFunc("/users", userHandler.GetAllUsers)
	mux.HandleFunc("/user", userHandler.GetUserByID)
}
