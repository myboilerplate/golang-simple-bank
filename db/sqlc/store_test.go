package db

import "testing"

func TestTransferTx(t *testing.T) {
	// write Test Store.TransferTx
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
}
