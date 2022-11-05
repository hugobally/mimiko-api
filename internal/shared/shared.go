package shared

import (
	"github.com/hugobally/mimiko_api/internal/authentication/jwt"
	"github.com/hugobally/mimiko_api/internal/authentication/spotify"
	"github.com/hugobally/mimiko_api/internal/config"
	"github.com/hugobally/mimiko_api/internal/db"
	"log"
	"net/http"
)

type Services struct {
	Logger     *log.Logger
	HttpClient *http.Client
	Config     *config.Config
	JwtUtil    *jwt.Util
	Spotify    *spotify.Client
	Database   *db.Client
}

func (s *Services) SetLogger(l *log.Logger) {
	s.Logger = l
}

func (s *Services) SetHttpClient(c *http.Client) {
	s.HttpClient = c
}

func (s *Services) SetConfig(c *config.Config) {
	s.Config = c
}

func (s *Services) SetJwtUtil(j *jwt.Util) {
	s.JwtUtil = j
}

func (s *Services) SetSpotify(sp *spotify.Client) {
	s.Spotify = sp
}

func (s *Services) SetDatabase(c *db.Client) {
	s.Database = c
}

func NewServices() *Services {
	return &Services{}
}
