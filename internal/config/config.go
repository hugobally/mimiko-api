package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Env string `yaml:"env" envconfig:"MIMIKO_ENV"`
	Tls struct {
		Cert string `yaml:"cert" envconfig:"MIMIKO_CERT"`
		Key  string `yaml:"key" envconfig:"MIMIKO_KEY"`
	} `yaml:"tls"`
	Server struct {
		Host       string `yaml:"host" envconfig:"MIMIKO_SERVER_HOST"`
		Port       int    `yaml:"port" envconfig:"MIMIKO_SERVER_PORT"`
		StaticPath string `yaml:"static_path" envconfig:"MIMIKO_STATIC_PATH"`
	} `yaml:"server"`
	Auth struct {
		JwtKey string `yaml:"jwt_key" envconfig:"MIMIKO_JWT_KEY"`
	} `yaml:"authentication"`
	Spotify struct {
		ClientId     string `yaml:"client_id" envconfig:"SPOTIFY_CLIENT_ID"`
		ClientSecret string `yaml:"client_secret" envconfig:"SPOTIFY_CLIENT_SECRET"`
		RedirectUri  string `yaml:"redirect_uri" envconfig:"SPOTIFY_REDIRECT_URI"`
	} `yaml:"spotify"`
}

func New() *Config {
	var v Config
	readFile(&v)
	readEnv(&v)
	return &v
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(v *Config) {
	f, err := os.Open("configs/config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(v)
	if err != nil {
		processError(err)
	}
}

func readEnv(v *Config) {
	err := envconfig.Process("", v)
	if err != nil {
		processError(err)
	}
}
