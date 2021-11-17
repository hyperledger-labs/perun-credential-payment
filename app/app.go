// Copyright 2021 PolyCrypt GmbH, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/perun-network/verifiable-credential-payment/app/data"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
)

const (
	AssetIdx = 0
)

var (
	ErrInvalidInitData     = errors.New("invalid init data")
	ErrInvalidNextData     = errors.New("invalid next data")
	ErrExpectedOffer       = errors.New("expected offer")
	ErrUnequalAllocation   = errors.New("unequal allocation")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidSigner       = errors.New("invalid signer")
)

// CredentialSwapApp is a channel app for atomically trading a credential against a payment.
type CredentialSwapApp struct {
	Addr wallet.Address
}

func NewCredentialSwapApp(addr wallet.Address) *CredentialSwapApp {
	return &CredentialSwapApp{
		Addr: addr,
	}
}

func (a *CredentialSwapApp) InitData() channel.Data {
	return &data.DefaultData{}
}

// Def returns the app address.
func (a *CredentialSwapApp) Def() wallet.Address {
	return a.Addr
}

// DecodeData decodes the channel data.
func (a *CredentialSwapApp) DecodeData(r io.Reader) (channel.Data, error) {
	return data.Decode(r)
}

// ValidInit checks that the initial state is valid.
func (a *CredentialSwapApp) ValidInit(p *channel.Params, s *channel.State) error {
	_, ok := s.Data.(*data.DefaultData)
	if !ok {
		return ErrInvalidInitData
	}
	return nil
}

// ValidTransition is called whenever the channel state transitions.
func (a *CredentialSwapApp) ValidTransition(params *channel.Params, cur, next *channel.State, actorIdx channel.Index) error {
	// We require that there is only a single constant asset.
	if err := assertSingleConstantAsset(cur, next); err != nil {
		return err
	}

	offer, ok := cur.Data.(*data.Offer)
	// If we are not in offer mode, we require that the allocation did not change and return.
	if !ok {
		if !cur.Balances.Equal(next.Balances) {
			return fmt.Errorf("unequal balances")
		}
		return nil
	}

	// We are dealing with an offer.

	// Verify signature.
	{
		// Decode next state.
		cert, ok := next.Data.(*data.Cert)
		if !ok {
			return ErrInvalidNextData
		}

		err := VerifySig(cert.Signature, offer.DataHash, offer.Issuer)
		if err != nil {
			return fmt.Errorf("verifying signature: %w", err)
		}
	}

	// Verify balances.

	// Verify buyer balance.
	{
		expectedBal := new(big.Int).Sub(cur.Balances[AssetIdx][offer.Buyer], offer.Price)
		if next.Balances[AssetIdx][offer.Buyer].Cmp(expectedBal) != 0 {
			return fmt.Errorf("wrong balance: buyer")
		}
	}

	// Verify seller balance.
	{
		expectedBal := new(big.Int).Add(cur.Balances[AssetIdx][actorIdx], offer.Price)
		if next.Balances[AssetIdx][actorIdx].Cmp(expectedBal) != 0 {
			return fmt.Errorf("wrong balance: seller")
		}
	}

	return nil
}

func assertSingleConstantAsset(cur, next *channel.State) error {
	const numAssets = 1
	if len(cur.Allocation.Assets) != numAssets {
		return fmt.Errorf("wrong number of assets: current state")
	} else if len(next.Allocation.Assets) != numAssets {
		return fmt.Errorf("wrong number of assets: next state")
	} else if err := channel.AssetsAssertEqual(cur.Assets, next.Assets); err != nil {
		return fmt.Errorf("asset not equal: %w", err)
	}
	return nil
}

type (
	Hash      = [data.HashLen]byte
	Signature = []byte
)

func ComputeDocumentHash(doc []byte) Hash {
	return crypto.Keccak256Hash(doc)
}

type Credential struct {
	Document  []byte
	Signature []byte
}

func (c *Credential) String() string {
	return fmt.Sprintf("Document: \"%s\" Signature: \"%x\"", c.Document, c.Signature)
}
