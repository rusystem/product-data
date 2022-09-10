package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MDB    Mongo
	Server Server
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

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("mongodb", &cfg.MDB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
