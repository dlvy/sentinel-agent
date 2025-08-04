package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// Data structures for the dashboard
type DashboardData struct {
	Portfolio    PortfolioData `json:"portfolio"`
	Strategies   []Strategy    `json:"strategies"`
	Transactions []Transaction `json:"transactions"`
	Stats        StatsData     `json:"stats"`
	Chains       []ChainData   `json:"chains"`
	LastUpdated  time.Time     `json:"lastUpdated"`
}

type PortfolioData struct {
	TotalValue   string         `json:"totalValue"`
	TotalChange  string         `json:"totalChange"`
	Assets       []Asset        `json:"assets"`
	Distribution []ChainHolding `json:"distribution"`
}

type Asset struct {
	Symbol  string `json:"symbol"`
	Amount  string `json:"amount"`
	Value   string `json:"value"`
	Change  string `json:"change"`
	Address string `json:"address"`
	ChainID uint64 `json:"chainId"`
}

type ChainHolding struct {
	ChainID    uint64  `json:"chainId"`
	Name       string  `json:"name"`
	Value      string  `json:"value"`
	Percentage int     `json:"percentage"`
	Assets     []Asset `json:"assets"`
}

type Strategy struct {
	ID            int                    `json:"id"`
	Name          string                 `json:"name"`
	Type          string                 `json:"type"`
	Status        string                 `json:"status"`
	NextExecution string                 `json:"nextExecution"`
	Performance   string                 `json:"performance"`
	Invested      string                 `json:"invested"`
	CurrentValue  string                 `json:"currentValue"`
	LastExecution time.Time              `json:"lastExecution"`
	Configuration map[string]interface{} `json:"configuration"`
}

type Transaction struct {
	Hash      string    `json:"hash"`
	Type      string    `json:"type"`
	Amount    string    `json:"amount"`
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Chain     string    `json:"chain"`
	ChainID   uint64    `json:"chainId"`
	GasUsed   string    `json:"gasUsed"`
	GasPrice  string    `json:"gasPrice"`
	From      string    `json:"from"`
	To        string    `json:"to"`
}

type StatsData struct {
	TotalPortfolio    string `json:"totalPortfolio"`
	ActiveStrategies  int    `json:"activeStrategies"`
	Volume24h         string `json:"volume24h"`
	GasOptimized      string `json:"gasOptimized"`
	TotalTransactions int    `json:"totalTransactions"`
	SuccessRate       string `json:"successRate"`
	AvgGasPrice       string `json:"avgGasPrice"`
}

type ChainData struct {
	ChainID        uint64    `json:"chainId"`
	Name           string    `json:"name"`
	RPC            string    `json:"rpc"`
	NativeCurrency string    `json:"nativeCurrency"`
	BlockExplorer  string    `json:"blockExplorer"`
	Status         string    `json:"status"`
	LastBlock      uint64    `json:"lastBlock"`
	GasPrice       string    `json:"gasPrice"`
	IsActive       bool      `json:"isActive"`
	LastChecked    time.Time `json:"lastChecked"`
}

type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type DashboardServer struct {
	clients    map[*websocket.Conn]bool
	upgrader   websocket.Upgrader
	data       *DashboardData
	ethClients map[uint64]*ethclient.Client
}

func NewDashboardServer() *DashboardServer {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for development
		},
	}

	return &DashboardServer{
		clients:    make(map[*websocket.Conn]bool),
		upgrader:   upgrader,
		data:       initializeMockData(),
		ethClients: make(map[uint64]*ethclient.Client),
	}
}

