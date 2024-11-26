package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/motty93/clean-architecture/internal/domain/service"
	"github.com/motty93/clean-architecture/internal/infrastructure"
	"github.com/motty93/clean-architecture/internal/interface/handler"
	"github.com/motty93/clean-architecture/internal/usecase"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	repo := infrastructure.NewSupabaseRepository(conn)
	us := service.NewUserService()
	uu := usecase.NewUseUsecase(repo, us)
	uh := handler.NewUserHandler(uu)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/user", uh.GetUserByID)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
