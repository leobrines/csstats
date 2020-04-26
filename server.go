package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func initServer() {
	router := gin.Default()
	router.GET("/top", TopHandler)

	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func TopHandler(c *gin.Context) {
	players, err := GetTopPlayers()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"players": players,
	})
}
