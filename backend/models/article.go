package models

import (
	"github.com/hyhkjiy/devin_test/backend/database"
	"time"
)

type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserID      int64     `json:"user_id"`
	PublishTime time.Time `json:"publish_time"`
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserID      int64     `json:"user_id"`
	PublishTime time.Time `json:"publish_time"`
}

func (a *Article) Create() error {
	query := `INSERT INTO articles (title, content, user_id) VALUES (?, ?, ?)`
	result, err := database.DB.Exec(query, a.Title, a.Content, a.UserID)
	if err != nil {
		return err
	}

	a.ID, _ = result.LastInsertId()
	return nil
}

func GetArticles() ([]Article, error) {
	query := `SELECT id, title, content, user_id, publish_time FROM articles ORDER BY publish_time DESC`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.UserID, &article.PublishTime)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func GetArticleByID(id int64) (*Article, error) {
	query := `SELECT id, title, content, user_id, publish_time FROM articles WHERE id = ?`
	var article Article
	err := database.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.UserID, &article.PublishTime)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
