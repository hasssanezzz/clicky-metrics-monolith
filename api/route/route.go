package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/api/middleware"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
)

func testPublicHandler(g *gin.Context) {
	g.JSON(200, map[string]string{
		"msg": "Hello world",
	})
}

func testPrivateHandler(g *gin.Context) {
	g.JSON(200, map[string]string{
		"msg": "some secret",
	})
}

func Setup(env *bootstrap.Env, gin *gin.Engine) {
	publicRouter := gin.Group("")
	publicRouter.GET("/", testPublicHandler)
	privateRouter := gin.Group("")
	privateRouter.Use(middleware.JWTAUthMiddleware(env.AccessTokenSecret))
	privateRouter.GET("/data", testPrivateHandler)

}
