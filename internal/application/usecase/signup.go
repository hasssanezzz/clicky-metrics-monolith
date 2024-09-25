package usecase

import (
	"context"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/utils"
)

type SignupUsecase struct {
	UserRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) *SignupUsecase {
	return &SignupUsecase{userRepository}
}

func (u *SignupUsecase) Create(c context.Context, user *domain.User) error {
	return u.UserRepository.Create(user)
}

func (u *SignupUsecase) GetByUsername(c context.Context, username string) (*domain.User, error) {
	return u.UserRepository.GetByUsername(username)
}

func (u *SignupUsecase) GetByEmail(c context.Context, email string) (*domain.User, error) {
	return u.UserRepository.GetByEmail(email)
}

func (u *SignupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (u *SignupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
