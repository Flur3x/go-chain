package transactions

import (
	"fmt"
	"strings"

	"github.com/Flur3x/go-chain/wallet"
	"github.com/google/uuid"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

// Address of an account. Can receive or send value.
type Address = wallet.Address

// Input defines how much value an Address receives.
type Input struct {
	Address Address
	Amount  uint64
}

// Output defines how much value an Address sends.
type Output struct {
	Address Address
	Amount  uint64
}

// Transaction gets created and signed by an account. Will be added to a "Block" by miners.
type Transaction struct {
	ID      uuid.UUID
	Input   Input
	Outputs []Output
}

// New creates a "Transaction" with the given data.
func New(from Address, to Address, amount uint64) Transaction {
	outputs := []Output{
		Output{to, amount},
		Output{from, 0}, // TODO - replace with something like "senderWallet.balance - amount"
	}

	return Transaction{uuid.New(), Input{from, amount}, outputs}
}

// JoinTransactionsToString takes a slice of transactions and returns it as a single string.
func JoinTransactionsToString(txs []Transaction) string {
	var stringSlice []string

	for _, tx := range txs {
		stringSlice = append(stringSlice, tx.String())
	}

	return strings.Join(stringSlice, ",")
}

func (t Transaction) String() string {
	return fmt.Sprintf("\n # ID: %s\n-> Input: %+v\n<- Outputs: %+v\n", t.ID, t.Input, t.Outputs)
}
