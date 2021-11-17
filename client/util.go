package client

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"perun.network/go-perun/wallet"
)

func (c *Client) PerunAddress() wallet.Address {
	return c.perunClient.Account.Address()
}

func (c *Client) Address() common.Address {
	return c.perunClient.Account.Account.Address
}

func (c *Client) challengeDurationInSeconds() uint64 {
	return uint64(c.challengeDuration.Seconds())
}

func (c *Client) Logf(format string, v ...interface{}) {
	log.Printf("Client %v: %v", c.Address(), fmt.Sprintf(format, v...))
}

func (c *Client) OnChainBalance() (b *big.Int, err error) {
	return c.perunClient.EthClient.BalanceAt(context.TODO(), c.Address(), nil)
}
