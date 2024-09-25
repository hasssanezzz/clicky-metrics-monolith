package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAUthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		autherHeader := c.Request.Header.Get("Authorization")
		if len(autherHeader) == 0 {
			c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Not authorized"})
			c.Abort()
		}
	}
}
