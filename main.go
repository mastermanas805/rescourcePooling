package main

import (
	"github.com/gin-gonic/gin"
	"main/api"
)

func main() {
	r := gin.Default()
	r.GET("/execute", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"uniqueID": api.Execute(),
		})
	})
	r.GET("/poll/:uniqueID", func(c *gin.Context) {
		uniqueID := c.Param("uniqueID")
		c.JSON(200, gin.H{
			"state": PoolByID(uniqueID),
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func PoolByID(id string) string {
	switch api.Poll(id) {
	case api.Completed:
		return "Completed"
	case api.Queued:
		return "Queued"
	case api.Processing:
		return "Processing"
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return "Unknown"
	}
}
