package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (con *LoginController) Execute(c *gin.Context) {
	req := new(loginRequest)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := con.LoginUsecase.GetByUsername(context.Background(), req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "username not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "invalid username/password combination"})
		return
	}

	accessToken, err := con.LoginUsecase.CreateAccessToken(user, con.Env.AccessTokenSecret, con.Env.AccessTokenExpiryHour)
	if err != nil {
		log.Printf("count not create access token: %v", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "internal server error"})
		return
	}

	refreshToken, err := con.LoginUsecase.CreateAccessToken(user, con.Env.RefreshTokenSecret, con.Env.RefreshTokenExpiryHour)
	if err != nil {
		log.Printf("count not refresh access token: %v", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
