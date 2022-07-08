package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
// album represents data about a record album.
// type album struct {
//     ID     string  `json:"id"`
//     Title  string  `json:"title"`
//     Artist string  `json:"artist"`
//     Price  float64 `json:"price"`
// }

// getAlbums responds with the list of all albums as JSON.
func getCloud(c *gin.Context) {
	return_message := gin.H{
		"response" : "CLOUD",
	}
    c.IndentedJSON(http.StatusOK, return_message)
}

func main() {
    router := gin.Default()
    router.GET("/cloud", getCloud)

    router.Run("0.0.0.0:8080")
}
