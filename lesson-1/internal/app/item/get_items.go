package item

import (
	"encoding/json"
	"net/http"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/api"
)

func (i *Implementation) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := i.service.GetItems(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to get items",
		})
		return
	}

	response := GetItemsResponse{
		Items: items,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