func initializeMockData() *DashboardData {
	return &DashboardData{
		Portfolio: PortfolioData{
			TotalValue:  "$25,420.50",
			TotalChange: "+12.5%",
			Assets: []Asset{
				{Symbol: "ETH", Amount: "8.5", Value: "$20,400.00", Change: "+5.2%", ChainID: 195},
				{Symbol: "USDC", Amount: "3000", Value: "$3,000.00", Change: "+0.1%", ChainID: 195},
				{Symbol: "OKB", Amount: "125", Value: "$2,020.50", Change: "+8.7%", ChainID: 195},
			},
			Distribution: []ChainHolding{
				{ChainID: 195, Name: "X Layer", Value: "$18,250.30", Percentage: 72},
				{ChainID: 1, Name: "Ethereum", Value: "$4,120.10", Percentage: 16},
				{ChainID: 137, Name: "Polygon", Value: "$2,050.10", Percentage: 8},
				{ChainID: 42161, Name: "Arbitrum", Value: "$1,000.00", Percentage: 4},
			},
		},
		Strategies: []Strategy{
			{
				ID: 1, Name: "ETH DCA Strategy", Type: "DCA", Status: "active",
				NextExecution: "23 mins", Performance: "+15.3%",
				Invested: "$2,400", CurrentValue: "$2,767",
				LastExecution: time.Now().Add(-1 * time.Hour),
			},
			{
				ID: 2, Name: "Grid Trading Bot", Type: "Grid", Status: "active",
				NextExecution: "12 mins", Performance: "+8.7%",
				Invested: "$5,000", CurrentValue: "$5,435",
				LastExecution: time.Now().Add(-30 * time.Minute),
			},
		},
		Transactions: []Transaction{
			{
				Hash: "0x1a2b3c4d5e6f7890abcdef1234567890abcdef12",
				Type: "DCA Buy", Amount: "0.1", Token: "ETH", Status: "completed",
				Timestamp: time.Now().Add(-2 * time.Minute), Chain: "X Layer", ChainID: 195,
			},
			{
				Hash: "0x2b3c4d5e6f7890abcdef1234567890abcdef123a",
				Type: "Grid Order", Amount: "500", Token: "USDC", Status: "pending",
				Timestamp: time.Now().Add(-5 * time.Minute), Chain: "X Layer", ChainID: 195,
			},
		},
		Stats: StatsData{
			TotalPortfolio:    "$25,420.50",
			ActiveStrategies:  2,
			Volume24h:         "$8,420",
			GasOptimized:      "$127",
			TotalTransactions: 1247,
			SuccessRate:       "98.3%",
			AvgGasPrice:       "12 gwei",
		},
		Chains: []ChainData{
			{ChainID: 195, Name: "X Layer Testnet", Status: "active", IsActive: true, LastChecked: time.Now()},
			{ChainID: 1, Name: "Ethereum", Status: "active", IsActive: true, LastChecked: time.Now()},
			{ChainID: 137, Name: "Polygon", Status: "active", IsActive: true, LastChecked: time.Now()},
		},
		LastUpdated: time.Now(),
	}
}

func (s *DashboardServer) initializeEthClients() {
	rpcEndpoints := map[uint64]string{
		195:   os.Getenv("X_LAYER_RPC"),
		1:     os.Getenv("ETHEREUM_RPC"),
		137:   os.Getenv("POLYGON_RPC"),
		42161: os.Getenv("ARBITRUM_RPC"),
		10:    os.Getenv("OPTIMISM_RPC"),
		8453:  os.Getenv("BASE_RPC"),
	}

	for chainID, rpc := range rpcEndpoints {
		if rpc != "" {
			client, err := ethclient.Dial(rpc)
			if err != nil {
				log.Printf("Failed to connect to chain %d: %v", chainID, err)
				continue
			}
			s.ethClients[chainID] = client
			log.Printf("Connected to chain %d (%s)", chainID, rpc)
		}
	}
}

func (s *DashboardServer) updateChainData() {
	for chainID, client := range s.ethClients {
		go func(id uint64, c *ethclient.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			// Get latest block
			header, err := c.HeaderByNumber(ctx, nil)
			if err != nil {
				log.Printf("Failed to get header for chain %d: %v", id, err)
				return
			}

			// Get gas price
			gasPrice, err := c.SuggestGasPrice(ctx)
			if err != nil {
				log.Printf("Failed to get gas price for chain %d: %v", id, err)
				return
			}

			// Update chain data
			for i := range s.data.Chains {
				if s.data.Chains[i].ChainID == id {
					s.data.Chains[i].LastBlock = header.Number.Uint64()
					s.data.Chains[i].GasPrice = fmt.Sprintf("%.2f gwei", float64(gasPrice.Uint64())/1e9)
					s.data.Chains[i].LastChecked = time.Now()
					s.data.Chains[i].Status = "active"
					break
				}
			}
		}(chainID, client)
	}
}

func (s *DashboardServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	s.clients[conn] = true
	defer delete(s.clients, conn)

	log.Printf("Client connected. Total clients: %d", len(s.clients))

	// Send initial data
	s.sendToClient(conn, "initial_data", s.data)

	// Handle incoming messages
	for {
		var msg WebSocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		s.handleWebSocketMessage(conn, msg)
	}
}

func (s *DashboardServer) handleWebSocketMessage(conn *websocket.Conn, msg WebSocketMessage) {
	switch msg.Type {
	case "get_portfolio":
		s.sendToClient(conn, "portfolio_update", s.data.Portfolio)
	case "get_strategies":
		s.sendToClient(conn, "strategies_update", s.data.Strategies)
	case "get_transactions":
		s.sendToClient(conn, "transactions_update", s.data.Transactions)
	case "pause_strategy":
		if strategyID, ok := msg.Data.(float64); ok {
			s.pauseStrategy(int(strategyID))
			s.broadcastToAll("strategies_update", s.data.Strategies)
		}
	case "resume_strategy":
		if strategyID, ok := msg.Data.(float64); ok {
			s.resumeStrategy(int(strategyID))
			s.broadcastToAll("strategies_update", s.data.Strategies)
		}
	}
}

