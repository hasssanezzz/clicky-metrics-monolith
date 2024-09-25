package domain

import (
	"context"
	"time"
)

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}

type LoginUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}

type URLUsecase interface {
	Create(c context.Context, url *URL) error
	FetchByUsername(c context.Context, username string) (URL, error)
	Resolve(c context.Context, short string) (URL, error)
	Deactivate(c context.Context, short string) error
	Update(c context.Context, url *URL) error
	Delete(c context.Context, short string) error
}

type URLAnalyticsUseCase interface {
	RecordAnalytics(c context.Context, urlID int, ipAddress, userAgent, browser, device, location string, accessedAt time.Time) error
	GetAnalyticsByURL(c context.Context, urlID int) ([]URLAnalytics, error)
	GetAnalyticsByDateRange(c context.Context, urlID int, startDate, endDate time.Time) ([]URLAnalytics, error)
}
