package domain

import "time"

type Article struct {
	ID    int
	Title string
	Text  string
	Tags  []Tag

	PublicationDate time.Time
	AuthorUsername  string
}
