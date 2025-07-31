package spotify

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type SpotifyResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *Outbound) GetTokenDetails() (string, string, error) {
	if o.accessToken == "" || time.Now().After(o.expiredTime) {
		if err := o.generateToken(); err != nil {
			return "", "", err
		}

	}
	return o.accessToken, o.tokenType, nil
}

func (o *Outbound) generateToken() error {
	formData := url.Values{}
	formData.Set("grand_type", "client_credentials")
	formData.Set("client_id", o.cfg.SpotifyConfig.ClientId)
	formData.Set("client_secret", o.cfg.SpotifyConfig.ClientSecret)

	encodedUrl := formData.Encode()
	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(encodedUrl))
	if err != nil {
		log.Error().Err(err).Msg("error create request for spotify")
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get response")
		return err
	}
	defer response.Body.Close()

	var spotifyResponse SpotifyResponse
	err = json.NewDecoder(response.Body).Decode(&spotifyResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decode parsing response api spotify")
		return err
	}
	o.accessToken = spotifyResponse.AccessToken
	o.tokenType = spotifyResponse.TokenType
	o.expiredTime = time.Now().Add(time.Duration(spotifyResponse.ExpiresIn) * time.Second)

	return nil
}
