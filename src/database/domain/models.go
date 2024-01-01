package domain

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model
	ISRC       string   `json:"isrc" db:"isrc"`
	ImageUrl   string   `json:"image_url" db:"image_url"`
	Title      string   `json:"title" db:"title"`
	Popularity int      `json:"popularity" db:"popularity"`
	Artists    []Artist `gorm:"many2many:track_artists;"`
}

type Artist struct {
	gorm.Model
	Name string `json:"name" db:"name"`
}
