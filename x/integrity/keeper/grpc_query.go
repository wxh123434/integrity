package keeper

import (
	"github.com/SBC/integrity/x/integrity/types"
)

var _ types.QueryServer = Keeper{}
