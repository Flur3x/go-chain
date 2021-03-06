package main

import (
	"math/rand"
	"time"

	"github.com/Flur3x/go-chain/api"
	"github.com/Flur3x/go-chain/blockchain"
	"github.com/Flur3x/go-chain/miner"
	"github.com/Flur3x/go-chain/transactions"
	"github.com/Flur3x/go-chain/wallet"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")
var errorReport = make(chan error)

func main() {
	go runSimulation()

	handleErrors()
}

func runSimulation() {
	blockchain.New()

	myWallet, err := wallet.New()

	if err != nil {
		errorReport <- err
	}

	foreignWallet, err := wallet.New()

	if err != nil {
		errorReport <- err
	}

	log.Infof("\nWallets created:\n%+v\n%+v\n", myWallet, foreignWallet)

	go api.Start(errorReport)
	go miner.Start(errorReport)

	log.Info("Simulation started 🌈\n\nFake Transactions are being created and Blocks mined ...\n\n")

	for range time.NewTicker(5 * time.Second).C {
		randomAmount := uint64(rand.Int63n(10000))
		fakeTransaction, err := transactions.New(myWallet.Address, foreignWallet.Address, randomAmount, myWallet)

		if err != nil {
			errorReport <- err
		}

		transactions.UpdateOrAdd(fakeTransaction)
	}
}

func handleErrors() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("%+v\n", r)
			log.Infof("More detailed logs in the errors.log file\n")
		}
	}()

	for r := range errorReport {
		panic("Client crashed. Error: " + r.Error())
	}
}
