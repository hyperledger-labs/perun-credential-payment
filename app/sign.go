package app

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/perun-network/perun-credential-payment/app/data"
	"perun.network/go-perun/backend/ethereum/wallet/simple"
)

const (
	sigVIndex    = 64
	sigVMagicNum = 27
)

func SignHash(acc *simple.Account, h [data.HashLen]byte) ([data.SigLen]byte, error) {
	sig, err := acc.SignHash(h[:])
	if err != nil {
		return [data.SigLen]byte{}, fmt.Errorf("signing hash: %w", err)
	}
	sig[sigVIndex] += sigVMagicNum

	var sigFixedLen [data.SigLen]byte
	copy(sigFixedLen[:], sig)
	return sigFixedLen, nil
}

func VerifySig(sig [data.SigLen]byte, h [data.HashLen]byte, addr common.Address) error {
	sig[sigVIndex] -= sigVMagicNum
	realSigner, err := crypto.Ecrecover(h[:], sig[:])
	if err != nil {
		return fmt.Errorf("failed to recover signer: %w", err)
	}

	signerPubKey, err := crypto.UnmarshalPubkey(realSigner)
	if err != nil {
		return fmt.Errorf("unmarshalling public key: %w", err)
	}

	signerAddr := crypto.PubkeyToAddress(*signerPubKey)
	if !bytes.Equal(signerAddr[:], addr[:]) {
		return ErrInvalidSigner
	}

	return nil
}
