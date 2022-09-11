package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	MDB    Mongo
	Server Server
	Client Client
}

type Mongo struct {
	URI        string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Server struct {
	Host string
	Port int
}

type Client struct {
	Timeout time.Duration
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.MDB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := envconfig.Process("client", &cfg.Client); err != nil {
		return nil, err
	}

	return cfg, nil
}
