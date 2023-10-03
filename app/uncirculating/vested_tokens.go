package uncirculating

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stakedotlink/token-supply-api/abi"
	"github.com/stakedotlink/token-supply-api/app"
)

var ctx = context.Background()

// VestedTokens removes vested tokens from the circulating supply
type VestedTokens struct {
	t             *app.TokenSupply
	DelegatorPool *abi.DelegatorPool
}

// Init sets the token api app
func (s *VestedTokens) Init(t *app.TokenSupply) error {
	s.t = t
	return nil
}

// Uncirculating sums the balances of all specified addresses
func (s *VestedTokens) Uncirculating() (*big.Int, error) {
	totalVested := new(big.Int)

	for _, a := range s.t.Config.VestingAddresses {
		vestedAmount, err := s.t.ERC20.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(a))
		if err != nil {
			return nil, err
		}
		totalVested = new(big.Int).Add(totalVested, vestedAmount)
	}

	return totalVested, nil
}
