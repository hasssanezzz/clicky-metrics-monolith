package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/api/middleware"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, db *sqlx.DB, gin *gin.Engine) {
	version := gin.Group("/v1")

	publicRouter := version.Group("/auth")
	privateRouter := version.Group("")
	privateRouter.Use(middleware.JWTAUthMiddleware(env.AccessTokenSecret))

	// public routes
	SetupLoginRoute(env, db, publicRouter)
	SetupSignupRoute(env, db, publicRouter)
	SetupRefreshTokenRoute(env, db, publicRouter)

	// private routes
}
