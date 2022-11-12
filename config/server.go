package config

import "time"

type Server struct {
	Host            string
	Port            int
	GracefulTimeout time.Duration `default:"3s"`
	ReadTimeout     time.Duration `default:"5s"`
	WriteTimeout    time.Duration `default:"5s"`
}
