package user

import "net/http"

type Profile struct {
	FirstName     string `json:"first_name"`
	ID            string `json:"id" `
	IsRecommended bool   `json:"is_recommended"`
	LastName      string `json:"last_name"`
}

func (h Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {}
