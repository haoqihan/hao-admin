package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog_server")
		c.Set("app_version", "0.1")
		c.Next()
	}
}
