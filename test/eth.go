package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/perun-network/verifiable-credential-payment/app"
	"github.com/pkg/errors"

	"perun.network/go-perun/backend/ethereum/bindings/adjudicator"
	"perun.network/go-perun/backend/ethereum/bindings/assetholdereth"
)

type EthClient struct {
	*ethclient.Client
	key     *ecdsa.PrivateKey
	chainID *big.Int
	nonce   uint64
}

func NewEthClient(ctx context.Context, nodeURL string, key *ecdsa.PrivateKey, chainID *big.Int) (*EthClient, error) {
	client, err := ethclient.DialContext(ctx, nodeURL)
	if err != nil {
		return nil, fmt.Errorf("dialing: %w", err)
	}

	addr := crypto.PubkeyToAddress(key.PublicKey)
	nonce, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return nil, fmt.Errorf("getting nonce: %w", err)
	}

	return &EthClient{
		Client:  client,
		key:     key,
		chainID: chainID,
		nonce:   nonce,
	}, nil
}

func (c *EthClient) DeployAdjudicator(ctx context.Context) (addr common.Address, tx *types.Transaction, err error) {
	return c.deployContract(ctx, func(to *bind.TransactOpts, c *ethclient.Client) (addr common.Address, tx *types.Transaction, err error) {
		addr, tx, _, err = adjudicator.DeployAdjudicator(to, c)
		return
	}, false)
}

func (c *EthClient) DeployApp(ctx context.Context, adjudicatorAddr common.Address) (addr common.Address, tx *types.Transaction, err error) {
	return c.deployContract(ctx, func(to *bind.TransactOpts, c *ethclient.Client) (addr common.Address, tx *types.Transaction, err error) {
		addr, tx, _, err = app.DeployCredentialSwap(to, c)
		return
	}, false)
}

func (c *EthClient) DeployAssetHolderETH(ctx context.Context, adjudicatorAddr common.Address, appAddr common.Address) (addr common.Address, tx *types.Transaction, err error) {
	return c.deployContract(ctx, func(to *bind.TransactOpts, c *ethclient.Client) (addr common.Address, tx *types.Transaction, err error) {
		addr, tx, _, err = assetholdereth.DeployAssetHolderETH(to, c, adjudicatorAddr)
		return
	}, false)
}

func (c *EthClient) deployContract(
	ctx context.Context,
	deployContract func(*bind.TransactOpts, *ethclient.Client) (common.Address, *types.Transaction, error),
	waitConfirmation bool,
) (common.Address, *types.Transaction, error) {
	tr, err := c.newTransactor(ctx)
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, tx, err := deployContract(tr, c.Client)
	if err != nil {
		return common.Address{}, nil, errors.WithMessage(err, "sending deployment transaction")
	}

	if waitConfirmation {
		addr, err = bind.WaitDeployed(ctx, c.Client, tx)
		if err != nil {
			return common.Address{}, nil, errors.WithMessage(err, "waiting for the deployment transaction to be mined")
		}
	}
	return addr, tx, nil
}

func (c *EthClient) WaitDeployment(ctx context.Context, txs ...*types.Transaction) (err error) {
	for _, tx := range txs {
		_, err = bind.WaitDeployed(ctx, c.Client, tx)
		if err != nil {
			return errors.WithMessagef(err, "waiting for deployment: %v", tx)
		}
	}
	return nil
}

func (c *EthClient) newTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	tr, err := bind.NewKeyedTransactorWithChainID(c.key, c.chainID)
	if err != nil {
		return nil, err
	}
	tr.Context = ctx
	tr.Nonce = new(big.Int).SetUint64(c.nonce)
	c.nonce++
	// tr.GasPrice = big.NewInt(20000000000)
	// tr.GasLimit = 6721975
	return tr, nil
}

func (c *EthClient) AccountBalance(a common.Address) (b *big.Int, err error) {
	return c.BalanceAt(context.Background(), a, nil)
}

func WeiToEth(weiAmount *big.Int) (ethAmount *big.Float) {
	weiPerEth := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	weiPerEthFloat := new(big.Float).SetInt(weiPerEth)
	weiAmountFloat := new(big.Float).SetInt(weiAmount)
	return new(big.Float).Quo(weiAmountFloat, weiPerEthFloat)
}

func EthToWei(ethAmount *big.Float) (weiAmount *big.Int) {
	weiPerEth := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	weiPerEthFloat := new(big.Float).SetInt(weiPerEth)
	weiAmountFloat := new(big.Float).Mul(ethAmount, weiPerEthFloat)
	weiAmount, _ = weiAmountFloat.Int(nil)
	return weiAmount
}
