package usecase

import (
	"context"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/utils"
)

type AuthenticationUsecase struct {
	UserRepository domain.UserRepository
}

func NewAuthenticationUsecase(userRepository domain.UserRepository) *AuthenticationUsecase {
	return &AuthenticationUsecase{userRepository}
}

func (u *AuthenticationUsecase) Create(c context.Context, user *domain.User) error {
	return u.UserRepository.Create(user)
}

func (u *AuthenticationUsecase) GetUserByID(c context.Context, ID int) (*domain.User, error) {
	return u.UserRepository.GetByID(ID)
}

func (u *AuthenticationUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	return u.UserRepository.GetByUsername(username)
}

func (u *AuthenticationUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	return u.UserRepository.GetByEmail(email)
}

func (u *AuthenticationUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (u *AuthenticationUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
func (u *AuthenticationUsecase) ExtractIDFromToken(token string, secret string) (int, error) {
	return utils.ExtractIDFromToken(token, secret)
}
