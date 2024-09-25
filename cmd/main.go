package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasssanezzz/clicky-metrics-monolith/api/route"
	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
)

func main() {
	env := bootstrap.NewEnv()
	gin := gin.Default()
	route.Setup(env, gin)
	gin.Run(env.ServerAddress)
}
