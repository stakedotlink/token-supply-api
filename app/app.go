package app

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stakedotlink/token-supply-api/abi"
)

var ctx = context.Background()

// Uncirculating is an interface for plugins that exclude tokens to calculate circulating supply
type Uncirculating interface {
	Init(t *TokenSupply) error
	Uncirculating() (*big.Int, error)
}

// TokenSupply represents application memory state
type TokenSupply struct {
	ETHClient *ethclient.Client
	Config    *Config
	ERC20     *abi.ERC20

	tickInterval      time.Duration
	ticker            *time.Ticker
	totalSupply       *big.Int
	circulatingSupply *big.Int
	Uncirculating     []Uncirculating
}

// NewTokenSupply returns a new instance of the app
func NewTokenSupply(configPath string, uncirculating ...Uncirculating) (*TokenSupply, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load config: %v", err)
	}
	duration, err := time.ParseDuration(cfg.Ticker)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse ticker: %v", err)
	}
	ethClient, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to rpc: %v", err)
	}
	erc20, err := abi.NewERC20(common.HexToAddress(cfg.TokenAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("couldn't instantiate erc20: %v", err)
	}

	return &TokenSupply{Config: cfg, ETHClient: ethClient, ERC20: erc20, tickInterval: duration, totalSupply: big.NewInt(0), circulatingSupply: big.NewInt(0), Uncirculating: uncirculating}, nil
}

// Start will start the app refreshing token supply
func (t *TokenSupply) Start() error {
	for _, u := range t.Uncirculating {
		if err := u.Init(t); err != nil {
			return err
		}
	}
	if err := t.tick(); err != nil {
		return err
	}
	go t.startTicker()
	return nil
}

// Stop will stop the app refreshing
func (t *TokenSupply) Stop() {
	t.ticker.Stop()
}

// TotalSupply returns the total supply of the token
func (t *TokenSupply) TotalSupply() *big.Float {
	return weiToEther(t.totalSupply)
}

// CirculatingSupply returns the circulating supply of the token
func (t *TokenSupply) CirculatingSupply() *big.Float {
	return weiToEther(t.circulatingSupply)
}

func (t *TokenSupply) tick() error {
	totalSupply, err := t.ERC20.TotalSupply(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	t.totalSupply = totalSupply

	circulatingSupply := new(big.Int)
	circulatingSupply.Set(totalSupply)
	for _, p := range t.Uncirculating {
		uncirculating, err := p.Uncirculating()
		if err != nil {
			return err
		}
		circulatingSupply = circulatingSupply.Sub(circulatingSupply, uncirculating)
	}
	t.circulatingSupply = circulatingSupply
	return nil
}

func (t *TokenSupply) startTicker() {
	t.ticker = time.NewTicker(t.tickInterval)
	for range t.ticker.C {
		if err := t.tick(); err != nil {
			fmt.Printf("error while refreshing token supply: %v\n", err)
		}
	}
}

// https://github.com/ethereum/go-ethereum/issues/21221#issuecomment-805852059
func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
