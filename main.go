package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/srrfrhmnn/short-url/handler"
	"github.com/srrfrhmnn/short-url/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Initialize the store service
	store.InitStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
