package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/demo/contracts"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	c := cron.New()
	c.AddFunc("0 */10 * * *", func() {
		contracts.CheckGameEndTime()
	})

	// Start cron with one scheduled job
	fmt.Println("Start cron")
	c.Start()

	router := gin.Default()
	router.GET("/biddingwar/result", getDeclareResult)
	router.GET("/biddingwar/events", getEventLogs)

	router.Run("localhost:8080")

}

// getAlbums responds with the list of all albums as JSON.
func getDeclareResult(c *gin.Context) {
	err := contracts.DeclareResultByOwner()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Placed successfully"})
}

// getEventLogs
func getEventLogs(c *gin.Context) {
	// id := c.Param("id")
	bids, wins := contracts.ReadLogs()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Found", "bids": bids, "wins": wins})
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Log not found"})
}
