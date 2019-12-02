package handler

import "github.com/gin-gonic/gin"

import "net/http"

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "ok",
		})
	}
}
