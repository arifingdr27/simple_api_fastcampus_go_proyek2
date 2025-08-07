package spotify

var searchResponse = `{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=20&query=bohemian%20rhapsody&type=track&market=ID",
    "limit": 10,
    "offset": 0,
    "next": "https://api.spotify.com/v1/search?offset=20&limit=20&query=bohemian%20rhapsody&type=track&market=ID",
    "previous": null,
    "total": 1000,
    "items": [
      {
        "album": {
          "album_type": "album",
          "total_tracks": 22,
          "images": [
            {
              "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
              "height": 640,
              "width": 640
            },
            {
              "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
              "height": 300,
              "width": 300
            },
            {
              "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
              "height": 64,
              "width": 64
            }
          ],
          "name": "Bohemian Rhapsody (The Original Soundtrack)",
          "release_date": "2018-10-19",
          "release_date_precision": "day"
        },
        "artists": [
          {
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen"
          }
        ],
        "explicit": false,
        "href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
        "id": "3z8h0TU7ReDPLIbEnYhWZb",
        "name": "Bohemian Rhapsody"
      },
      {
        "album": {
          "album_type": "album",
          "total_tracks": 12,
          "images": [
            {
              "url": "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
              "height": 640,
              "width": 640
            },
            {
              "url": "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
              "height": 300,
              "width": 300
            },
            {
              "url": "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
              "height": 64,
              "width": 64
            }
          ],
          "name": "A Night At The Opera (2011 Remaster)",
          "release_date": "1975-11-21",
          "release_date_precision": "day"
        },
        "artists": [
          {
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen"
          }
        ],
        "explicit": false,
        "href": "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
        "id": "4u7EnebtmKWzUH433cf5Qv",
        "name": "Bohemian Rhapsody - Remastered 2011"
      }
    ]
  }
}`
