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
