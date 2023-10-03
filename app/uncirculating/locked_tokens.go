package uncirculating

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stakedotlink/token-supply-api/abi"
	"github.com/stakedotlink/token-supply-api/app"
)

// LockedTokens removes locked tokens from the circulating supply
type LockedTokens struct {
	t       *app.TokenSupply
	SdlPool *abi.SdlPool
}

// Init sets the token api app
func (s *LockedTokens) Init(t *app.TokenSupply) error {
	s.t = t
	sdlPool, err := abi.NewSdlPool(common.HexToAddress(s.t.Config.SDLPool), t.ETHClient)
	if err != nil {
		return fmt.Errorf("couldn't initialise the sdl pool: %v", err)
	}
	s.SdlPool = sdlPool
	return nil
}

// Uncirculating sums the balances of all locked tokens
func (s *LockedTokens) Uncirculating() (*big.Int, error) {
	totalLocked := new(big.Int)
	lastLockId, err := s.SdlPool.LastLockId(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, err
	}

	for i := 0; i < int(lastLockId.Int64()); i++ {
		lockId := big.NewInt(int64(i) + 1)
		lockArray := []*big.Int{lockId}

		locks, err := s.SdlPool.GetLocks(&bind.CallOpts{Context: context.Background()}, lockArray)
		if err != nil {
			fmt.Printf("Error fetching lock %d: %s\n", i+1, err.Error())
			continue
		}
		if locks[0].Duration > 0 {
			totalLocked.Add(totalLocked, locks[0].Amount)
		}
	}

	return totalLocked, nil
}
