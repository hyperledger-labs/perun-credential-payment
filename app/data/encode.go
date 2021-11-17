package data

import (
	"fmt"
	"io"

	pio "perun.network/go-perun/pkg/io"
)

func writeBytesWithLengthUint16(w io.Writer, b []byte) error {
	l := uint16safe(len(b))
	return pio.Encode(w, l, b)
}

func uint16safe(i int) uint16 {
	const uint16max = 1<<16 - 1 // 2 ^ 16 - 1
	if i > uint16max {
		panic("out of bounds")
	}
	return uint16(i)
}

func readBytesWithLengthUint16(r io.Reader) ([]byte, error) {
	var l uint16
	err := pio.Decode(r, &l)
	if err != nil {
		return nil, fmt.Errorf("reading length: %w", err)
	}

	b := make([]byte, l)
	err = pio.Decode(r, &b)
	return b, err
}
