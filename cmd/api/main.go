package main

import (
	"os"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/rk23/hinge/pkg/db"
	"github.com/rk23/hinge/pkg/server"
	"github.com/rk23/hinge/pkg/user"
	"github.com/rs/zerolog"
)

func main() {
	r := chi.NewRouter()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// In a production system, the values in this connection string would be
	// in environment variables
	pg := db.Postgres{
		Log:     log,
		ConnStr: "dbname=hinge sslmode=disable user=postgres password=password",
	}
	err := pg.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db, shutting down")
		os.Exit(-1)
	}

	// While the two databases here look like a stutter, it allows for both the server and the user handler to
	// have two separate database connections. Each has its own interface, if you wanted to swap out for something
	// else you'd just have to satisfy that interface.
	s := server.Server{
		Log:      log,
		Database: pg,
		UserHandler: user.Handler{
			Log: log,
			DB:  pg,
		},
	}

	err = s.InitRoutes(r).Run()
	if err != nil {
		log.Fatal().Err(err).Msg("server failed, check logs")
	}
	log.Info().Msg("server shutting down")
}
