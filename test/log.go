package test

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/perun-network/perun-credential-payment/client"
)

// balanceLogger is a utility for logging client balances.
type balanceLogger struct {
	ethClient *ethclient.Client
}

// newBalanceLogger creates a new balance logger for the specified ledger.
func newBalanceLogger(chainURL string) balanceLogger {
	c, err := ethclient.Dial(chainURL)
	if err != nil {
		panic(err)
	}
	return balanceLogger{ethClient: c}
}

// LogBalances prints the balances of the specified clients.
func (l balanceLogger) LogBalances(clients ...*client.Client) {
	bals := make([]*big.Float, len(clients))
	for i, c := range clients {
		bal, err := l.ethClient.BalanceAt(context.TODO(), c.EthAddress(), nil)
		if err != nil {
			log.Fatal(err)
		}
		bals[i] = WeiToEth(bal)
	}
	log.Println("Client balances (ETH):", bals)
}
