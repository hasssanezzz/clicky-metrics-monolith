package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
)

type RefreshTokenController struct {
	AuthenticationUsecase domain.AuthenticationUsecase
	Env                   *bootstrap.Env
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (con *RefreshTokenController) Execute(c *gin.Context) {
	req := new(refreshTokenRequest)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := con.AuthenticationUsecase.ExtractIDFromToken(req.RefreshToken, con.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := con.AuthenticationUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := con.AuthenticationUsecase.CreateAccessToken(user, con.Env.AccessTokenSecret, con.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := con.AuthenticationUsecase.CreateRefreshToken(user, con.Env.RefreshTokenSecret, con.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
