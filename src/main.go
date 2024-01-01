package main

import (
	"ltitest/src/database"
	"ltitest/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the database

	database.InitDB()

	// Define routes
	router.POST("/api/tracks", handlers.CreateTrack)
	router.GET("/api/artists/:artistName/tracks", handlers.GetTracksByArtist)
	router.GET("/api/tracks/:isrc", handlers.GetTrackByISRC)

	router.Run(":8081")
}
