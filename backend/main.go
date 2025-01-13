package main

import (
	"github.com/hyhkjiy/devin_test/backend/database"
	"github.com/hyhkjiy/devin_test/backend/handlers"
	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
)

func main() {
	app := iris.New()

	// Initialize database
	database.InitDB()

	// CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	app.WrapRouter(iris.FromStd(corsMiddleware.Handler))

	// Routes
	api := app.Party("/api")
	{
		users := api.Party("/users")
		{
			users.Post("/register", handlers.RegisterUser)
		}
	}

	app.Listen(":8000")
}
