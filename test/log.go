package test

import (
	"fmt"
	"log"
	"math/big"

	"github.com/perun-network/verifiable-credential-payment/client"
)

func LogAccountBalance(clients ...*client.Client) {
	for _, c := range clients {
		globalBalance, err := c.OnChainBalance()
		if err != nil {
			log.Panicf("Could not retrieve balance for %v: %v", c.Address(), err)
		}
		log.Printf("%v: Account Balance: %v", c.Address(), toEth(globalBalance))
	}
}

func toEth(weiAmount *big.Int) string {
	return fmt.Sprintf("%vETH", WeiToEth(weiAmount))
}
