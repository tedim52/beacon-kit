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

package components

import (
	"os"
	"path/filepath"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

//nolint:gochecknoglobals // todo:fix from sdk.
var DefaultNodeHome string

//nolint:gochecknoinits // annoying from sdk.
func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	DefaultNodeHome = filepath.Join(userHomeDir, ".beacond")
}

//
//nolint:lll // link.
const TermsOfServiceURL = "https://github.com/berachain/beacon-kit/blob/main/TERMS_OF_SERVICE.md"

// ProvideClientContext returns a new client context with the given options.
func ProvideClientContext(
	appCodec codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
	txConfig client.TxConfig,
	addressCodec address.Codec,
	validatorAddressCodec address.ValidatorAddressCodec,
	consensusAddressCodec address.ConsensusAddressCodec,
) (client.Context, error) {
	var err error

	clientCtx := client.Context{}.
		WithCodec(appCodec).
		WithInterfaceRegistry(interfaceRegistry).
		WithInput(os.Stdin).
		WithAddressCodec(addressCodec).
		WithValidatorAddressCodec(validatorAddressCodec).
		WithConsensusAddressCodec(consensusAddressCodec).
		WithHomeDir(DefaultNodeHome).
		WithViper("") // uses by default the binary name as prefix

	// Read the config to overwrite the default values with the values from the
	// config file
	customClientTemplate, customClientConfig := InitClientConfig()
	clientCtx, err = config.ReadDefaultValuesFromDefaultClientConfig(
		clientCtx,
		customClientTemplate,
		customClientConfig,
	)
	if err != nil {
		return clientCtx, err
	}

	clientCtx = clientCtx.WithTxConfig(txConfig)

	return clientCtx, nil
}

// InitClientConfig sets up the default client configuration, allowing for
// overrides.
func InitClientConfig() (string, interface{}) {
	clientCfg := config.DefaultConfig()
	clientCfg.KeyringBackend = keyring.BackendTest
	return config.DefaultClientConfigTemplate, clientCfg
}
