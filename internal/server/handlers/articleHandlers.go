package handlers

import (
	_ "article-service/cmd/docs"
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

// GetArticleByID
// @Summary Получение статьи по айди
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "айди нужной статьи"
// @Success 200 {object} dto.Success
// @Failure 500 {object} dto.ApiError
// @Router /api/article/ [get]
func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
		return
	}

	article, err := h.articleService.GetArticleByID(c.Request.Context(), id)
	if err != nil {
		apiErr := dto.NewApiError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, dto.ArticleToArticleDTO(article))
}

// GetAllArticles
// @Summary Получение всех статей
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {object} dto.Success
// @Failure 500 {object} dto.ApiError
// @Router /api/article [get]
func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.articleService.GetAllArticles(c.Request.Context())
	if err != nil {
		apiErr := dto.NewApiError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, dto.ArticlesToArticlesDTO(articles))
}

// CreateArticle
// @Summary Добавление статьи
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {object} dto.Success
// @Failure 500 {object} dto.ApiError
// @Router /api/article/create [post]
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var newArticleDTO dto.ArticleDTO
	if err := c.ShouldBindJSON(&newArticleDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Некорректный article")
		return
	}

	err := h.articleService.CreateArticle(c.Request.Context(), dto.ArticleDTOtoArticle(newArticleDTO))
	if err != nil {
		apiErr := dto.NewApiError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccess(true))
}

// DeleteArticle
// @Summary Удаление статьи по айди
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "айди нужной статьи"
// @Success 200 {object} dto.Success
// @Failure 500 {object} dto.ApiError
// @Router /api/article/delete [delete]
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
		return
	}

	err = h.articleService.DeleteArticle(c.Request.Context(), id)
	if err != nil {
		apiErr := dto.NewApiError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccess(true))
}

// UpdateArticle
// @Summary Обновление статьи по нужному айди
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "айди нужной статьи"
// @Success 200 {object} dto.Success
// @Failure 500 {object} dto.ApiError
// @Router /api/article/delete [update]
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Некорректный id")
	}

	var newArticleDTO dto.ArticleDTO
	if err := c.ShouldBindJSON(&newArticleDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Некорректный article")
		return
	}

	err = h.articleService.UpdateArticle(c.Request.Context(), id, dto.ArticleDTOtoArticle(newArticleDTO))
	if err != nil {
		apiErr := dto.NewApiError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccess(true))
}
