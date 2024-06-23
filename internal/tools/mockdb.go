package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"andrew": {
		AuthToken: "123ABC",
		Username: "loremipsum",
	},
	"mark": {
		AuthToken: "456DEF",
		Username: "dolor",

	},
	"miguel": {
		AuthToken: "789GHI",
		Username: "sitamet",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"andrew": {
		Coins: 400,
		Username: "loremipsum",
	},
	"mark": {
		Coins: 150,
		Username: "dolor",

	},
	"miguel": {
		Coins: 350,
		Username: "sitamet",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil

	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
