package handler

import (
	"net/http"
	"newsfeeder/cmd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {

	return func(c *gin.Context) {
		request := newsFeedPostRequest{}
		c.Bind(&request)
		item := newsfeed.Item{
			Title: request.Title,
			Post:  request.Body,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
