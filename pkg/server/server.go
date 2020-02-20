package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type userHandler interface {
	GetLikes(w http.ResponseWriter, r *http.Request)
	EditProfile(w http.ResponseWriter, r *http.Request)
}

type db interface {
	BasicAuth(string, string) (int, error)
}

type Server struct {
	Router      chi.Router
	Database    db
	Log         zerolog.Logger
	UserHandler userHandler
}

func (s *Server) AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, apiKey, ok := r.BasicAuth()
		if !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Could cache / memoize this in a production system
		userID, err := s.Database.BasicAuth(username, apiKey)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		if userID < 0 {
			s.Log.Warn().Str("username", username).Msg("unauthorized access of api")
			http.Error(w, http.StatusText(403), 403)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) InitRoutes(r chi.Router) *Server {
	r.Route("/user", func(r chi.Router) {
		r.Use(s.AuthCtx)
		r.Get("/likes", s.UserHandler.GetLikes)
		r.Put("/profile", s.UserHandler.EditProfile)
	})
	s.Router = r
	return s
}

func (s *Server) Run() error {
	err := http.ListenAndServe(":8000", s.Router)
	if err != nil {
		return err
	}
	return nil
}
