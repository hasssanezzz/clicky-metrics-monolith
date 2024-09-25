package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/api/route"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
)

func main() {
	gin := gin.Default()
	env := bootstrap.NewEnv()
	db := bootstrap.MustSetup(env)
	route.Setup(env, db, gin)
	gin.Run(env.ServerAddress)
}
