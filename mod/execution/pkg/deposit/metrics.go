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

package deposit

import (
	"strconv"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// depositMetrics is a struct that contains metrics for the deposit service.
type depositMetrics struct {
	// sink is the telemetry sink.
	sink TelemetrySink
}

// newDepositMetrics creates a new instance of the depositMetrics struct.
func newDepositMetrics(sink TelemetrySink) *depositMetrics {
	return &depositMetrics{
		sink: sink,
	}
}

// markFailedToGetBlockLogs increments the counter for failed to get block logs.
func (m *depositMetrics) markFailedToGetBlockLogs(blockNum math.U64) {
	m.sink.IncrementCounter(
		"beacon_kit.execution.deposit.failed_to_get_block_logs",
		"block_num",
		strconv.FormatUint(uint64(blockNum), 10),
	)
}
