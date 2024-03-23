package dbModels

import "article-service/internal/domain"

type Tag struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

func TagsDBToTags(tags []Tag) []domain.Tag {
	res := make([]domain.Tag, 0)
	for _, tag := range tags {
		res = append(res, domain.Tag(tag))
	}
	return res
}

func TagsToTagsDB(tags []domain.Tag) []Tag {
	res := make([]Tag, 0)
	for _, tag := range tags {
		res = append(res, Tag(tag))
	}
	return res
}
