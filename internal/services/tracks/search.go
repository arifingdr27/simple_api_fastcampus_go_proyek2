package tracks

import (
	"context"
	spotifymodel "proyek3-catalog-music/internal/models/spotify"
	"proyek3-catalog-music/internal/repository/spotify"
)

func (s *Service) Search(ctx context.Context, query string, pageLimit, pageIndex int) (*spotifymodel.SearchResponse, error) {
	limit := pageLimit
	offset := (pageIndex - 1) * pageLimit

	trackDetail, err := s.SpotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return modelToResponse(trackDetail), nil

}

func modelToResponse(data *spotify.SpotifySearchResponse) *spotifymodel.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotifymodel.SpotifyItemTrack, 0)

	for _, v := range data.Tracks.Items {

		var groupArtistName []string
		for _, artist := range v.Artists {
			groupArtistName = append(groupArtistName, artist.Name)
		}

		var imageAlbumUrl []string
		for _, image := range v.Albums.Images {
			imageAlbumUrl = append(imageAlbumUrl, image.Url)
		}

		items = append(items, spotifymodel.SpotifyItemTrack{
			AlbumType:        v.Albums.AlbumType,
			AlbumTotalTracks: v.Albums.TotalTracks,
			AlbumImages:      imageAlbumUrl,
			AlbumName:        v.Albums.Name,

			ArtistsName: groupArtistName,
			Explicit:    v.Explicit,
			ID:          v.ID,
			Name:        v.Name,
		})
	}

	return &spotifymodel.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
