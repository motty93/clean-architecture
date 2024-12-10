package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/motty93/clean-architecture/internal/domain/service"
	"github.com/motty93/clean-architecture/internal/infrastructure"
	"github.com/motty93/clean-architecture/internal/interface/handler"
	"github.com/motty93/clean-architecture/internal/interface/routes"
	"github.com/motty93/clean-architecture/internal/repository"
	"github.com/motty93/clean-architecture/internal/usecase"
	"go.uber.org/dig"
)

var (
	dbUrl string
)

func buildContainer() *dig.Container {
	container := dig.New()

	// register service
	container.Provide(func() string { return os.Getenv("DATABASE_URL") })
	container.Provide(infrastructure.NewDatabaseConnection)
	container.Provide(infrastructure.NewSupabaseRepository)

	// user
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	container.Provide(usecase.NewUserUsecase)
	container.Provide(handler.NewUserHandler)

	return container
}

func main() {
	container := buildContainer()
	cleanupManager := infrastructure.NewCleanupManager()

	// init and run server
	err := container.Invoke(func(
		userHandler *handler.UserHandler,
		db *infrastructure.SupabaseRepository,
	) {
		// cleanup
		cleanupManager.Add(func() error {
			log.Println("Closing database connection...")
			return db.Conn.Close(context.Background())
		})

		mux := http.NewServeMux()

		// register routes
		routes.ResisterApplicationRoutes(mux)
		routes.RegisterUserRoutes(mux, userHandler)
	})

	if err != nil {
		log.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)
	}
}
