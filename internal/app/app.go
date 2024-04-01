package app

import (
	"article-service/cmd/migrator"
	"article-service/internal/config"
	"article-service/internal/server/handlers"
	"article-service/internal/server/router"
	"article-service/internal/services"
	"article-service/internal/storage/repositories"
	adapter "article-service/lib/adapter/db"
)

func MustRun(cfg config.Config) {
	migrator.MustRun(cfg.ConnStr, cfg.DriverName, cfg.MigrationsDir)

	db := adapter.NewPostgresAdapter(cfg.TimeoutDbContext, cfg.ConnStr)
	defer db.Close()

	articleRepo := repositories.NewArticleRepository(db)

	articleService := services.NewArticleService(articleRepo)

	articleHandler := handlers.NewArticleHandler(articleService)

	router := router.NewRouter(cfg.Address, *articleHandler)

	router.Run()
}
