package repositories

import (
	"article-service/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
)

type ArticleRepository struct {
	db *sqlx.Conn
}

func NewArticleRepository(conn *sqlx.Conn) *ArticleRepository {
	return &ArticleRepository{
		db: conn,
	}
}

func (a *ArticleRepository) FindArticleByID(ctx context.Context, id int) (domain.Article, error) {
	panic("IMPLEMENT ME")
}

func (a *ArticleRepository) GetAllArticles(ctx context.Context) ([]domain.Article, error) {
	panic("IMPLEMENT ME")
}

func (a *ArticleRepository) AddArticle(ctx context.Context, article domain.Article) error {
	panic("IMPLEMENT ME")
}

func (a *ArticleRepository) DeleteArticle(ctx context.Context, id int) error {
	panic("IMPLEMENT ME")
}

func (a *ArticleRepository) UpdateArticle(ctx context.Context, id int, newArticle domain.Article) error {
	panic("IMPLEMENT ME")
}
