package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/motty93/clean-architecture/internal/domain/service"
	"github.com/motty93/clean-architecture/internal/infrastructure"
	"github.com/motty93/clean-architecture/internal/interface/handler"
	"github.com/motty93/clean-architecture/internal/repository"
	"github.com/motty93/clean-architecture/internal/usecase"
)

var (
	dbUrl string
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func init() {
	dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}

	log.Printf("Database URL: %s", dbUrl)
}

func main() {
	conn, err := infrastructure.NewDatabaseConnection(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	supabaseRepo := infrastructure.NewSupabaseRepository(conn)
	us := service.NewUserService()
	userRepo := repository.NewUserRepository(supabaseRepo)
	userUsecase := usecase.NewUseUsecase(userRepo, us)
	userHandler := handler.NewUserHandler(userUsecase)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/user", userHandler.GetUserByID)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
