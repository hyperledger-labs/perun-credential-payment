package data

import (
	"io"

	"github.com/ethereum/go-ethereum/accounts/abi"
	appabi "github.com/perun-network/perun-credential-payment/app/abi"
)

type dataFrame struct {
	Mode Mode
	Data []byte
}

var frameType = func() abi.Type {
	t, err := abi.NewType(
		"tuple",
		"frame",
		[]abi.ArgumentMarshaling{
			{Type: "uint8", Name: "mode"},
			{Type: "bytes", Name: "data"},
		},
	)
	if err != nil {
		panic(err)
	}
	return t
}()

var frameArgs = appabi.Arguments{
	{Name: "frame", Type: frameType},
}

func (f *dataFrame) Encode(w io.Writer) error {
	b, err := frameArgs.Pack(f)
	if err != nil {
		return err
	}

	return writeBytesWithLengthUint16(w, b)
}

func (f *dataFrame) Decode(r io.Reader) error {
	b, err := readBytesWithLengthUint16(r)
	if err != nil {
		return err
	}

	return appabi.Unpack(b, f, frameArgs)
}
