package multichain

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ChainConfig represents configuration for a supported blockchain
type ChainConfig struct {
	ChainID       uint64
	Name          string
	RPC           string
	DEXAggregator string
	NativeToken   common.Address
	IsTestnet     bool
	BlockTime     uint64 // Average block time in seconds
}

// MultiChainManager handles operations across multiple blockchains
type MultiChainManager struct {
	chains  map[uint64]*ChainConfig
	clients map[uint64]*ethclient.Client
}

func NewMultiChainManager() *MultiChainManager {
	return &MultiChainManager{
		chains:  make(map[uint64]*ChainConfig),
		clients: make(map[uint64]*ethclient.Client),
	}
}

func (m *MultiChainManager) Initialize() error {
	// Configure supported chains
	chains := []*ChainConfig{
		{
			ChainID:       1,
			Name:          "Ethereum Mainnet",
			RPC:           "https://eth.llamarpc.com",
			DEXAggregator: "https://api.1inch.io/v5.0/1",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     12,
		},
		{
			ChainID:       137,
			Name:          "Polygon",
			RPC:           "https://polygon-rpc.com",
			DEXAggregator: "https://api.1inch.io/v5.0/137",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     2,
		},
		{
			ChainID:       42161,
			Name:          "Arbitrum One",
			RPC:           "https://arb1.arbitrum.io/rpc",
			DEXAggregator: "https://api.1inch.io/v5.0/42161",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     1,
		},
		{
			ChainID:       10,
			Name:          "Optimism",
			RPC:           "https://mainnet.optimism.io",
			DEXAggregator: "https://api.1inch.io/v5.0/10",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     2,
		},
		{
			ChainID:       8453,
			Name:          "Base",
			RPC:           "https://mainnet.base.org",
			DEXAggregator: "https://api.1inch.io/v5.0/8453",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     2,
		},
		{
			ChainID:       195,
			Name:          "X Layer Testnet",
			RPC:           "https://testrpc.xlayer.tech",
			DEXAggregator: "https://www.okx.com/api/v5/dex/aggregator",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     true,
			BlockTime:     3,
		},
		{
			ChainID:       196,
			Name:          "X Layer Mainnet",
			RPC:           "https://rpc.xlayer.tech",
			DEXAggregator: "https://www.okx.com/api/v5/dex/aggregator",
			NativeToken:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
			IsTestnet:     false,
			BlockTime:     3,
		},
	}

	for _, chain := range chains {
		err := m.AddChain(chain)
		if err != nil {
			log.Printf("⚠️  Failed to add chain %s: %v", chain.Name, err)
		} else {
			log.Printf("✅ Added chain: %s (ID: %d)", chain.Name, chain.ChainID)
		}
	}

	return nil
}

func (m *MultiChainManager) AddChain(config *ChainConfig) error {
	// Add chain configuration
	m.chains[config.ChainID] = config

	// Create RPC client
	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", config.Name, err)
	}

	// Verify connection
	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to verify chain ID for %s: %v", config.Name, err)
	}

	if chainID.Uint64() != config.ChainID {
		return fmt.Errorf("chain ID mismatch for %s: expected %d, got %d", 
			config.Name, config.ChainID, chainID.Uint64())
	}

	m.clients[config.ChainID] = client
	return nil
}

func (m *MultiChainManager) GetChain(chainID uint64) (*ChainConfig, error) {
	chain, exists := m.chains[chainID]
	if !exists {
		return nil, fmt.Errorf("chain %d not supported", chainID)
	}
	return chain, nil
}

func (m *MultiChainManager) GetClient(chainID uint64) (*ethclient.Client, error) {
	client, exists := m.clients[chainID]
	if !exists {
		return nil, fmt.Errorf("client for chain %d not available", chainID)
	}
	return client, nil
}

func (m *MultiChainManager) GetSupportedChains() []*ChainConfig {
	chains := make([]*ChainConfig, 0, len(m.chains))
	for _, chain := range m.chains {
		chains = append(chains, chain)
	}
	return chains
}

