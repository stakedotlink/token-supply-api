package uncirculating

import (
	"context"
	"math/big"
	"time"

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
	vesting       []*abi.Vesting
}

// Init sets the token api app
func (s *VestedTokens) Init(t *app.TokenSupply) error {
	s.t = t

	for _, v := range s.t.Config.VestingAddresses {
		vestingContract, err := abi.NewVesting(common.HexToAddress(v), t.ETHClient)
		if err != nil {
			return err
		}
		s.vesting = append(s.vesting, vestingContract)
	}

	return nil
}

// Uncirculating sums the balances of all specified addresses
func (s *VestedTokens) Uncirculating() (*big.Int, error) {
	totalVested := new(big.Int)
	sdlToken := common.HexToAddress(s.t.Config.TokenAddress)

	for _, v := range s.vesting {
		vestedAmount, err := v.VestedAmount0(&bind.CallOpts{Context: ctx}, sdlToken, uint64(time.Now().Unix()))
		if err != nil {
			return totalVested, err
		}
		totalVested = new(big.Int).Add(totalVested, vestedAmount)
	}
	return totalVested, nil
}
