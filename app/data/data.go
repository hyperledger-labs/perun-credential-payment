package data

import (
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	appabi "github.com/perun-network/perun-credential-payment/app/abi"
	"perun.network/go-perun/channel"
)

const (
	HashLen = 32
	SigLen  = 65
)

type Mode = uint8

const (
	defaultMode Mode = iota
	offerMode
	certMode
)

// DefaultData represents the default state.
type DefaultData struct{}

// Encode encodes the data onto an io.Writer.
func (d *DefaultData) Encode(w io.Writer) error {
	f := &dataFrame{
		Mode: defaultMode,
		Data: []byte{},
	}
	return f.Encode(w)
}

// Clone returns a deep copy of the data.
func (d *DefaultData) Clone() channel.Data {
	_d := *d
	return &_d
}

// Offer represents an offer.
type Offer struct {
	Issuer   common.Address
	DataHash [HashLen]byte
	Price    *big.Int
	Buyer    uint16
}

func (a Offer) Equal(b *Offer) bool {
	return a.Issuer == b.Issuer &&
		a.DataHash == b.DataHash &&
		a.Price.Cmp(b.Price) == 0 &&
		a.Buyer == b.Buyer
}

var offerType = func() abi.Type {
	t, err := abi.NewType(
		"tuple",
		"offer",
		[]abi.ArgumentMarshaling{
			{Type: "address", Name: "issuer"},
			{Type: "bytes32", Name: "dataHash"},
			{Type: "uint256", Name: "price"},
			{Type: "uint16", Name: "buyer"},
		},
	)
	if err != nil {
		panic(err)
	}
	return t
}()

var offerArgs = appabi.Arguments{
	{Name: "offer", Type: offerType},
}

// Encode encodes app data onto an io.Writer.
func (d *Offer) Encode(w io.Writer) error {
	body, err := offerArgs.Pack(d)
	if err != nil {
		return err
	}

	f := &dataFrame{
		Mode: offerMode,
		Data: body,
	}
	return f.Encode(w)
}

func (d *Offer) Unmarshal(b []byte) error {
	return appabi.Unpack(b, d, offerArgs)
}

// Clone returns a deep copy of the app data.
func (d *Offer) Clone() channel.Data {
	_d := *d
	_d.Price = new(big.Int).Set(d.Price)
	return &_d
}

// Cert represents an offer response.
type Cert struct {
	Signature [SigLen]byte
}

// Encode encodes the data onto an io.Writer.
func (d *Cert) Encode(w io.Writer) error {
	f := &dataFrame{
		Mode: certMode,
		Data: d.Signature[:],
	}
	return f.Encode(w)
}

// Clone returns a deep copy of the app data.
func (d *Cert) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *Cert) Unmarshal(b []byte) error {
	if len(b) != len(d.Signature) {
		return fmt.Errorf("invalid signature length")
	}
	copy(d.Signature[:], b)
	return nil
}

func Decode(r io.Reader) (channel.Data, error) {
	var f dataFrame
	err := f.Decode(r)
	if err != nil {
		return nil, err
	}

	switch f.Mode {
	case defaultMode:
		return &DefaultData{}, nil
	case offerMode:
		var offer Offer
		return &offer, offer.Unmarshal(f.Data)
	case certMode:
		var cert Cert
		return &cert, cert.Unmarshal(f.Data)
	default:
		return nil, fmt.Errorf("unknown mode")
	}
}
