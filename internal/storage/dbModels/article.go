package dbModels

import (
	"article-service/internal/domain"
	"time"
)

type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Text  string `db:"text"`

	PublicationDate time.Time `db:"publication_date"`
	AuthorUsername  string    `db:"author_username"`
}

func ArticleDBToArticle(article Article, tags []Tag) domain.Article {
	return domain.Article{
		ID:    article.ID,
		Title: article.Title,
		Text:  article.Text,
		Tags:  TagsDBToTags(tags),

		PublicationDate: article.PublicationDate,
		AuthorUsername:  article.AuthorUsername,
	}
}

func ArticleToArticleDB(article domain.Article) (Article, []ArticleTagPair) {
	articleTagPairs := make([]ArticleTagPair, 0)

	for _, tag := range article.Tags {
		articleTagPairs = append(articleTagPairs, ArticleTagPair{
			ArticleID: article.ID,
			TagID:     tag.ID,
		})
	}

	resArticle := Article{
		ID:    article.ID,
		Title: article.Title,
		Text:  article.Text,

		PublicationDate: article.PublicationDate,
		AuthorUsername:  article.AuthorUsername,
	}

	return resArticle, articleTagPairs
}

func ArticlesDBToArticles(articles []Article, tags []Tag) []domain.Article {
	res := make([]domain.Article, 0)
	for _, article := range articles {
		res = append(res, ArticleDBToArticle(article, tags))
	}
	return res
}

func ArticlesToArticlesDB(articles []domain.Article) ([]Article, []ArticleTagPair) {
	panic("implement me")
}
