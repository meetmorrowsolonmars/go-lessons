package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/meetmorrowsolonmars/go-lessons/testing/coverage/internal/domain"
)

type IsEvenRequest struct {
	Number int64 `json:"number"`
}

type IsEvenResponse struct {
	IsEven bool `json:"isEven"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Change the function declaration to
//
// var IsEvenHandleFunc http.HandlerFunc = func (w http.ResponseWriter, r *http.Request) { }
//
// and run the unit tests. After that, check the code coverage with -func
// option and explain the program output.

func IsEvenHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := IsEvenRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := ErrorResponse{
			Code:    "invalid_request",
			Message: fmt.Sprintf("invalid request body: %s", err),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			slog.Error("failed to encode response", "error", err)
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	response := IsEvenResponse{
		IsEven: domain.IsEvenNumber(body.Number),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
