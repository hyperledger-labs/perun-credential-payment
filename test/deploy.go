package test

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/perun-network/perun-credential-payment/app"
	"github.com/pkg/errors"
	"perun.network/go-perun/backend/ethereum/wallet"
	"perun.network/go-perun/channel"
)

type ContractAddresses struct {
	Adjudicator, AssetHolder, App common.Address
}

func deployContracts(
	ctx context.Context,
	nodeURL string,
	chainID *big.Int,
	deploymentKey *ecdsa.PrivateKey,
) (ContractAddresses, error) {
	c, err := NewEthClient(ctx, nodeURL, deploymentKey, chainID)
	if err != nil {
		return ContractAddresses{}, errors.WithMessage(err, "creating ethereum client")
	}

	// Deploy adjudicator.
	adj, txAdj, err := c.DeployAdjudicator(ctx)
	if err != nil {
		return ContractAddresses{}, errors.WithMessage(err, "deploying adjudicator")
	}

	// Deploy app.
	appAddr, txApp, err := c.DeployApp(ctx, adj)
	if err != nil {
		return ContractAddresses{}, errors.WithMessage(err, "deploying CollateralApp")
	}

	// Register app.
	swapApp := app.NewCredentialSwapApp(wallet.AsWalletAddr(appAddr))
	channel.RegisterApp(swapApp)

	// Deploy asset holder.
	assetHolderAddr, txAss, err := c.DeployAssetHolderETH(ctx, adj, appAddr)
	if err != nil {
		return ContractAddresses{}, errors.WithMessage(err, "deploying CollateralAssetHolderETH")
	}

	err = c.WaitDeployment(ctx, txAdj, txApp, txAss)
	if err != nil {
		return ContractAddresses{}, errors.WithMessage(err, "waiting for contract deployment")
	}

	return ContractAddresses{
		Adjudicator: adj,
		AssetHolder: assetHolderAddr,
		App:         appAddr,
	}, nil
}
