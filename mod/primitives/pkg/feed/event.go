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

package feed

import (
	"context"
)

// Event represents a generic event in the beacon chain.
type Event[DataT any] struct {
	// ctx is the context associated with the event.
	ctx context.Context
	// name is the name of the event.
	name string
	// event is the actual beacon event.
	data DataT
}

// NewEvent creates a new Event with the given context and beacon event.
func NewEvent[
	DataT any,
](ctx context.Context, name string, data DataT) *Event[DataT] {
	return &Event[DataT]{
		ctx:  ctx,
		name: name,
		data: data,
	}
}

// Name returns the name of the event.
func (e Event[DataT]) Name() string {
	return e.name
}

// Context returns the context associated with the event.
func (e Event[DataT]) Context() context.Context {
	return e.ctx
}

// Event returns the beacon event.
func (e Event[DataT]) Data() DataT {
	return e.data
}

// Is returns true if the event has the given name.
func (e Event[DataT]) Is(name string) bool {
	return e.name == name
}
