package handlers

import (
	"github.com/hyhkjiy/devin_test/backend/models"
	"github.com/kataras/iris/v12"
)

func RegisterUser(ctx iris.Context) {
	var req models.RegisterRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request format"})
		return
	}

	if req.Username == "" || req.Password == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Username and password are required"})
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.Create(); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to create user"})
		return
	}

	ctx.JSON(models.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
	})
}
