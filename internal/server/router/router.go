package router

import (
	_ "article-service/docs"
	"article-service/internal/server/handlers"
	"article-service/internal/server/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Router struct {
	address        string
	articleHandler handlers.ArticleHandler
}

func NewRouter(address string, articleHandler handlers.ArticleHandler) *Router {
	return &Router{
		address:        address,
		articleHandler: articleHandler,
	}
}

func (r *Router) Run() {
	g := gin.New()
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := g.Group("/api")

	api.GET("/article", r.articleHandler.GetAllArticles)
	article := api.Group("/article")
	{
		article.GET("/", r.articleHandler.GetArticleByID)
		article.POST("/create", r.articleHandler.CreateArticle).Use(middlewares.AuthMiddleware)
		article.DELETE("/delete", r.articleHandler.DeleteArticle).Use(middlewares.AuthMiddleware)
		article.PUT("/update", r.articleHandler.UpdateArticle).Use(middlewares.AuthMiddleware)
	}

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	g.Run(r.address)
}
