package app

import (
	"article-service/internal/config"
	"article-service/internal/server/handlers"
	"article-service/internal/server/router"
	"article-service/internal/services"
	"article-service/internal/storage/repositories"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func MustRun(cfg config.Config) {
	conn := ConnPostgress(cfg.ConnStr)

	articleRepo := repositories.NewArticleRepository(conn)

	articleService := services.NewArticleService(articleRepo)

	articleHandler := handlers.NewArticleHandler(articleService)

	router := router.NewRouter(cfg.Address, *articleHandler)

	router.Run()
}

func ConnPostgress(connStr string) *sqlx.DB {
	db := sqlx.MustOpen("postgres", connStr)
	return db
}
