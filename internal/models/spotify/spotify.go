package spotify

type SearchResponse struct {
	Limit  int                `json:"limit"`
	Offset int                `json:"offset"`
	Items  []SpotifyItemTrack `json:"tracks"`
	Total  int                `json:"total"`
}

// type TrackSearchSpotify struct {
// 	Href     string           `json:"href"`
// 	Limit    int              `json:"limit"`
// 	Offset   int              `json:"offset"`
// 	Next     string           `json:"next"`
// 	Previous string           `json:"previous"`
// 	Total    int              `json:"total"`
// 	Items    SpotifyItemTrack `json:"items"`
// }

type SpotifyItemTrack struct {
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int      `json:"album_total_tracks"`
	AlbumImages      []string `json:"album_images"`
	AlbumName        string   `json:"album_name"`
	ArtistsName      []string `json:"artists_name"`
	Explicit         bool     `json:"explicit"`
	Href             string   `json:"href"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
}

type SpotifyArtistsObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}
