package ethereum

import (
	"math/big"

	"github.com/alpe/community-bifrost/pkg/common"
)

func (t Transaction) ValueToStellar() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.StellarAmountPrecision)
}
