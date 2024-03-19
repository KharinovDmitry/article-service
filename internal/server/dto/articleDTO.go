package dto

import (
	"article-service/internal/domain"
	"time"
)

type ArticleDTO struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Tags  []TagDTO `json:"tags"`

	PublicationDate time.Time `json:"publicationDate"`
	AuthorUsername  string    `json:"authorUsername"`
}

func ArticleDTOtoArticle(dto ArticleDTO) domain.Article {
	return domain.Article{
		ID:              dto.ID,
		Title:           dto.Title,
		Text:            dto.Text,
		Tags:            TagsDTOtoTags(dto.Tags),
		PublicationDate: dto.PublicationDate,
		AuthorUsername:  dto.AuthorUsername,
	}
}

func ArticleToArticleDTO(article domain.Article) ArticleDTO {
	return ArticleDTO{
		ID:              article.ID,
		Title:           article.Title,
		Text:            article.Text,
		Tags:            TagsToTagsDTO(article.Tags),
		PublicationDate: article.PublicationDate,
		AuthorUsername:  article.AuthorUsername,
	}
}

func ArticlesDTOtoArticles(dto []ArticleDTO) []domain.Article {
	res := make([]domain.Article, 0)
	for _, articleDTO := range dto {
		res = append(res, ArticleDTOtoArticle(articleDTO))
	}
	return res
}

func ArticlesToArticlesDTO(dto []domain.Article) []ArticleDTO {
	res := make([]ArticleDTO, 0)
	for _, article := range dto {
		res = append(res, ArticleToArticleDTO(article))
	}
	return res
}
