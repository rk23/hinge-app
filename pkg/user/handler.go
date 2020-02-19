package user

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type Handler struct {
	Log zerolog.Logger
	DB  *sql.DB
}
