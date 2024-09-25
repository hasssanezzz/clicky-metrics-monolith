package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/api/controller"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/application/repository"
	"github.com/hasssanezzz/clicky-metrics-monolith/internal/application/usecase"
	"github.com/jmoiron/sqlx"
)

func SetupLoginRoute(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	con := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(userRepo),
		Env:          env,
	}
	group.POST("/login", con.Execute)
}

func SetupSignupRoute(env *bootstrap.Env, db *sqlx.DB, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	con := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepo),
		Env:           env,
	}
	group.POST("/signup", con.Execute)
}
