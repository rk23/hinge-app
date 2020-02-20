package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h Handler) GetLikes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		// UserID ought to be in the context - if not it's server's fault for failing request
		h.Log.Error().Int("userID", userID).Msg("failed to get userID from context")
		http.Error(w, string(internalServerErr), http.StatusInternalServerError)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	// Could do more validation for app-based constraints here i.e. no more than 100 likes
	// show at a time
	// Note: A nil limit value is equivalent to not having a limit -
	// validation to prevent non-existent limit setting to return 0 rows
	lPtr := &limit
	if limit <= 0 {
		lPtr = nil
	}
	if offset < 0 {
		offset = 0
	}

	likes, err := h.DB.GetLikes(userID, lPtr, offset)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Msg("failed to get likes")
		http.Error(w, string(internalServerErr), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(likes)
	if err != nil {
		h.Log.Error().Err(err).Int("userID", userID).Msg("failed to marshal GetLikes response")
		http.Error(w, string(internalServerErr), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
