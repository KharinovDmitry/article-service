package dbModels

import "article-service/internal/domain"

type Tag struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

func TagsDBToTags(tag []Tag) []domain.Tag {

}

func TagsToTagsDB(tag []domain.Tag) []Tag {

}
