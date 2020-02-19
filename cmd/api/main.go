package main

import (
	"os"

	"github.com/go-chi/chi"
	"github.com/rk23/hinge/pkg/server"
	"github.com/rk23/hinge/pkg/user"
	"github.com/rs/zerolog"
)

func main() {
	r := chi.NewRouter()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	userHandler := user.Handler{
		Log: log,
	}
	server.Init(r, log, userHandler).Run()
}
