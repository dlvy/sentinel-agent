package strategies

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TradingStrategy represents the interface for all trading strategies
type TradingStrategy interface {
	Execute(ctx context.Context) error
	ShouldExecute(ctx context.Context) (bool, error)
	GetType() string
	GetID() uint64
}

// DCAStrategy implements Dollar Cost Averaging
type DCAStrategy struct {
	ID                 uint64
	TokenIn            common.Address
	TokenOut           common.Address
	AmountPerExecution *big.Int
	IntervalSeconds    uint64
	LastExecution      time.Time
	TotalExecutions    uint64
	MaxExecutions      uint64
	Active             bool
	client             *ethclient.Client
	contractAddress    common.Address
	auth               *bind.TransactOpts
}

func NewDCAStrategy(
	id uint64,
	tokenIn, tokenOut common.Address,
	amountPerExecution *big.Int,
	intervalSeconds uint64,
	maxExecutions uint64,
	client *ethclient.Client,
	contractAddress common.Address,
	auth *bind.TransactOpts,
) *DCAStrategy {
	return &DCAStrategy{
		ID:                 id,
		TokenIn:            tokenIn,
		TokenOut:           tokenOut,
		AmountPerExecution: amountPerExecution,
		IntervalSeconds:    intervalSeconds,
		MaxExecutions:      maxExecutions,
		Active:             true,
		client:             client,
		contractAddress:    contractAddress,
		auth:               auth,
		LastExecution:      time.Now(),
	}
}

func (d *DCAStrategy) ShouldExecute(ctx context.Context) (bool, error) {
	if !d.Active {
		return false, nil
	}

	if d.TotalExecutions >= d.MaxExecutions {
		d.Active = false
		return false, nil
	}

	timeSinceLastExecution := time.Since(d.LastExecution)
	return timeSinceLastExecution.Seconds() >= float64(d.IntervalSeconds), nil
}

func (d *DCAStrategy) Execute(ctx context.Context) error {
	log.Printf("🔄 Executing DCA Strategy #%d: %s -> %s",
		d.ID, d.TokenIn.Hex()[:8], d.TokenOut.Hex()[:8])

	// Get swap quote
	quote, err := GetSwapQuote(
		d.TokenIn.Hex(),
		d.TokenOut.Hex(),
		d.AmountPerExecution.String(),
	)
	if err != nil {
		return fmt.Errorf("failed to get swap quote: %v", err)
	}

	// Execute the swap through Smart Account
	err = ExecuteSwapThroughSmartAccount(ctx, d.client, d.contractAddress, d.auth, quote)
	if err != nil {
		return fmt.Errorf("failed to execute swap: %v", err)
	}

	// Update strategy state
	d.LastExecution = time.Now()
	d.TotalExecutions++

	if d.TotalExecutions >= d.MaxExecutions {
		d.Active = false
		log.Printf("✅ DCA Strategy #%d completed all %d executions", d.ID, d.MaxExecutions)
	}

	return nil
}

func (d *DCAStrategy) GetType() string {
	return "DCA"
}

func (d *DCAStrategy) GetID() uint64 {
	return d.ID
}

// GridStrategy implements Grid Trading
type GridStrategy struct {
	ID              uint64
	TokenA          common.Address
	TokenB          common.Address
	GridSize        uint64
	PriceStep       *big.Int
	BasePrice       *big.Int
	GridLevels      map[uint64]bool
	Active          bool
	client          *ethclient.Client
	contractAddress common.Address
	auth            *bind.TransactOpts
}

func NewGridStrategy(
	id uint64,
	tokenA, tokenB common.Address,
	gridSize uint64,
	priceStep, basePrice *big.Int,
	client *ethclient.Client,
	contractAddress common.Address,
	auth *bind.TransactOpts,
) *GridStrategy {
	return &GridStrategy{
		ID:              id,
		TokenA:          tokenA,
		TokenB:          tokenB,
		GridSize:        gridSize,
		PriceStep:       priceStep,
		BasePrice:       basePrice,
		GridLevels:      make(map[uint64]bool),
		Active:          true,
		client:          client,
		contractAddress: contractAddress,
		auth:            auth,
	}
}

func (g *GridStrategy) ShouldExecute(ctx context.Context) (bool, error) {
	if !g.Active {
		return false, nil
	}

	// Check current price and see if any grid levels should be executed
	currentPrice, err := GetCurrentPrice(g.TokenA, g.TokenB)
	if err != nil {
		return false, err
	}

	// Calculate which grid level the current price falls into
	priceDiff := new(big.Int).Sub(currentPrice, g.BasePrice)
	gridLevel := new(big.Int).Div(priceDiff, g.PriceStep).Uint64()

	// Check if this grid level hasn't been executed yet
	return !g.GridLevels[gridLevel], nil
}

func (g *GridStrategy) Execute(ctx context.Context) error {
	log.Printf("📊 Executing Grid Strategy #%d", g.ID)

	currentPrice, err := GetCurrentPrice(g.TokenA, g.TokenB)
	if err != nil {
		return err
	}

	priceDiff := new(big.Int).Sub(currentPrice, g.BasePrice)
	gridLevel := new(big.Int).Div(priceDiff, g.PriceStep).Uint64()

	// Mark this grid level as executed
	g.GridLevels[gridLevel] = true

	// Determine trade direction based on price level
	var tokenIn, tokenOut common.Address
	if currentPrice.Cmp(g.BasePrice) > 0 {
		// Price above base: sell TokenA for TokenB
		tokenIn, tokenOut = g.TokenA, g.TokenB
	} else {
		// Price below base: buy TokenA with TokenB
		tokenIn, tokenOut = g.TokenB, g.TokenA
	}

	// Calculate trade amount (simplified - could be more sophisticated)
	tradeAmount := new(big.Int).SetUint64(1000000000000000000) // 1 token

	quote, err := GetSwapQuote(tokenIn.Hex(), tokenOut.Hex(), tradeAmount.String())
	if err != nil {
		return err
	}

	return ExecuteSwapThroughSmartAccount(ctx, g.client, g.contractAddress, g.auth, quote)
}

