package handlers

import (
	"github.com/hyhkjiy/devin_test/backend/models"
	"github.com/hyhkjiy/devin_test/backend/utils"
	"github.com/kataras/iris/v12"
)

func LoginUser(ctx iris.Context) {
	var req models.LoginRequest
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

	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid credentials"})
		return
	}

	if !user.VerifyPassword(req.Password) {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(models.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}

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
