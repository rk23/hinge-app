package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Profile is the representation of a user passed internally from servers to iOS or Android devices running
// the apllication
type Profile struct {
	FirstName *string `json:"first_name"`
	ID        *int    `json:"id"`
	LastName  *string `json:"last_name"`
}

func validateProfile(p Profile) error {
	if p.FirstName != nil && len(*p.FirstName) == 0 {
		return errors.New("first name must not be empty")
	}
	// More validation, like checking character counts in prompts or verifying a picture url is valid,
	// would go here
	return nil
}

func (h Handler) EditProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		// UserID ought to be in the context - if not it's server's fault for failing request
		h.Log.Error().Int("userID", userID).Msg("failed to get userID from context")
		http.Error(w, string(internalServerErr), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Msg("failed read UpdateProfile body")
		http.Error(w, string(unprocessableErr), http.StatusUnprocessableEntity)
		return
	}

	var p Profile
	err = json.Unmarshal(bodyBytes, &p)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Msg("failed to unmarshal UpdateProfile request body")
		http.Error(w, string(unprocessableErr), http.StatusUnprocessableEntity)
		return
	}

	err = validateProfile(p)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Interface("profile", p).Msg("invalid profile update")
		http.Error(w, string(unprocessableErr)+": "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.DB.EditProfile(userID, p)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Msg("failed to update profile")
		http.Error(w, string(internalServerErr), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}
