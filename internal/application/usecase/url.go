package usecase

import "github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"

type URLUsecase struct {
	UrlRepository *domain.URLRepository
}

func NewURLUsecase(urlRepository *domain.URLRepository) *URLUsecase {
	return &URLUsecase{urlRepository}
}