// CrossChainPortfolio represents a user's portfolio across multiple chains
type CrossChainPortfolio struct {
	UserAddress common.Address
	Balances    map[uint64]map[common.Address]*big.Int // chainID -> token -> balance
	TotalValue  *big.Int
	manager     *MultiChainManager
}

func NewCrossChainPortfolio(userAddress common.Address, manager *MultiChainManager) *CrossChainPortfolio {
	return &CrossChainPortfolio{
		UserAddress: userAddress,
		Balances:    make(map[uint64]map[common.Address]*big.Int),
		TotalValue:  big.NewInt(0),
		manager:     manager,
	}
}

func (p *CrossChainPortfolio) UpdateBalances(ctx context.Context) error {
	log.Printf("🔍 Updating cross-chain portfolio for %s", p.UserAddress.Hex())

	totalValue := big.NewInt(0)

	for chainID, client := range p.manager.clients {
		chainBalances := make(map[common.Address]*big.Int)

		// Get native token balance
		nativeBalance, err := client.BalanceAt(ctx, p.UserAddress, nil)
		if err != nil {
			log.Printf("⚠️  Failed to get native balance on chain %d: %v", chainID, err)
			continue
		}

		chain := p.manager.chains[chainID]
		chainBalances[chain.NativeToken] = nativeBalance

		// Convert to USD value (simplified - would use price oracles in practice)
		valueInUSD := new(big.Int).Mul(nativeBalance, big.NewInt(2000)) // Mock $2000 per ETH
		totalValue.Add(totalValue, valueInUSD)

		p.Balances[chainID] = chainBalances

		log.Printf("📊 Chain %s: %s ETH", 
			chain.Name, 
			nativeBalance.String())
	}

	p.TotalValue = totalValue
	log.Printf("💰 Total portfolio value: $%s", totalValue.String())

	return nil
}

func (p *CrossChainPortfolio) GetBalanceOnChain(chainID uint64, token common.Address) *big.Int {
	if chainBalances, exists := p.Balances[chainID]; exists {
		if balance, exists := chainBalances[token]; exists {
			return balance
		}
	}
	return big.NewInt(0)
}

func (p *CrossChainPortfolio) GetTotalBalance(token common.Address) *big.Int {
	total := big.NewInt(0)
	for _, chainBalances := range p.Balances {
		if balance, exists := chainBalances[token]; exists {
			total.Add(total, balance)
		}
	}
	return total
}

// CrossChainStrategy represents a trading strategy that operates across multiple chains
type CrossChainStrategy struct {
	ID          uint64
	Name        string
	ChainIDs    []uint64
	Active      bool
	manager     *MultiChainManager
	portfolio   *CrossChainPortfolio
}

func NewCrossChainStrategy(
	id uint64,
	name string,
	chainIDs []uint64,
	manager *MultiChainManager,
	portfolio *CrossChainPortfolio,
) *CrossChainStrategy {
	return &CrossChainStrategy{
		ID:        id,
		Name:      name,
		ChainIDs:  chainIDs,
		Active:    true,
		manager:   manager,
		portfolio: portfolio,
	}
}

func (s *CrossChainStrategy) Execute(ctx context.Context) error {
	log.Printf("🌐 Executing cross-chain strategy: %s", s.Name)

	// Update portfolio balances across all chains
	err := s.portfolio.UpdateBalances(ctx)
	if err != nil {
		return fmt.Errorf("failed to update portfolio: %v", err)
	}

	// Find arbitrage opportunities
	opportunities, err := s.FindArbitrageOpportunities(ctx)
	if err != nil {
		return fmt.Errorf("failed to find opportunities: %v", err)
	}

	// Execute profitable trades
	for _, opportunity := range opportunities {
		log.Printf("💡 Found arbitrage: %s -> %s (profit: %s%%)", 
			s.manager.chains[opportunity.ChainA].Name,
			s.manager.chains[opportunity.ChainB].Name,
			opportunity.ProfitPercentage.String())

		// Execute the arbitrage (simplified)
		err := s.ExecuteArbitrage(ctx, opportunity)
		if err != nil {
			log.Printf("❌ Failed to execute arbitrage: %v", err)
		}
	}

	return nil
}

