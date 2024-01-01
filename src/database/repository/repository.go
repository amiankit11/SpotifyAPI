package repository

import (
	"ltitest/src/database"
	"ltitest/src/database/domain"
)

func CreateTrack(track *domain.Track) error {
	return database.DB.Create(track).Error
}

func GetTrackByISRC(isrc string) (*domain.Track, error) {
	var track domain.Track
	if err := database.DB.Where("isrc = ?", isrc).Preload("Artists").First(&track).Error; err != nil {
		return nil, err
	}

	return &track, nil
}

func GetTracksByArtist(artistName string) ([]domain.Track, error) {
	var tracks []domain.Track
	if err := database.DB.Joins("JOIN track_artists ON tracks.id = track_artists.track_id").
		Joins("JOIN artists ON track_artists.artist_id = artists.id").
		Where("artists.name LIKE ?", "%"+artistName+"%").
		Preload("Artists").
		Find(&tracks).Error; err != nil {
		return nil, err
	}

	return tracks, nil
}
