package uncirculating

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stakedotlink/token-supply-api/app"
)

// ExcludedAddresses calculates uncirculating tokens by excluding balances of specified addresses
type ExcludedAddresses struct {
	t *app.TokenSupply
}

// Init sets the token api app
func (e *ExcludedAddresses) Init(t *app.TokenSupply) error {
	e.t = t
	return nil
}

// Uncirculating sums the balances of all specified addresses
func (e *ExcludedAddresses) Uncirculating() (*big.Int, error) {
	totalBalance := big.NewInt(0)
	for _, a := range e.t.Config.ExcludedAddresses {
		addr := common.HexToAddress(a)
		balance, err := e.t.ERC20.BalanceOf(&bind.CallOpts{Context: context.Background()}, addr)
		if err != nil {
			return nil, err
		}
		totalBalance = totalBalance.Add(totalBalance, balance)
	}

	return totalBalance, nil
}
