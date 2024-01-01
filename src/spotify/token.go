package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	model "ltitest/src/model"
	"net/http"
	"net/url"
)

const (
	clientID     = "2bc793dedbad4f54b376b96e2737a21a"
	clientSecret = "31a6ae0e360b465cb69a573e257d2d9f"
)

// GetToken from spotify
func GetToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access token not found in response")
	}

	return accessToken, nil
}

// Fetch the data from spotify on ISRC
func GetTrackDataByISRC(accessToken string, isrc string) (*model.TracksResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?q=isrc:"+isrc+"&amp;type=track", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(isrc)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tracksResponse model.TracksResponse
	err = json.NewDecoder(resp.Body).Decode(&tracksResponse)

	fmt.Println(tracksResponse)
	if err != nil {
		return nil, err
	}

	return &tracksResponse, nil
}
