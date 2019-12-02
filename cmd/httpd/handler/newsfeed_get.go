package handler

import (
	"net/http"
	"newsfeeder/cmd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
