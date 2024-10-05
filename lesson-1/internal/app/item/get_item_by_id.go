package item

import (
	"encoding/json"
	stderr "errors"
	"net/http"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/api"
	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/errors"
)

func (i *Implementation) GetItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemID, ok := api.PathValueOrError(w, r, "item_id")
	if !ok {
		return
	}

	item, err := i.service.GetItemByID(r.Context(), itemID)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "item not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to get item by id",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(GetItemByIDResponse{
		Item: item,
	})
}
