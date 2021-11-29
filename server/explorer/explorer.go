package explorer

import (
	"encoding/json"
	"reflect"

	"github.com/B-0-B-B-Y/bobflix/db"
	"github.com/gin-gonic/gin"
)

const VIDEO_DIR = "/home/zer0/Documents/github/bobflix/server/data"

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
	client := db.New().Client

	keys, err := client.Do(client.Context(), "keys", "*").Result()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Error": "Could not retrieve information from redis",
		})
		return
	}

	var movieList []db.Movie
	keyArray := reflect.ValueOf(keys)

	for i := 0; i < keyArray.Len(); i++ {
		movie, err := client.Get(client.Context(), keyArray.Index(i).Elem().String()).Result()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"Error": "Could not retrieve data for all movies",
			})
			return
		}

		var jsonMovie db.Movie
		err = json.Unmarshal([]byte(movie), &jsonMovie)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"Error": "Could not serialise data from redis to JSON",
			})
			return
		}

		movieList = append(movieList, jsonMovie)
	}

	c.JSON(200, gin.H{
		"movies": movieList,
	})
}
