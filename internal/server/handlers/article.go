package handlers

import (
	"article-service/internal/domain"
	"article-service/internal/server/dto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleServiceContract interface {
	GetArticleByID(ctx context.Context, id int) (domain.Article, error)
	GetAllArticles(ctx context.Context) ([]domain.Article, error)
	CreateArticle(ctx context.Context, article domain.Article) error
	DeleteArticle(ctx context.Context, id int) error
	UpdateArticle(ctx context.Context, id int, newArticle domain.Article) error
}

type ArticleHandler struct {
	articleService ArticleServiceContract
}

func NewArticleHandler(service ArticleServiceContract) *ArticleHandler {
	return &ArticleHandler{articleService: service}
}

func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
	}

	article, err := h.articleService.GetArticleByID(c.Request.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, dto.ArticleToArticleDTO(article))
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.articleService.GetAllArticles(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, dto.ArticlesToArticlesDTO(articles))
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var newArticleDTO dto.ArticleDTO
	if err := c.ShouldBindJSON(&newArticleDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Некорректный article")
	}

	err := h.articleService.CreateArticle(c.Request.Context(), dto.ArticleDTOtoArticle(newArticleDTO))
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
	}

	err = h.articleService.DeleteArticle(c.Request.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
	}

	var newArticleDTO dto.ArticleDTO
	if err := c.ShouldBindJSON(&newArticleDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Некорректный article")
	}

	err = h.articleService.UpdateArticle(c.Request.Context(), id, dto.ArticleDTOtoArticle(newArticleDTO))
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}
