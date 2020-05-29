package http

import (
	"github.com/gin-gonic/gin"
	"upper-tweet/posts"
	"upper-tweet/posts/web"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()
	postsRepository := posts.RepositoryImpl{}

	router.GET("/posts", web.GetAllRoute(&postsRepository))
	return router
}