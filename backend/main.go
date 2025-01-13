package main

import (
	"github.com/hyhkjiy/devin_test/backend/database"
	"github.com/hyhkjiy/devin_test/backend/handlers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
)

func main() {
	app := iris.New()

	// Initialize database
	database.InitDB()

	// CORS middleware
	crs := cors.New().
		AllowedOrigins([]string{"*"}).
		AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}).
		AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	app.UseRouter(crs)

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
