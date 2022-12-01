package uncirculating

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"token-supply-api/abi"
	"token-supply-api/app"
)

var ctx = context.Background()

// SDLNops removes
type SDLNops struct {
	t             *app.TokenSupply
	DelegatorPool *abi.DelegatorPool
}

// Init sets the token api app
func (s *SDLNops) Init(t *app.TokenSupply) error {
	s.t = t
	delegatorPool, err := abi.NewDelegatorPool(common.HexToAddress(s.t.Config.DelegatorPool), t.ETHClient)
	if err != nil {
		return err
	}
	s.DelegatorPool = delegatorPool
	return nil
}

// Uncirculating sums the balances of all specified addresses
func (s *SDLNops) Uncirculating() (*big.Int, error) {
	uncirculating := big.NewInt(0)
	for _, a := range s.t.Config.NodeOperatorAddresses {
		addr := common.HexToAddress(a)
		vestingSchedule, err := s.DelegatorPool.GetVestingSchedule(&bind.CallOpts{Context: ctx}, addr)
		if err != nil {
			return nil, err
		}
		balance, err := s.DelegatorPool.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
		if err != nil {
			return nil, err
		}
		if balance.Cmp(vestingSchedule.TotalAmount) == 1 {
			continue
		}
		uncirculating = uncirculating.Add(uncirculating, balance.Sub(vestingSchedule.TotalAmount, balance))
	}
	return uncirculating, nil
}