func (s *DashboardServer) sendToClient(conn *websocket.Conn, msgType string, data interface{}) {
	msg := WebSocketMessage{
		Type: msgType,
		Data: data,
	}

	err := conn.WriteJSON(msg)
	if err != nil {
		log.Printf("WebSocket write error: %v", err)
		delete(s.clients, conn)
	}
}

func (s *DashboardServer) broadcastToAll(msgType string, data interface{}) {
	msg := WebSocketMessage{
		Type: msgType,
		Data: data,
	}

	for conn := range s.clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Broadcast error: %v", err)
			delete(s.clients, conn)
			conn.Close()
		}
	}
}

func (s *DashboardServer) pauseStrategy(id int) {
	for i := range s.data.Strategies {
		if s.data.Strategies[i].ID == id {
			s.data.Strategies[i].Status = "paused"
			s.data.Strategies[i].NextExecution = "Paused"
			break
		}
	}
}

func (s *DashboardServer) resumeStrategy(id int) {
	for i := range s.data.Strategies {
		if s.data.Strategies[i].ID == id {
			s.data.Strategies[i].Status = "active"
			s.data.Strategies[i].NextExecution = "15 mins"
			break
		}
	}
}

// REST API Handlers
func (s *DashboardServer) getDashboardData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.data)
}

func (s *DashboardServer) getPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.data.Portfolio)
}

func (s *DashboardServer) getStrategies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.data.Strategies)
}

func (s *DashboardServer) getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Support pagination
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageNum := 1
	limitNum := 20

	if page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			pageNum = p
		}
	}

	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			limitNum = l
		}
	}

	// Simple pagination for mock data
	start := (pageNum - 1) * limitNum
	end := start + limitNum

	if start >= len(s.data.Transactions) {
		json.NewEncoder(w).Encode([]Transaction{})
		return
	}

	if end > len(s.data.Transactions) {
		end = len(s.data.Transactions)
	}

	paginatedTxs := s.data.Transactions[start:end]
	json.NewEncoder(w).Encode(map[string]interface{}{
		"transactions": paginatedTxs,
		"total":        len(s.data.Transactions),
		"page":         pageNum,
		"limit":        limitNum,
	})
}

func (s *DashboardServer) getChains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.data.Chains)
}

func (s *DashboardServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"version":   "1.0.0",
		"clients":   len(s.clients),
		"chains":    len(s.ethClients),
	}
	json.NewEncoder(w).Encode(response)
}

func (s *DashboardServer) startPeriodicUpdates() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			s.updateChainData()
			s.data.LastUpdated = time.Now()

			// Simulate some data changes
			s.simulateDataChanges()

			// Broadcast updates to all connected clients
			s.broadcastToAll("data_update", s.data)
		}
	}()
}

func (s *DashboardServer) simulateDataChanges() {
	// Simulate new transactions
	if len(s.data.Transactions) < 50 { // Keep reasonable number for demo
		newTx := Transaction{
			Hash:      fmt.Sprintf("0x%x", time.Now().UnixNano()),
			Type:      "Auto Swap",
			Amount:    "0.05",
			Token:     "ETH",
			Status:    "completed",
			Timestamp: time.Now(),
			Chain:     "X Layer",
			ChainID:   195,
		}
		s.data.Transactions = append([]Transaction{newTx}, s.data.Transactions...)
	}

	// Update strategy execution times
	for i := range s.data.Strategies {
		if s.data.Strategies[i].Status == "active" {
			// Simulate countdown
			if s.data.Strategies[i].NextExecution != "Paused" {
				s.data.Strategies[i].NextExecution = fmt.Sprintf("%d mins", 30-int(time.Now().Unix()%30))
			}
		}
	}
}

func main() {
	log.Println("🚀 Starting Sentinel Agent Dashboard Server...")

	server := NewDashboardServer()
	server.initializeEthClients()
	server.startPeriodicUpdates()

	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/dashboard", server.getDashboardData).Methods("GET")
	api.HandleFunc("/portfolio", server.getPortfolio).Methods("GET")
	api.HandleFunc("/strategies", server.getStrategies).Methods("GET")
	api.HandleFunc("/transactions", server.getTransactions).Methods("GET")
	api.HandleFunc("/chains", server.getChains).Methods("GET")
	api.HandleFunc("/health", server.healthCheck).Methods("GET")

	// WebSocket endpoint
	r.HandleFunc("/ws", server.handleWebSocket)

	// Serve static files (for production builds)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../out/")))

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🌐 Dashboard server running on http://localhost:%s", port)
	log.Printf("📊 API available at http://localhost:%s/api/v1/", port)
	log.Printf("🔌 WebSocket available at ws://localhost:%s/ws", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
