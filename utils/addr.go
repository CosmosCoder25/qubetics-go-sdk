package utils

import (
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	qubetics "github.com/qubetics/qubetics-blockchain/v2/types"
)

// MustAccAddrFromBech32 converts a Bech32-encoded string to a cosmossdk.AccAddress,
// panicking if there is an error during the conversion.
func MustAccAddrFromBech32(v string) cosmossdk.AccAddress {
	// If the input string is empty, return nil
	if v == "" {
		return nil
	}

	// Attempt to convert the Bech32 string to a cosmossdk.AccAddress
	addr, err := cosmossdk.AccAddressFromBech32(v)

	// If there is an error during the conversion, panic
	if err != nil {
		panic(err)
	}

	// Return the converted address
	return addr
}

// MustNodeAddrFromBech32 converts a Bech32-encoded string to a qubetics.NodeAddress,
// panicking if there is an error during the conversion.
func MustNodeAddrFromBech32(v string) qubetics.NodeAddress {
	// If the input string is empty, return nil
	if v == "" {
		return nil
	}

	// Attempt to convert the Bech32 string to a qubetics.NodeAddress
	addr, err := qubetics.NodeAddressFromBech32(v)

	// If there is an error during the conversion, panic
	if err != nil {
		panic(err)
	}

	// Return the converted address
	return addr
}
