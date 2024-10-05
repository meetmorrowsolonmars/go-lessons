package card

import (
	"encoding/json"
	stderr "errors"
	"net/http"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/api"
	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/errors"
)

func (i *Implementation) GetByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}

	card, err := i.cardService.GetByUserID(r.Context(), userID)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "card not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to get card by user id",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(DefaultResponse{
		Card: card,
	})
}
