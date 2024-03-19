package dto

import "article-service/internal/domain"

type TagDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TagsDTOtoTags(dto []TagDTO) []domain.Tag {
	res := make([]domain.Tag, 0)
	for _, tagDTO := range dto {
		res = append(res, domain.Tag(tagDTO))
	}
	return res
}

func TagsToTagsDTO(dto []domain.Tag) []TagDTO {
	res := make([]TagDTO, 0)
	for _, tag := range dto {
		res = append(res, TagDTO(tag))
	}
	return res
}
