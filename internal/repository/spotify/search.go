package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SpotifySearchResponse struct {
	Tracks TrackSearchSpotify `json:"tracks"`
}

type TrackSearchSpotify struct {
	Href     string             `json:"href"`
	Limit    int                `json:"limit"`
	Offset   int                `json:"offset"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Total    int                `json:"total"`
	Items    []SpotifyItemTrack `json:"items"`
}

type SpotifyItemTrack struct {
	Albums   SpotifyAlbumObject     `json:"albums"`
	Artists  []SpotifyArtistsObject `json:"artists"`
	Explicit bool                   `json:"explicit"`
	Href     string                 `json:"href"`
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
}

type SpotifyAlbumObject struct {
	AlbumType            string                     `json:"album_type"`
	TotalTracks          int                        `json:"total_tracks"`
	Images               []SpotifyImagesAlbumObject `json:"images"`
	Name                 string                     `json:"name"`
	ReleaseDate          string                     `json:"release_date"`
	ReleaseDatePrecision string                     `json:"release_date_precision"`
}

type SpotifyImagesAlbumObject struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyArtistsObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

func (o *Outbound) Search(ctx context.Context, query string, limit, offset int) (*SpotifySearchResponse, error) {
	urlparams := "https://api.spotify.com/v1/search"
	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("offset", strconv.Itoa(offset))

	urlpath := fmt.Sprintf("%s?%s", urlparams, params.Encode())

	req, err := http.NewRequest(http.MethodPost, urlpath, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create request for spotify")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("error get token details")
		return nil, err
	}

	bearertoken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearertoken)

	response, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get response")
		return nil, err
	}
	defer response.Body.Close()

	var spotifyResponse SpotifySearchResponse
	err = json.NewDecoder(response.Body).Decode(&spotifyResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decode parsing response api spotify")
		return nil, err
	}

	return &spotifyResponse, nil
}
