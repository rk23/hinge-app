package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type UserHandler interface {
	GetLikes(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	Router   chi.Router
	Database *sql.DB
	Log      zerolog.Logger
	User     UserHandler
}

func Init(r chi.Router, log zerolog.Logger, user UserHandler) Server {
	r.Route("/user", func(r chi.Router) {
		r.Get("/likes", user.GetLikes)
		r.Put("/profile", user.UpdateProfile)
	})
	return Server{
		Router: r,
		Log:    log,
		User:   user,
	}
}

func (s Server) Run() error {
	http.ListenAndServe(":3000", s.Router)
	return nil
}
