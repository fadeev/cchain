package keeper

import (
	"github.com/fadeev/cchain/x/cchain/types"
)

var _ types.QueryServer = Keeper{}
