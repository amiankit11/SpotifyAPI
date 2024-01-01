package handlers

import (
	"fmt"
	"ltitest/src/database/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Track info on the basis of artist
func GetTracksByArtist(c *gin.Context) {
	artistName := c.Param("artistName")
	fmt.Println("GetTracksByArtist", artistName)
	tracks, err := repository.GetTracksByArtist(artistName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tracks"})
		return
	}

	c.JSON(http.StatusOK, tracks)
}
