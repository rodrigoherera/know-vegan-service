package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement your authentication logic here

		// continue processing the request
		c.Next()
	}
}
