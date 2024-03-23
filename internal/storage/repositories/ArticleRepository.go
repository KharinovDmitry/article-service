package repositories

import (
	"article-service/internal/domain"
	"article-service/internal/storage/dbModels"
	"context"
	"github.com/jmoiron/sqlx"
)

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(conn *sqlx.DB) *ArticleRepository {
	return &ArticleRepository{
		db: conn,
	}
}

func (a *ArticleRepository) FindArticleByID(ctx context.Context, id int) (domain.Article, error) {
	sql := "SELECT (id, title, text, tags, publication_date, author_username) FROM articles WHERE id = $1"
	var article dbModels.Article
	err := a.db.Get(&article, sql, id)
	if err != nil {
		return domain.Article{}, err
	}

	sql = "SELECT (id, title) FROM tags WHERE id IN (SELECT tag_id FROM article_tag_pairs WHERE article_id = $1)"
	var articleTagPairs []dbModels.Tag
	err = a.db.Select(&articleTagPairs, sql, id)
	if err != nil {
		return domain.Article{}, err
	}

	return dbModels.ArticleDBToArticle(article, articleTagPairs), nil
}

func (a *ArticleRepository) GetAllArticles(ctx context.Context) ([]domain.Article, error) {
	sql := "SELECT (id, title, text, tags, publication_date, author_username) FROM articles"
	var articles []dbModels.Article
	err := a.db.Select(&articles, sql)
	if err != nil {
		return nil, err
	}

	sql = "SELECT (id, title) FROM tags"
	var tags []dbModels.Tag
	err = a.db.Select(&tags, sql)
	if err != nil {
		return nil, err
	}

	return dbModels.ArticlesDBToArticles(articles, tags), err
}

func (a *ArticleRepository) AddArticle(ctx context.Context, article domain.Article) error {
	sql := "INSERT INTO articles(text, title, author_username, publication_date)  VALUES ($1, $2, $3, $4)"
	_, err := a.db.Exec(sql, article.Text, article.Title, article.AuthorUsername, article.PublicationDate)
	if err != nil {
		return err
	}

	for _, tag := range article.Tags {
		err = a.AddTagToArticle(ctx, article.ID, tag.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *ArticleRepository) DeleteArticle(ctx context.Context, id int) error {
	sql := "DELETE FROM articles WHERE id = $1"
	_, err := a.db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepository) UpdateArticle(ctx context.Context, id int, newArticle domain.Article) error {
	sql := "UPDATE articles SET text=$1, title=$2, author_username=$3, publication_date=$4 WHERE id = $5"
	_, err := a.db.Exec(sql, newArticle.Text, newArticle.Title, newArticle.AuthorUsername, newArticle.PublicationDate, id)
	if err != nil {
		return err
	}

	sql = "DELETE FROM article_tag_pairs WHERE article_id = $1"
	_, err = a.db.Exec(sql, id)
	if err != nil {
		return err
	}

	for _, tag := range newArticle.Tags {
		err = a.AddTagToArticle(ctx, id, tag.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *ArticleRepository) AddTagToArticle(ctx context.Context, articleId int, tagId int) error {
	sql := "INSERT INTO article_tag_pairs(article_id, tag_id) VALUES ($1, $2)"
	_, err := a.db.Exec(sql, articleId, tagId)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepository) RemoveTagFromArticle(ctx context.Context, articleId int, tagId int) error {
	sql := "DELETE FROM article_tag_pairs WHERE article_id = $1 AND tag_id = $2"
	_, err := a.db.Exec(sql, articleId, tagId)
	if err != nil {
		return err
	}
	return nil
}
