package main

import (
	"newsfeeder/cmd/httpd/handler"
	"newsfeeder/cmd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello")
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed/post", handler.NewsFeedPost(feed))
	r.Run()
	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"title", "post"})
	// fmt.Println(feed)
}
