package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

type signupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (con *SignupController) Execute(c *gin.Context) {
	req := new(signupRequest)

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err := con.SignupUsecase.GetUserByEmail(context.Background(), req.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "email is in use"})
	}

	_, err = con.SignupUsecase.GetUserByUsername(context.Background(), req.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "username is in use"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	req.Password = string(hashedPassword)

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = con.SignupUsecase.Create(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
}
