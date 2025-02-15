// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package math

import (
	"encoding/binary"
	"math/big"
	"math/bits"
	"reflect"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/hex"
)

const (
	// U64NumBytes is the number of bytes in a U64.
	U64NumBytes = 8
	// U64NumBits is the number of bits in a U64.
	U64NumBits = U64NumBytes * 8
)

//nolint:gochecknoglobals // stores the reflect type of U64.
var uint64T = reflect.TypeOf(U64(0))

//nolint:lll
type (
	// U64 represents a 64-bit unsigned integer that is both SSZ and JSON
	// marshallable. We marshal U64 as hex strings in JSON in order to keep the
	// execution client apis happy, and we marshal U64 as little-endian in SSZ
	// to be
	// compatible with the spec.
	U64 uint64

	// Gwei is a denomination of 1e9 Wei represented as a U64.
	Gwei = U64

	// Slot as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Slot = U64

	// CommitteeIndex as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	CommitteeIndex = U64

	// ValidatorIndex as per the Ethereum 2.0  Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	ValidatorIndex = U64

	// Epoch as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Epoch = U64
)

// -------------------------- SSZMarshallable --------------------------

// MarshalSSZTo serializes the U64 into a byte slice.
func (u U64) MarshalSSZTo(buf []byte) ([]byte, error) {
	binary.LittleEndian.PutUint64(buf, uint64(u))
	return buf, nil
}

// MarshalSSZ serializes the U64 into a byte slice.
func (u U64) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, U64NumBytes)
	if _, err := u.MarshalSSZTo(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// UnmarshalSSZ deserializes the U64 from a byte slice.
func (u *U64) UnmarshalSSZ(buf []byte) error {
	if len(buf) != U64NumBytes {
		return ErrUnexpectedInputLength(U64NumBytes, len(buf))
	}
	if u == nil {
		u = new(U64)
	}
	*u = U64(binary.LittleEndian.Uint64(buf))
	return nil
}

// SizeSSZ returns the size of the U64 in bytes.
func (u U64) SizeSSZ() int {
	return U64NumBytes
}

// HashTreeRoot computes the Merkle root of the U64 using SSZ hashing rules.
func (u U64) HashTreeRoot() ([32]byte, error) {
	buf := make([]byte, U64NumBytes)
	binary.LittleEndian.PutUint64(buf, uint64(u))
	var hashRoot [32]byte
	copy(hashRoot[:], buf)
	return hashRoot, nil
}

// -------------------------- JSONMarshallable -------------------------

// MarshalText implements encoding.TextMarshaler.
func (u U64) MarshalText() ([]byte, error) {
	return hex.MarshalText(u.Unwrap())
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *U64) UnmarshalJSON(input []byte) error {
	return hex.UnmarshalJSONText(input, u, uint64T)
}

// ---------------------------------- Hex ----------------------------------

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *U64) UnmarshalText(input []byte) error {
	dec, err := hex.UnmarshalUint64Text(input)
	if err != nil {
		return err
	}
	*u = U64(dec)
	return nil
}

// String returns the hex encoding of b.
func (u U64) String() hex.String {
	return hex.FromUint64(u.Unwrap())
}

// ----------------------- U64 Mathematical Methods -----------------------

// Unwrap returns a copy of the underlying uint64 value of U64.
func (u U64) Unwrap() uint64 {
	return uint64(u)
}

// UnwrapPtr returns a pointer to the underlying uint64 value of U64.
func (u U64) UnwrapPtr() *uint64 {
	return (*uint64)(&u)
}

// Get the power of 2 for given input, or the closest higher power of 2 if the
// input is not a power of 2. Commonly used for "how many nodes do I need for a
// bottom tree layer fitting x elements?"
// Example: 0->1, 1->1, 2->2, 3->4, 4->4, 5->8, 6->8, 7->8, 8->8, 9->16.
//
// https://github.com/ethereum/consensus-specs/blob/dev/ssz/merkle-proofs.md#helper-functions
//
//nolint:mnd,lll // powers of 2.
func (u U64) NextPowerOfTwo() U64 {
	if u == 0 {
		return 1
	}
	if u > 1<<63 {
		panic("Next power of 2 is 1 << 64.")
	}
	u--
	u |= u >> 1
	u |= u >> 2
	u |= u >> 4
	u |= u >> 8
	u |= u >> 16
	u |= u >> 32
	u++
	return u
}

// Get the power of 2 for given input, or the closest lower power of 2 if the
// input is not a power of 2. The zero case is a placeholder and not used for
// math with generalized indices. Commonly used for "what power of two makes up
// the root bit of the generalized index?"
// Example: 0->1, 1->1, 2->2, 3->2, 4->4, 5->4, 6->4, 7->4, 8->8, 9->8.
//
// https://github.com/ethereum/consensus-specs/blob/dev/ssz/merkle-proofs.md#helper-functions
//
//nolint:mnd,lll // From Ethereum 2.0 spec.
func (u U64) PrevPowerOfTwo() U64 {
	if u == 0 {
		return 1
	}
	u |= u >> 1
	u |= u >> 2
	u |= u >> 4
	u |= u >> 8
	u |= u >> 16
	u |= u >> 32
	return u - (u >> 1)
}

// ILog2Ceil returns the ceiling of the base 2 logarithm of the U64.
func (u U64) ILog2Ceil() uint8 {
	// Log2(0) is undefined, should we panic?
	if u == 0 {
		return 0
	}
	//#nosec:G701 // we handle the case of u == 0 above, so this is safe.
	return uint8(bits.Len64(uint64(u - 1)))
}

// ILog2Floor returns the floor of the base 2 logarithm of the U64.
func (u U64) ILog2Floor() uint8 {
	// Log2(0) is undefined, should we panic?
	if u == 0 {
		return 0
	}
	//#nosec:G701 // we handle the case of u == 0 above, so this is safe.
	return uint8(bits.Len64(uint64(u))) - 1
}

// ---------------------------- Gwei Methods ----------------------------

// GweiToWei returns the value of Wei in Gwei.
func GweiFromWei(i *big.Int) Gwei {
	intToGwei := big.NewInt(0).SetUint64(constants.GweiPerWei)
	i.Div(i, intToGwei)
	return Gwei(i.Uint64())
}

// ToWei converts a value from Gwei to Wei.
func (u Gwei) ToWei() *big.Int {
	gweiAmount := big.NewInt(0).SetUint64(u.Unwrap())
	intToGwei := big.NewInt(0).SetUint64(constants.GweiPerWei)
	return gweiAmount.Mul(gweiAmount, intToGwei)
}
