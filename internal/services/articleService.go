package services

import (
	"article-service/internal/domain"
	"context"
	"fmt"
)

type ArticleRepositoryContract interface {
	FindArticleByID(ctx context.Context, id int) (domain.Article, error)
	GetAllArticles(ctx context.Context) ([]domain.Article, error)
	AddArticle(ctx context.Context, article domain.Article) error
	DeleteArticle(ctx context.Context, id int) error
	UpdateArticle(ctx context.Context, id int, newArticle domain.Article) error
	AddTagToArticle(ctx context.Context, articleId int, tagId int) error
	RemoveTagFromArticle(ctx context.Context, articleId int, tagId int) error
}

type ArticleService struct {
	ArticleRepo ArticleRepositoryContract
}

func NewArticleService(articleRepo ArticleRepositoryContract) *ArticleService {
	return &ArticleService{ArticleRepo: articleRepo}
}

func (s *ArticleService) GetArticleByID(ctx context.Context, id int) (domain.Article, error) {
	article, err := s.ArticleRepo.FindArticleByID(ctx, id)
	if err != nil {
		return domain.Article{}, fmt.Errorf("In ArticleService(GetArticleByID): %w", err)
	}
	return article, nil
}

func (s *ArticleService) GetAllArticles(ctx context.Context) ([]domain.Article, error) {
	articles, err := s.ArticleRepo.GetAllArticles(ctx)
	if err != nil {
		return nil, fmt.Errorf("In ArticleService(GetAllArticles): %w", err)
	}
	return articles, nil
}

func (s *ArticleService) CreateArticle(ctx context.Context, article domain.Article) error {
	err := s.ArticleRepo.AddArticle(ctx, article)
	if err != nil {
		return fmt.Errorf("In ArticleService(CreateArticle): %w", err)
	}
	if err != nil {
		return fmt.Errorf("In ArticleService(CreateArticle): %w", err)
	}
	return nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id int) error {
	err := s.ArticleRepo.DeleteArticle(ctx, id)
	if err != nil {
		return fmt.Errorf("In ArticleService(DeleteArticle): %w", err)
	}
	return nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id int, newArticle domain.Article) error {
	err := s.ArticleRepo.UpdateArticle(ctx, id, newArticle)
	if err != nil {
		return fmt.Errorf("In ArticleService(UpdateArticle): %w", err)
	}
	return nil
}
