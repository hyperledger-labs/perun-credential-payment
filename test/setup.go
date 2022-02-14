package test

import (
	"context"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/perun-network/perun-credential-payment/client"
	"github.com/stretchr/testify/require"
	"perun.network/go-perun/wire"
)

const (
	chainURL = "ws://127.0.0.1:8545"
	chainID  = 1337

	DeployerSK = "50b4713b4ba55b6fbcb826ae04e66c03a12fc62886a90ca57ab541959337e897"
	HolderSK   = "1af2e950272dd403de7a5760d41c6e44d92b6d02797e51810795ff03cc2cda4f"
	IssuerSK   = "f63d7d8e930bccd74e93cf5662fde2c28fd8be95edb70c73f1bdd863d07f412e"

	txFinality      = 1
	disputeDuration = 3 * time.Second
)

type Environment struct {
	Holder, Issuer *client.Client
	balanceLogger  balanceLogger
}

func (e *Environment) LogAccountBalances() {
	e.balanceLogger.LogBalances(e.Holder, e.Issuer)
}

// Setup deploys the contracts, and then creates and starts two Perun clients
// which represent the holder and the issuer.
func Setup(t *testing.T) *Environment {
	t.Helper()
	require := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// Deploy contracts.
	log.Print("Deploying contracts...")
	deployerKey, err := crypto.HexToECDSA(DeployerSK)
	require.NoError(err, "creating deployer private key")
	contracts, err := deployContracts(ctx, chainURL, big.NewInt(chainID), deployerKey)
	require.NoError(err, "deploying contracts")

	log.Print("Setting up clients...")

	// Create message bus used for off-chain communication.
	bus := wire.NewLocalBus()

	// Setup holder.
	holderConfig := newClientConfig(chainURL, contracts, HolderSK, bus)
	holder, err := client.StartClient(ctx, holderConfig)
	require.NoError(err, "Holder setup")
	t.Cleanup(holder.Shutdown)

	// Setup issuer.
	issuerConfig := newClientConfig(chainURL, contracts, IssuerSK, bus)
	issuer, err := client.StartClient(ctx, issuerConfig)
	require.NoError(err, "Issuer setup")
	t.Cleanup(issuer.Shutdown)
	log.Print("Setup done.")

	return &Environment{
		Holder:        holder,
		Issuer:        issuer,
		balanceLogger: newBalanceLogger(chainURL),
	}
}

func newClientConfig(
	nodeURL string,
	contracts ContractAddresses,
	privateKey string,
	bus *wire.LocalBus,
) client.ClientConfig {
	sk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return client.ClientConfig{
		PrivateKey:        sk,
		ETHNodeURL:        nodeURL,
		Adjudicator:       contracts.Adjudicator,
		AssetHolder:       contracts.AssetHolder,
		AppAddress:        contracts.App,
		TxFinality:        txFinality,
		ChainID:           big.NewInt(chainID),
		Bus:               bus,
		ChallengeDuration: disputeDuration,
	}
}
