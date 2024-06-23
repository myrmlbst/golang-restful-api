package api

import (
	"encoding/json"
	"net/http"
)

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	Code int // success code, usually 200
	
	Balance int64 // account balance
}

// error response
type Error struct {
	Code int // error code

	Message string // error message
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "applicacion/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), http.StatusBadRequest)	
	}
	InternalErrorHandler = func (w http.ResponseWriter)  {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}
)
