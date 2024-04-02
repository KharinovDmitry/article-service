package dto

type AddArticleRequest struct {
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Tags  []TagDTO `json:"tags"`

	AuthorUsername string `json:"authorUsername"`
}
