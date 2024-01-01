package handlers

import (
	"fmt"
	"net/http"
	"sort"

	"ltitest/src/database/domain"
	"ltitest/src/database/repository"
	"ltitest/src/spotify"

	"github.com/gin-gonic/gin"
)

type TrackRequest struct {
	ISRC string `json:"isrc" binding:"required"`
}

// Store the data from fetching the data from spotify by ISRC
func CreateTrack(c *gin.Context) {
	var trackRequest TrackRequest

	if err := c.ShouldBindJSON(&trackRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Getting token from spotify
	token, err := spotify.GetToken()
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	//fetching track from spotify based on isrc
	tracksResponse, err := spotify.GetTrackDataByISRC(token, trackRequest.ISRC)
	if err != nil {
		fmt.Println("Error getting track info:", err)
		return
	}

	fmt.Println("tracksResponse", tracksResponse)

	// Sort tracks by popularity in descending order
	sort.Slice(tracksResponse.Tracks.Items, func(i, j int) bool {
		return tracksResponse.Tracks.Items[i].Popularity > tracksResponse.Tracks.Items[j].Popularity
	})

	// Get data from the track with the highest popularity
	highestPopularityTrack := tracksResponse.Tracks.Items[0]
	previewURL := highestPopularityTrack.PreviewURL
	trackName := highestPopularityTrack.Name

	// Build a list of artist names
	var listOfArtists []domain.Artist
	for _, artist := range highestPopularityTrack.Artists {
		listOfArtists = append(listOfArtists, domain.Artist{
			Name: artist.Name,
		})
	}

	track := domain.Track{
		ISRC:       trackRequest.ISRC,
		ImageUrl:   previewURL,
		Title:      trackName,
		Popularity: highestPopularityTrack.Popularity,
		Artists:    listOfArtists,
	}

	//saving it to db
	err = repository.CreateTrack(&track)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create track"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Track created successfully"})
}

// Fetch track info on the basis of ISRC
func GetTrackByISRC(c *gin.Context) {
	isrc := c.Param("isrc")

	track, err := repository.GetTrackByISRC(isrc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch track"})
		return
	}

	c.JSON(http.StatusOK, track)
}
