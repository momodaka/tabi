// Copyright 2024 Tabi Foundation
// This file is part of the Tabi Network packages.
//
// Tabi is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Tabi packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tabi/tabi/v1/x/evm/types"
)

// GetParams returns the total set of evm parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixParams)
	if len(bz) == 0 {
		return k.GetLegacyParams(ctx)
	}
	k.cdc.MustUnmarshal(bz, &params)
	return
}

// SetParams sets the EVM params each in their individual key for better get performance
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(types.KeyPrefixParams, bz)
	return nil
}

// GetLegacyParams returns param set for version before migrate
func (k Keeper) GetLegacyParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.ss.GetParamSetIfExists(ctx, &params)
	return params
}