package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/pkg/httpclient"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOutbound_Search(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	// Create mock HTTP client
	mockHTTPClient := httpclient.NewMockHTTPClient(gomockController)

	type args struct {
		ctx    context.Context
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(mockClient *httpclient.MockHTTPClient, args args)
	}{
		{
			name: "success case",
			args: args{
				ctx:    context.Background(),
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: TrackSearchSpotify{
					Href: func() *string {
						s := `https://api.spotify.com/v1/search?offset=0&limit=20&query=bohemian%20rhapsody&type=track&market=ID`
						return &s
					}(),
					Limit:  10,
					Offset: 0,
					Next: func() *string {
						s := `https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID`
						return &s
					}(),
					Previous: nil,
					Total:    1000,
					Items: []SpotifyItemTrack{
						{
							Albums: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []SpotifyImagesAlbumObject{
									{
										Url:    "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
										Height: 640,
										Width:  640,
									},
									{
										Url:    "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
										Height: 300,
										Width:  300,
									},
									{
										Url:    "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
										Height: 64,
										Width:  64,
									},
								},
								Name:                 "Bohemian Rhapsody (The Original Soundtrack)",
								ReleaseDate:          "2018-10-19",
								ReleaseDatePrecision: "day",
							},
							Artists: []SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							ID:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
						{
							Albums: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 12,
								Images: []SpotifyImagesAlbumObject{
									{
										Url:    "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
										Height: 640,
										Width:  640,
									},
									{
										Url:    "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
										Height: 300,
										Width:  300,
									},
									{
										Url:    "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
										Height: 64,
										Width:  64,
									},
								},
								Name:                 "A Night At The Opera (2011 Remaster)",
								ReleaseDate:          "1975-11-21",
								ReleaseDatePrecision: "day",
							},
							Artists: []SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
							ID:       "4u7EnebtmKWzUH433cf5Qv",
							Name:     "Bohemian Rhapsody - Remastered 2011",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(mockClient *httpclient.MockHTTPClient, args args) {
				urlparams := "https://api.spotify.com/v1/search"
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("market", "ID")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				urlpath := fmt.Sprintf("%s?%s", urlparams, params.Encode())

				request, err := http.NewRequest(http.MethodPost, urlpath, nil)
				assert.NoError(t, err)
				request.Header.Set("Authorization", "Bearer  test-token")

				// JSON response yang sesuai dengan expected struct
				// jsonResponse := `{
				// 	"tracks": {
				// 		"href": "https://api.spotify.com/v1/search?q=bohemian+rhapsody&type=track&market=ID&limit=20&offset=0",
				// 		"limit": 10,
				// 		"offset": 0,
				// 		"next": "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID",
				// 		"previous": nil,
				// 		"total": 1000,
				// 		"items": [
				// 			{
				// 				"albums": {
				// 					"album_type": "album",
				// 					"total_tracks": 22,
				// 					"images": [
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
				// 							"height": 640,
				// 							"width": 640
				// 						},
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
				// 							"height": 300,
				// 							"width": 300
				// 						},
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
				// 							"height": 64,
				// 							"width": 64
				// 						}
				// 					],
				// 					"name": "Bohemian Rhapsody (The Original Soundtrack)",
				// 					"release_date": "2018-10-19",
				// 					"release_date_precision": "day"
				// 				},
				// 				"artists": [
				// 					{
				// 						"href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
				// 						"name": "Queen"
				// 					}
				// 				],
				// 				"explicit": false,
				// 				"href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
				// 				"id": "3z8h0TU7ReDPLIbEnYhWZb",
				// 				"name": "Bohemian Rhapsody"
				// 			},
				// 			{
				// 				"albums": {
				// 					"album_type": "album",
				// 					"total_tracks": 12,
				// 					"images": [
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
				// 							"height": 640,
				// 							"width": 640
				// 						},
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
				// 							"height": 300,
				// 							"width": 300
				// 						},
				// 						{
				// 							"url": "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
				// 							"height": 64,
				// 							"width": 64
				// 						}
				// 					],
				// 					"name": "A Night At The Opera (2011 Remaster)",
				// 					"release_date": "1975-11-21",
				// 					"release_date_precision": "day"
				// 				},
				// 				"artists": [
				// 					{
				// 						"href": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d",
				// 						"name": "Queen"
				// 					}
				// 				],
				// 				"explicit": false,
				// 				"href": "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
				// 				"id": "4u7EnebtmKWzUH433cf5Qv",
				// 				"name": "Bohemian Rhapsody - Remastered 2011"
				// 			}
				// 		]
				// 	}
				// }`

				response := &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(searchResponse))),
					Header:     make(http.Header),
				}

				mockClient.EXPECT().Do(request).Return(response, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(mockHTTPClient, tt.args)
			}

			o := &Outbound{
				cfg:         configs.Config{},
				client:      mockHTTPClient,
				accessToken: "test-token",
				tokenType:   "Bearer ",
				expiredTime: time.Now().Add(time.Duration(36000) * time.Second),
			}
			got, err := o.Search(tt.args.ctx, tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// wantttt, _ := json.Marshal(tt.want)
			// fmt.Printf("%+v\n", string(wantttt))
			// fmt.Printf("%s\n", "=======================WANT TO GOT===============================")
			// gotttt, _ := json.Marshal(got)
			// fmt.Printf("%+v\n", string(gotttt))
			t.Logf("GOT:  %+v\n\n", got)
			t.Logf("WANT: %+v", tt.want)
			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("Outbound.Search() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
