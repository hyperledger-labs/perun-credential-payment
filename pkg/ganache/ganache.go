package ganache

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type GanacheConfig struct {
	Cmd           string
	Host          string
	Port          uint
	BlockTime     time.Duration
	Funding       []KeyWithBalance
	StartupTime   time.Duration
	ChainID       *big.Int
	PrintToStdOut bool
}

type Ganache struct {
	Accounts []Account
	Cmd      *exec.Cmd
}

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	Amount     *big.Int
}

type KeyWithBalance struct {
	PrivateKey string
	BalanceEth uint
}

func StartGanacheWithPrefundedAccounts(cfg GanacheConfig) (ganache *Ganache, err error) {
	// Create accounts
	accounts := make([]Account, len(cfg.Funding))
	for i, funding := range cfg.Funding {
		accountKey, err := crypto.HexToECDSA(funding.PrivateKey[2:])
		if err != nil {
			return nil, errors.WithMessage(err, "parsing private key")
		}
		accounts[i] = Account{PrivateKey: accountKey, Amount: ethToWei(big.NewFloat(float64(funding.BalanceEth)))}
	}

	// Build ganache command line arguments
	var ganacheArgs []string
	ganacheArgs = append(ganacheArgs, "ganache-cli", "--host", cfg.Host, "--port", fmt.Sprint(cfg.Port))
	for _, a := range accounts {
		key := hexutil.Encode(crypto.FromECDSA(a.PrivateKey))
		ganacheArgs = append(ganacheArgs, "--account", fmt.Sprintf("%v,%v", key, a.Amount))
	}
	ganacheArgs = append(ganacheArgs, fmt.Sprintf("--blockTime=%v", int(cfg.BlockTime.Seconds())))
	ganacheArgs = append(ganacheArgs, fmt.Sprintf("--chainId=%d", cfg.ChainID.Uint64()))

	// Start command
	ganacheCmdTokens := strings.Split(cfg.Cmd, " ")
	cmdName := ganacheCmdTokens[0]
	var cmdArgs []string
	cmdArgs = append(cmdArgs, ganacheCmdTokens[1:]...)
	cmdArgs = append(cmdArgs, ganacheArgs...)
	cmd := exec.Command(cmdName, cmdArgs...)

	// This is needed for correctly shutting down ganache-cli.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if cfg.PrintToStdOut {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Panic(err)
		}
		go func() {
			rd := bufio.NewReader(stdout)
			for {
				str, err := rd.ReadString('\n')
				if err != nil {
					log.Print("Failed to read ganache output:", err)
					return
				}
				log.Print(str)
			}
		}()
	}

	if err := cmd.Start(); err != nil {
		return nil, errors.WithMessage(err, "starting ganache")
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- cmd.Wait()
	}()
	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(cfg.StartupTime):
	}
	return &Ganache{accounts, cmd}, nil
}

func (g *Ganache) Shutdown() error {
	// Running Process.Kill() does not kill child processes.
	// The below kills the process group referenced by the negative process ID
	// and therefore correctly shuts down ganache-cli.
	// May only work on unix-like systems.
	return syscall.Kill(-g.Cmd.Process.Pid, syscall.SIGKILL)
}

func ethToWei(eth *big.Float) (wei *big.Int) {
	var weiPerEth = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	var weiPerEthFloat = new(big.Float).SetInt(weiPerEth)
	wei, _ = new(big.Float).Mul(eth, weiPerEthFloat).Int(nil)
	return
}

func (a *Account) Address() common.Address {
	return crypto.PubkeyToAddress(a.PrivateKey.PublicKey)
}

func (cfg GanacheConfig) NodeURL() string {
	return fmt.Sprintf("ws://%s:%d", cfg.Host, cfg.Port)
}
