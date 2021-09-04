package explorer

import (
	"github.com/gin-gonic/gin"
)

// WIP
// Search : Performs a lookup in the database using user-defined keyword
func Search(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
		return
	}

	c.JSON(200, gin.H{
		"query": query,
	})
}

func List(c *gin.Context) {
	extensions := []string{".mkv", ".mp4", ".avi", ".wmv", ".mov", ".webm"}
	videoList := find("/data", extensions)

	if len(videoList) == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"Error": "No movies were found",
		})
		return
	}

	c.JSON(200, videoList)
}
