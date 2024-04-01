package dto

import (
	"article-service/internal/storage"
	"github.com/pkg/errors"
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `yaml:"message"`
}

func NewApiError(err error) ApiError {
	switch {
	case errors.Is(err, storage.ErrArticleNotFound),
		errors.Is(err, storage.ErrTagNotFound):
		return ApiError{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		}
	default:
		return ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    "",
		}
	}
}
