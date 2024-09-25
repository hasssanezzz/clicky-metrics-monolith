package usecase

import (
	"context"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/utils"
)

type LoginUsecase struct {
	UserRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{userRepository}
}

func (u *LoginUsecase) GetByUsername(c context.Context, username string) (*domain.User, error) {
	return u.UserRepository.GetByUsername(username)
}

func (u *LoginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (u *LoginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
