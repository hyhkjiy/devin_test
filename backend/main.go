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
		AllowOrigin("*").
		AllowCredentials().
		AllowHeader("Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization").
		AllowMethod("POST", "GET", "OPTIONS", "PUT", "DELETE")
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
