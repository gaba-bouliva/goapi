package tools

import (
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"eagle": {
		AuthToken: "1234ABC",
		Username:  "Eagle Bald",
	},
	"bald": {
		AuthToken: "5678DEF",
		Username:  "Bald eagle",
	},
	"soldierx": {
		AuthToken: "0238XYZ",
		Username:  "Soldier X",
	},
	"king": {
		AuthToken: "9083KNG",
		Username:  "Cyber King",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"eagle": {
		Coins:    100,
		Username: "Eagle Bald",
	},
	"bald": {
		Coins:    200,
		Username: "Bald eagle",
	},
	"soldierx": {
		Coins:    300,
		Username: "Soldier X",
	},
	"king": {
		Coins:    400,
		Username: "Cyber King",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB Call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simple DB call
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