type ArbitrageOpportunity struct {
	ChainA          uint64
	ChainB          uint64
	Token           common.Address
	PriceA          *big.Int
	PriceB          *big.Int
	ProfitPercentage *big.Int
}

func (s *CrossChainStrategy) FindArbitrageOpportunities(ctx context.Context) ([]*ArbitrageOpportunity, error) {
	opportunities := []*ArbitrageOpportunity{}

	// Compare prices across different chains
	for i, chainA := range s.ChainIDs {
		for j, chainB := range s.ChainIDs {
			if i >= j {
				continue // Avoid duplicates and self-comparison
			}

			// Mock price comparison - in practice, you'd query DEX APIs
			priceA := big.NewInt(2000) // $2000 on chain A
			priceB := big.NewInt(2050) // $2050 on chain B

			// Calculate profit percentage
			priceDiff := new(big.Int).Sub(priceB, priceA)
			profitPercentage := new(big.Int).Mul(priceDiff, big.NewInt(100))
			profitPercentage.Div(profitPercentage, priceA)

			// If profit > 1%, it's an opportunity
			if profitPercentage.Cmp(big.NewInt(1)) > 0 {
				opportunity := &ArbitrageOpportunity{
					ChainA:           chainA,
					ChainB:           chainB,
					Token:            s.manager.chains[chainA].NativeToken,
					PriceA:           priceA,
					PriceB:           priceB,
					ProfitPercentage: profitPercentage,
				}
				opportunities = append(opportunities, opportunity)
			}
		}
	}

	return opportunities, nil
}

func (s *CrossChainStrategy) ExecuteArbitrage(ctx context.Context, opportunity *ArbitrageOpportunity) error {
	log.Printf("⚡ Executing arbitrage between chains %d and %d", 
		opportunity.ChainA, opportunity.ChainB)

	// 1. Buy on cheaper chain
	clientA, err := s.manager.GetClient(opportunity.ChainA)
	if err != nil {
		return err
	}

	// 2. Bridge to more expensive chain (simplified)
	log.Printf("🌉 Bridging assets from chain %d to chain %d", 
		opportunity.ChainA, opportunity.ChainB)

	// 3. Sell on more expensive chain
	clientB, err := s.manager.GetClient(opportunity.ChainB)
	if err != nil {
		return err
	}

	// Mock execution
	_ = clientA
	_ = clientB

	log.Printf("✅ Arbitrage executed successfully")
	return nil
}

// GasOptimizer helps choose the best chain for transactions based on gas costs
type GasOptimizer struct {
	manager *MultiChainManager
}

func NewGasOptimizer(manager *MultiChainManager) *GasOptimizer {
	return &GasOptimizer{manager: manager}
}

func (g *GasOptimizer) GetBestChainForTransaction(ctx context.Context, txType string) (uint64, error) {
	bestChainID := uint64(0)
	lowestCost := big.NewInt(0)

	for chainID, client := range g.manager.clients {
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Printf("⚠️  Failed to get gas price for chain %d: %v", chainID, err)
			continue
		}

		// Estimate gas cost (simplified)
		gasLimit := big.NewInt(100000) // Base gas limit
		gasCost := new(big.Int).Mul(gasPrice, gasLimit)

		chain := g.manager.chains[chainID]
		log.Printf("⛽ %s: Gas cost = %s wei", chain.Name, gasCost.String())

		if bestChainID == 0 || gasCost.Cmp(lowestCost) < 0 {
			bestChainID = chainID
			lowestCost = gasCost
		}
	}

	if bestChainID == 0 {
		return 0, fmt.Errorf("no available chains")
	}

	chain := g.manager.chains[bestChainID]
	log.Printf("🏆 Best chain for %s: %s (cost: %s wei)", 
		txType, chain.Name, lowestCost.String())

	return bestChainID, nil
}