func (g *GridStrategy) GetType() string {
	return "Grid"
}

func (g *GridStrategy) GetID() uint64 {
	return g.ID
}

// RebalanceStrategy implements Portfolio Rebalancing
type RebalanceStrategy struct {
	ID                 uint64
	Tokens             []common.Address
	TargetPercentages  []uint64 // basis points
	RebalanceThreshold uint64   // percentage deviation to trigger
	MinInterval        time.Duration
	LastRebalance      time.Time
	Active             bool
	client             *ethclient.Client
	contractAddress    common.Address
	auth               *bind.TransactOpts
}

func NewRebalanceStrategy(
	id uint64,
	tokens []common.Address,
	targetPercentages []uint64,
	rebalanceThreshold uint64,
	minInterval time.Duration,
	client *ethclient.Client,
	contractAddress common.Address,
	auth *bind.TransactOpts,
) *RebalanceStrategy {
	return &RebalanceStrategy{
		ID:                 id,
		Tokens:             tokens,
		TargetPercentages:  targetPercentages,
		RebalanceThreshold: rebalanceThreshold,
		MinInterval:        minInterval,
		Active:             true,
		client:             client,
		contractAddress:    contractAddress,
		auth:               auth,
		LastRebalance:      time.Now(),
	}
}

func (r *RebalanceStrategy) ShouldExecute(ctx context.Context) (bool, error) {
	if !r.Active {
		return false, nil
	}

	if time.Since(r.LastRebalance) < r.MinInterval {
		return false, nil
	}

	// Check if portfolio deviation exceeds threshold
	currentBalances, err := GetPortfolioBalances(r.Tokens, r.contractAddress)
	if err != nil {
		return false, err
	}

	totalValue := big.NewInt(0)
	for _, balance := range currentBalances {
		totalValue.Add(totalValue, balance)
	}

	// Calculate current percentages and check deviation
	for i, balance := range currentBalances {
		currentPercentage := new(big.Int).Mul(balance, big.NewInt(10000))
		currentPercentage.Div(currentPercentage, totalValue)

		targetPercentage := big.NewInt(int64(r.TargetPercentages[i]))
		deviation := new(big.Int).Sub(currentPercentage, targetPercentage)
		deviation.Abs(deviation)

		threshold := big.NewInt(int64(r.RebalanceThreshold))
		if deviation.Cmp(threshold) > 0 {
			return true, nil
		}
	}

	return false, nil
}

func (r *RebalanceStrategy) Execute(ctx context.Context) error {
	log.Printf("⚖️  Executing Rebalance Strategy #%d", r.ID)

	// Get current balances
	currentBalances, err := GetPortfolioBalances(r.Tokens, r.contractAddress)
	if err != nil {
		return err
	}

	totalValue := big.NewInt(0)
	for _, balance := range currentBalances {
		totalValue.Add(totalValue, balance)
	}

	// Execute rebalancing trades
	for i, token := range r.Tokens {
		currentBalance := currentBalances[i]
		targetValue := new(big.Int).Mul(totalValue, big.NewInt(int64(r.TargetPercentages[i])))
		targetValue.Div(targetValue, big.NewInt(10000))

		difference := new(big.Int).Sub(targetValue, currentBalance)

		if difference.Sign() != 0 {
			// Need to trade to reach target
			log.Printf("🔄 Rebalancing %s: current=%s, target=%s, diff=%s",
				token.Hex()[:8], currentBalance.String(), targetValue.String(), difference.String())

			// Execute the necessary trades (simplified implementation)
			// In practice, this would involve complex multi-token swaps
		}
	}

	r.LastRebalance = time.Now()
	return nil
}

func (r *RebalanceStrategy) GetType() string {
	return "Rebalance"
}

func (r *RebalanceStrategy) GetID() uint64 {
	return r.ID
}

// Helper functions (these would be implemented based on your DEX integration)

func GetCurrentPrice(tokenA, tokenB common.Address) (*big.Int, error) {
	// This would integrate with price oracles or DEX APIs
	// For now, return a mock price
	return big.NewInt(1000000000000000000), nil // 1 ETH
}

func GetPortfolioBalances(tokens []common.Address, account common.Address) ([]*big.Int, error) {
	// This would query token balances for the account
	balances := make([]*big.Int, len(tokens))
	for i := range tokens {
		balances[i] = big.NewInt(1000000000000000000) // Mock 1 token balance
	}
	return balances, nil
}

func ExecuteSwapThroughSmartAccount(ctx context.Context, client *ethclient.Client, contractAddress common.Address, auth *bind.TransactOpts, quote *SwapQuote) error {
	// This would execute the swap through the Smart Account
	log.Printf("💱 Executing swap through Smart Account: %s", contractAddress.Hex())
	return nil
}

// SwapQuote represents a swap quote from DEX
type SwapQuote struct {
	Price string `json:"price"`
	To    string `json:"to"`
	Data  string `json:"data"`
}

func GetSwapQuote(tokenIn, tokenOut, amount string) (*SwapQuote, error) {
	// This would integrate with your existing OKX DEX API
	return &SwapQuote{
		Price: "1.0",
		To:    "0x1234567890abcdef1234567890abcdef12345678",
		Data:  "0x",
	}, nil
}
