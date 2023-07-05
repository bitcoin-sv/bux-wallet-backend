package data

import (
	buxclient "bux-wallet/transports/bux/client"

	"github.com/brianvoe/gofakeit/v6"
)

// CreateTestTransactions returns 'count' randomly generated transactions.
func CreateTestTransactions(count int) []buxclient.FullTransaction {
	result := make([]buxclient.FullTransaction, count)
	gofakeit.Slice(&result)

	return result
}
