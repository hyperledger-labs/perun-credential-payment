package abi

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Arguments = abi.Arguments

// Unpack unpacks `b` into `dst` according to `args`.
func Unpack(b []byte, dst interface{}, args abi.Arguments) error {
	argsI, err := args.Unpack(b)
	if err != nil {
		return fmt.Errorf("unpacking: %w", err)
	}

	return argumentsFixed{args}.Copy(dst, argsI)
}

// argumentsFixed is a representation of go-ethereum's `abi.Arguments` with a
// fixed Copy function.
type argumentsFixed struct {
	Arguments
}

// Copy writes `values` into `dst`.
func (a argumentsFixed) Copy(dst interface{}, values []interface{}) error {
	//TODO add comment
	prepareDst := func(dst interface{}) interface{} {
		val := reflect.ValueOf(dst).Elem()
		if len(a.Arguments) <= 1 && val.Kind() == reflect.Struct {
			return &struct{ Dst interface{} }{dst}
		}
		return dst
	}

	dstFixed := prepareDst(dst)
	err := a.Arguments.Copy(dstFixed, values)
	if err != nil {
		return fmt.Errorf("copying: %w", err)
	}

	return nil
}
