package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"upper-tweet/posts"
)

func GetAllRoute(repository posts.Repository) gin.HandlerFunc {
	return func (c *gin.Context) {
		posts, err := repository.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{ "posts:": posts })
	}
}
