package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	Address = createType("address")
	Bytes32 = createType("bytes32")
	Bytes   = createType("bytes")
	Uint8   = createType("uint8")
	Uint16  = createType("uint16")
	Uint256 = createType("uint256")
)

func createType(name string) abi.Type {
	t, err := abi.NewType(name, "", nil)
	if err != nil {
		panic(err)
	}
	return t
}
