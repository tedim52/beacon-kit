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

package types_test

import (
	"encoding/json"
	"testing"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/version"
	"github.com/stretchr/testify/require"
)

func generateExecutableDataDeneb() *types.ExecutableDataDeneb {
	return &types.ExecutableDataDeneb{
		ParentHash:    common.ExecutionHash{},
		FeeRecipient:  common.ExecutionAddress{},
		StateRoot:     bytes.B32{},
		ReceiptsRoot:  bytes.B32{},
		LogsBloom:     make([]byte, 256),
		Random:        bytes.B32{},
		Number:        math.U64(0),
		GasLimit:      math.U64(0),
		GasUsed:       math.U64(0),
		Timestamp:     math.U64(0),
		ExtraData:     []byte{},
		BaseFeePerGas: math.Wei{},
		BlockHash:     common.ExecutionHash{},
		Transactions:  [][]byte{},
		Withdrawals:   []*engineprimitives.Withdrawal{},
		BlobGasUsed:   math.U64(0),
		ExcessBlobGas: math.U64(0),
	}
}
func TestExecutableDataDeneb_Serialization(t *testing.T) {
	original := generateExecutableDataDeneb()

	data, err := original.MarshalSSZ()
	require.NoError(t, err)
	require.NotNil(t, data)

	var unmarshalled types.ExecutableDataDeneb
	err = unmarshalled.UnmarshalSSZ(data)
	require.NoError(t, err)

	require.Equal(t, original, &unmarshalled)
}

func TestExecutableDataDeneb_SizeSSZ(t *testing.T) {
	payload := generateExecutableDataDeneb()
	size := payload.SizeSSZ()
	require.Equal(t, 528, size)
}

func TestExecutableDataDeneb_HashTreeRoot(t *testing.T) {
	payload := generateExecutableDataDeneb()
	_, err := payload.HashTreeRoot()
	require.NoError(t, err)
}

func TestExecutableDataDeneb_GetTree(t *testing.T) {
	payload := generateExecutableDataDeneb()
	tree, err := payload.GetTree()
	require.NoError(t, err)
	require.NotNil(t, tree)
}

func TestExecutableDataDeneb_Getters(t *testing.T) {
	payload := generateExecutableDataDeneb()

	require.Equal(t, common.ExecutionHash{}, payload.GetParentHash())
	require.Equal(t, common.ExecutionAddress{}, payload.GetFeeRecipient())
	require.Equal(t, bytes.B32{}, payload.GetStateRoot())
	require.Equal(t, bytes.B32{}, payload.GetReceiptsRoot())
	require.Equal(t, make([]byte, 256), payload.GetLogsBloom())
	require.Equal(t, bytes.B32{}, payload.GetPrevRandao())
	require.Equal(t, math.U64(0), payload.GetNumber())
	require.Equal(t, math.U64(0), payload.GetGasLimit())
	require.Equal(t, math.U64(0), payload.GetGasUsed())
	require.Equal(t, math.U64(0), payload.GetTimestamp())
	require.Equal(t, []byte{}, payload.GetExtraData())
	require.Equal(t, math.Wei{}, payload.GetBaseFeePerGas())
	require.Equal(t, common.ExecutionHash{}, payload.GetBlockHash())
	require.Equal(t, [][]byte{}, payload.GetTransactions())
	require.Equal(t, []*engineprimitives.Withdrawal{}, payload.GetWithdrawals())
	require.Equal(t, math.U64(0), payload.GetBlobGasUsed())
	require.Equal(t, math.U64(0), payload.GetExcessBlobGas())
}

func TestExecutableDataDeneb_MarshalJSON(t *testing.T) {
	payload := generateExecutableDataDeneb()

	data, err := payload.MarshalJSON()
	require.NoError(t, err)
	require.NotNil(t, data)

	var unmarshalled types.ExecutableDataDeneb
	err = unmarshalled.UnmarshalJSON(data)
	require.NoError(t, err)
	require.Equal(t, payload, &unmarshalled)
}

func TestExecutableDataDeneb_IsNil(t *testing.T) {
	var payload *types.ExecutableDataDeneb
	require.True(t, payload.IsNil())

	payload = generateExecutableDataDeneb()
	require.False(t, payload.IsNil())
}

func TestExecutableDataDeneb_IsBlinded(t *testing.T) {
	payload := generateExecutableDataDeneb()
	require.False(t, payload.IsBlinded())
}

func TestExecutableDataDeneb_Version(t *testing.T) {
	payload := generateExecutableDataDeneb()
	require.Equal(t, version.Deneb, payload.Version())
}

func TestExecutionPayload_Empty(t *testing.T) {
	payload := new(types.ExecutionPayload)
	emptyPayload := payload.Empty(version.Deneb)

	require.NotNil(t, emptyPayload)
	require.Equal(t, version.Deneb, emptyPayload.Version())
}

