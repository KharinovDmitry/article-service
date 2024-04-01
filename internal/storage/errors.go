package storage

import "github.com/pkg/errors"

var (
	ErrTagNotFound     = errors.New("tag not found")
	ErrArticleNotFound = errors.New("article not found")
)
