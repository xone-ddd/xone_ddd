package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hello-xone/xone-chain/x/xone/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
