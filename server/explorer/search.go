package explorer

import "github.com/gin-gonic/gin"

// WIP
// Search : Performs a lookup in the database using user-defined keyword
func Search(c *gin.Context) {
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

func List(c *gin.Context) {
	var videoList ListResponse
	var videoItem ListItem

	videoItem.Title = "Test title"
	videoItem.Runtime = "1:43:57"
	videoItem.Genre = "Test genre"
	videoItem.Description = "Test description"

	videoList.Items = append(videoList.Items, videoItem)

	c.JSON(200, videoList.Items)
}
