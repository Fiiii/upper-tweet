package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRoute(repository INotesRepository) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"posts": repository.GetAll()})
	}
}
