package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hyhkjiy/devin_test/backend/models"
	"github.com/hyhkjiy/devin_test/backend/utils"
	"github.com/kataras/iris/v12"
	"strconv"
)

func CreateArticle(ctx iris.Context) {
	// Get token from Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid or missing authorization token"})
		return
	}

	token := authHeader[7:]
	// Validate token
	parsedToken, err := utils.ValidateToken(token)
	if err != nil || !parsedToken.Valid {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid token"})
		return
	}

	// Get user ID from token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid token claims"})
		return
	}

	userID := int64(claims["user_id"].(float64))

	var req models.CreateArticleRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request format"})
		return
	}

	if req.Title == "" || req.Content == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Title and content are required"})
		return
	}

	article := &models.Article{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := article.Create(); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to create article"})
		return
	}

	ctx.JSON(models.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Content:     article.Content,
		UserID:      article.UserID,
		PublishTime: article.PublishTime,
	})
}

func ListArticles(ctx iris.Context) {
	articles, err := models.GetArticles()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to fetch articles"})
		return
	}

	var response []models.ArticleResponse
	for _, article := range articles {
		response = append(response, models.ArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Content:     article.Content,
			UserID:      article.UserID,
			PublishTime: article.PublishTime,
		})
	}

	ctx.JSON(response)
}

func GetArticle(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid article ID"})
		return
	}

	article, err := models.GetArticleByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Article not found"})
		return
	}

	ctx.JSON(models.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Content:     article.Content,
		UserID:      article.UserID,
		PublishTime: article.PublishTime,
	})
}
