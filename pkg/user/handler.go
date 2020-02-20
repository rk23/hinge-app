package user

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type db interface {
	GetLikes(userID int, limit *int, offset int) ([]Profile, error)
	EditProfile(userID int, profile Profile) error
}

type errorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var internalServerErr, _ = json.Marshal(errorMessage{
	Message: http.StatusText(http.StatusInternalServerError),
	Status:  http.StatusInternalServerError,
})

var unprocessableErr, _ = json.Marshal(errorMessage{
	Message: http.StatusText(http.StatusUnprocessableEntity),
	Status:  http.StatusUnprocessableEntity,
})

type Handler struct {
	Log zerolog.Logger
	DB  db
}
