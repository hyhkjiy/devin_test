package main

import (
	"github.com/hyhkjiy/devin_test/backend/database"
	"github.com/hyhkjiy/devin_test/backend/handlers"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Initialize database
	database.InitDB()

	// CORS configuration
	app.AllowMethods(iris.MethodOptions)
	app.Use(func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if ctx.Method() == iris.MethodOptions {
			ctx.StatusCode(204)
			return
		}
		ctx.Next()
	})

	// Routes
	api := app.Party("/api")
	{
		users := api.Party("/users")
		{
			users.Post("/register", handlers.RegisterUser)
			users.Post("/login", handlers.LoginUser)
		}

		articles := api.Party("/articles")
		{
			articles.Post("/", handlers.CreateArticle)
			articles.Get("/", handlers.ListArticles)
			articles.Get("/{id:int64}", handlers.GetArticle)
		}
	}

	app.Listen(":8000")
}
