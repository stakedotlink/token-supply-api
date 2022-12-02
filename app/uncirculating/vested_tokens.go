package uncirculating

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stakedotlink/token-supply-api/abi"
	"github.com/stakedotlink/token-supply-api/app"
	"math/big"
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
	delegatorPool, err := abi.NewDelegatorPool(common.HexToAddress(s.t.Config.DelegatorPool), t.ETHClient)
	if err != nil {
		return fmt.Errorf("couldn't initialise the delegator pool: %v", err)
	}
	s.DelegatorPool = delegatorPool
	return nil
}

// Uncirculating sums the balances of all specified addresses
func (s *VestedTokens) Uncirculating() (*big.Int, error) {
	totalVested := new(big.Int)
	totalWithdrawn := new(big.Int)

	var addresses []common.Address
	for _, a := range s.t.Config.NodeOperatorAddresses {
		addresses = append(addresses, common.HexToAddress(a))
		vestingSchedule, err := s.DelegatorPool.GetVestingSchedule(&bind.CallOpts{Context: ctx}, common.HexToAddress(a))
		if err != nil {
			return nil, err
		}
		totalVested = new(big.Int).Add(totalVested, vestingSchedule.TotalAmount)
	}
	iterator, err := s.DelegatorPool.FilterAllowanceWithdrawn(&bind.FilterOpts{Context: ctx, Start: s.t.Config.TokenMintBlockNumber}, addresses)
	if err != nil {
		return nil, err
	}
	for {
		if !iterator.Next() {
			break
		}
		if iterator.Error() != nil {
			return nil, iterator.Error()
		}
		event := iterator.Event
		if event == nil {
			break
		}
		totalWithdrawn = new(big.Int).Add(totalWithdrawn, event.Amount)
	}

	if totalWithdrawn.Cmp(totalVested) == 1 {
		return big.NewInt(0), nil
	}
	return new(big.Int).Sub(totalVested, totalWithdrawn), nil
}
