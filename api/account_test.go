package api

import (
	"testing"

	db "github.com/myboilerplate/golang-simple-bank/db/sqlc"
	"github.com/myboilerplate/golang-simple-bank/util"
)

func TestGetAccountAPI(t *testing.T) {

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
