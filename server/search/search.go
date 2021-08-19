package search

import "github.com/gin-gonic/gin"

// WIP
// GETSearch : Performs a lookup in the database using user-defined keyword
func GETSearch(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
	}

	c.JSON(200, gin.H{
		"query": query,
	})
}
