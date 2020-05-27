package http

import (
	"github.com/gin-gonic/gin"
	"upper-tweet/posts"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()
	postsRepository := posts.Repository{}

	router.GET("/posts", posts.GetAllRoute(&postsRepository))
	return router
}