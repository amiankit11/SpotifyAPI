
# Music Metadata API

This API allows you to store and retrieve metadata for music tracks from the Spotify API.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [API Endpoints](#api-endpoints)
  - [Create Track](#create-track)
  - [Get Track by ISRC](#get-track-by-isrc)
  - [Get Tracks by Artist](#get-tracks-by-artist)
- [Usage](#usage)

## Features

- Create tracks with metadata fetched from the Spotify API
- Retrieve track metadata by ISRC
- Retrieve tracks by artist name

## Getting Started

### Prerequisites

- Go installed on your machine
- PostgreSQL installed and running
- Spotify Developer account for API access

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/amiankit11/SpotifyAPI.git
   cd SpotifyAPI

1. Install dependencies:

go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm/dialects/postgres
go get -u github.com/lib/pq
go get -u github.com/joho/godotenv

2. Set up your PostgreSQL database and update the .env file with your configuration.

3. Run the application: go run main.go

## API Endpoints

### Create Track
Endpoint:

POST /api/tracks
Description:
Create a new track with metadata fetched from the Spotify API.

Request:

Body: JSON with the ISRC of the track

```bash 
curl -X POST -d '{"isrc": "USVT10300001"}' http://localhost:8080/api/tracks
```

### Get Track by ISRC
Endpoint:

```bash 
GET http://localhost:8080/api/tracks/:isrc
```

Description:
Retrieve metadata for a specific track using its ISRC.

Request:

Path Parameter: ISRC (International Standard Recording Code)

```bash 
curl http://localhost:8080/api/tracks/USVT10300001
```

### Get Tracks by Artist
Endpoint:

```bash 
GET http://localhost:8080/api/artists/:artistName/tracks
```
Description:
Retrieve tracks by a specific artist using a "like" search in the database.

Request:

Path Parameter: Artist Name

```bash 
curl http://localhost:8080/api/artists/artistName/tracks
```

# Usage
Follow the examples provided in each endpoint's description to interact with the API.
