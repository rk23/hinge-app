package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/rk23/hinge/pkg/user"
	"github.com/rs/zerolog"
)

type Postgres struct {
	Log     zerolog.Logger
	ConnStr string
	DB      *sql.DB
}

func (p *Postgres) Open() error {
	db, err := sql.Open("postgres", p.ConnStr)
	if err != nil {
		return err
	}
	p.DB = db
	return nil
}

func (p Postgres) BasicAuth(username string, apiKey string) (int, error) {
	rows, err := p.DB.Query(`
		SELECT id
		FROM users
		WHERE first_name = $1 AND api_key = $2`, username, apiKey)
	if err != nil {
		return 0, err
	}

	if !rows.Next() {
		return -1, nil
	}

	var userID int
	err = rows.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (p Postgres) EditProfile(userID int, profile user.Profile) error {
	var args []interface{}
	query := []string{}
	if profile.FirstName != nil {
		query = append(query, "first_name = $")
		args = append(args, profile.FirstName)
	}
	if profile.LastName != nil {
		query = append(query, "last_name = $")
		args = append(args, profile.LastName)
	}

	query = append(query, "WHERE id = $")
	args = append(args, userID)

	q := ""
	for i, val := range query {
		// 1 based index, not 0
		i++
		if len(q) == 0 || strings.Contains(val, "WHERE") {
			q += fmt.Sprintf(" "+val+"%d", i)
		} else {
			q += fmt.Sprintf(", "+val+"%d", i)
		}
	}

	_, err := p.DB.Exec(`
		UPDATE users 
		SET `+q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) GetLikes(userID int, limit *int, offset int) ([]user.Profile, error) {
	rows, err := p.DB.Query(`
		SELECT users.id, users.first_name, users.last_name
		FROM users 
		JOIN relationships ON relationships.initiator_id = users.id
		WHERE relationships.receiver_id = $1
		ORDER BY relationships.last_updated
		LIMIT $2
		OFFSET $3
			`, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	likes := []user.Profile{}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		err = rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			return nil, err
		}
		likes = append(likes, user.Profile{
			ID:        &id,
			FirstName: &firstName,
			LastName:  &lastName,
		})
	}

	return likes, nil
}
