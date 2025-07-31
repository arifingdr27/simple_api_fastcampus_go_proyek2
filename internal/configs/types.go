package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotifyClient"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJwt string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientId     string `mapstructure:"clientId"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)
