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

package bytes

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/hex"
)

// B8 represents a 4-byte array.
type B8 [8]byte

// UnmarshalJSON implements the json.Unmarshaler interface for B8.
func (h *B8) UnmarshalJSON(input []byte) error {
	return unmarshalJSONHelper(h[:], input)
}

// ToBytes8 is a utility function that transforms a byte slice into a fixed
// 8-byte array. If the input exceeds 4 bytes, it gets truncated.
func ToBytes8(input []byte) B8 {
	//nolint:mnd // 8 bytes.
	return [8]byte(ExtendToSize(input, 8))
}

// String returns the hex string representation of B8.
func (h B8) String() string {
	return hex.FromBytes(h[:]).Unwrap()
}

// MarshalText implements the encoding.TextMarshaler interface for B8.
func (h B8) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for B8.
func (h *B8) UnmarshalText(text []byte) error {
	return UnmarshalTextHelper(h[:], text)
}
