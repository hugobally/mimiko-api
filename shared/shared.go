package shared

import (
	"github.com/hugobally/mimiko/backend/auth/jwt"
	"github.com/hugobally/mimiko/backend/config"
	"github.com/hugobally/mimiko/backend/prisma"
	"github.com/hugobally/mimiko/backend/spotify"
	"log"
	"net/http"
)

type Services struct {
	Logger     *log.Logger
	Prisma     *prisma.Client
	HttpClient *http.Client
	Config     *config.Config
	JwtUtil    *jwt.Util
	Spotify    *spotify.Client
}

func (s *Services) SetLogger(l *log.Logger) {
	s.Logger = l
}

func (s *Services) SetPrisma(p *prisma.Client) {
	s.Prisma = p
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

func NewServices() *Services {
	return &Services{}
}