func TestExecutionPayload_ToHeader(t *testing.T) {
	payload := types.ExecutionPayload{
		InnerExecutionPayload: &types.ExecutableDataDeneb{
			ParentHash:    common.ExecutionHash{},
			FeeRecipient:  common.ExecutionAddress{},
			StateRoot:     bytes.B32{},
			ReceiptsRoot:  bytes.B32{},
			LogsBloom:     make([]byte, 256),
			Random:        bytes.B32{},
			Number:        math.U64(0),
			GasLimit:      math.U64(0),
			GasUsed:       math.U64(0),
			Timestamp:     math.U64(0),
			ExtraData:     []byte{},
			BaseFeePerGas: math.Wei{},
			BlockHash:     common.ExecutionHash{},
			Transactions:  [][]byte{},
			Withdrawals:   []*engineprimitives.Withdrawal{},
			BlobGasUsed:   math.U64(0),
			ExcessBlobGas: math.U64(0),
		},
	}

	header, err := payload.ToHeader()
	require.NoError(t, err)
	require.NotNil(t, header)

	require.Equal(t, payload.GetParentHash(), header.GetParentHash())
	require.Equal(t, payload.GetFeeRecipient(), header.GetFeeRecipient())
	require.Equal(t, payload.GetStateRoot(), header.GetStateRoot())
	require.Equal(t, payload.GetReceiptsRoot(), header.GetReceiptsRoot())
	require.Equal(t, payload.GetLogsBloom(), header.GetLogsBloom())
	require.Equal(t, payload.GetPrevRandao(), header.GetPrevRandao())
	require.Equal(t, payload.GetNumber(), header.GetNumber())
	require.Equal(t, payload.GetGasLimit(), header.GetGasLimit())
	require.Equal(t, payload.GetGasUsed(), header.GetGasUsed())
	require.Equal(t, payload.GetTimestamp(), header.GetTimestamp())
	require.Equal(t, payload.GetExtraData(), header.GetExtraData())
	require.Equal(t, payload.GetBaseFeePerGas(), header.GetBaseFeePerGas())
	require.Equal(t, payload.GetBlockHash(), header.GetBlockHash())
	require.Equal(t, payload.GetBlobGasUsed(), header.GetBlobGasUsed())
	require.Equal(t, payload.GetExcessBlobGas(), header.GetExcessBlobGas())
}

//nolint:lll
func TestExecutableDataDeneb_UnmarshalJSON_Error(t *testing.T) {
	original := generateExecutableDataDeneb()
	validJSON, err := original.MarshalJSON()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		removeField   string
		expectedError string
	}{
		{
			name:          "missing required field 'parentHash'",
			removeField:   "parentHash",
			expectedError: "missing required field 'parentHash' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'feeRecipient'",
			removeField:   "feeRecipient",
			expectedError: "missing required field 'feeRecipient' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'stateRoot'",
			removeField:   "stateRoot",
			expectedError: "missing required field 'stateRoot' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'receiptsRoot'",
			removeField:   "receiptsRoot",
			expectedError: "missing required field 'receiptsRoot' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'logsBloom'",
			removeField:   "logsBloom",
			expectedError: "missing required field 'logsBloom' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'prevRandao'",
			removeField:   "prevRandao",
			expectedError: "missing required field 'prevRandao' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'blockNumber'",
			removeField:   "blockNumber",
			expectedError: "missing required field 'blockNumber' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'gasLimit'",
			removeField:   "gasLimit",
			expectedError: "missing required field 'gasLimit' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'gasUsed'",
			removeField:   "gasUsed",
			expectedError: "missing required field 'gasUsed' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'timestamp'",
			removeField:   "timestamp",
			expectedError: "missing required field 'timestamp' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'extraData'",
			removeField:   "extraData",
			expectedError: "missing required field 'extraData' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'baseFeePerGas'",
			removeField:   "baseFeePerGas",
			expectedError: "missing required field 'baseFeePerGas' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'blockHash'",
			removeField:   "blockHash",
			expectedError: "missing required field 'blockHash' for ExecutableDataDeneb",
		},
		{
			name:          "missing required field 'transactions'",
			removeField:   "transactions",
			expectedError: "missing required field 'transactions' for ExecutableDataDeneb",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var payload types.ExecutableDataDeneb
			var jsonMap map[string]interface{}

			errUnmarshal := json.Unmarshal(validJSON, &jsonMap)
			require.NoError(t, errUnmarshal)

			delete(jsonMap, tc.removeField)

			malformedJSON, errMarshal := json.Marshal(jsonMap)
			require.NoError(t, errMarshal)

			err = payload.UnmarshalJSON(malformedJSON)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedError)
		})
	}
}
