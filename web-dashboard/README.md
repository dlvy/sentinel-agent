# 🌐 Sentinel Agent Web Dashboard

> **Real-time monitoring and control dashboard for your DeFi Guardian on X Layer**

![Dashboard Preview](https://img.shields.io/badge/Dashboard-Live-brightgreen)
![React](https://img.shields.io/badge/React-18+-blue)
![TypeScript](https://img.shields.io/badge/TypeScript-5+-blue)
![Go](https://img.shields.io/badge/Go-1.21+-blue)
![WebSocket](https://img.shields.io/badge/WebSocket-Real--time-orange)

## ✨ Features

### 📊 **Real-Time Dashboard**
- **Live Portfolio Tracking** - Monitor your multi-chain assets in real-time
- **Strategy Performance** - Track DCA, Grid Trading, and Rebalancing strategies
- **Transaction History** - Complete transaction log with status tracking
- **Gas Optimization** - Real-time gas price monitoring across chains

### 🔄 **Live Data Updates**
- **WebSocket Integration** - Real-time data streaming
- **Auto-Refresh** - Periodic updates every 30 seconds
- **Connection Status** - Live connection indicator
- **Notification System** - Instant alerts for important events

### 📱 **Responsive Design**
- **Mobile-First** - Works perfectly on all devices
- **Modern UI** - Beautiful glass-morphism design
- **Dark Theme** - Eye-friendly dark interface
- **Smooth Animations** - Fluid transitions and micro-interactions

### 🔗 **Multi-Chain Support**
- **X Layer** - Primary chain integration
- **Ethereum** - Full Ethereum support
- **Polygon** - Polygon network monitoring
- **Arbitrum** - Layer 2 optimization
- **Optimism** - OP Stack support
- **Base** - Coinbase L2 integration

## 🚀 Quick Start

### Prerequisites
- **Node.js** 18+ installed
- **Go** 1.21+ installed
- **Sentinel Agent** running (from main project)

### 1. One-Command Setup
```bash
# Navigate to the web dashboard directory
cd web-dashboard

# Run the automated setup script
./start-dashboard.sh
```

This script will:
- ✅ Install all dependencies (frontend & backend)
- ✅ Build the React application
- ✅ Start the Go backend server
- ✅ Launch the development server
- ✅ Set up environment variables

### 2. Manual Setup (Alternative)

#### Frontend Setup
```bash
# Install dependencies
npm install

# Create environment file
cp .env.example .env.local

# Start development server
npm run dev
```

#### Backend Setup
```bash
# Navigate to server directory
cd server

# Install Go dependencies
go mod tidy

# Build and run server
go build -o dashboard-server main.go
./dashboard-server
```

## 🌍 Access Points

Once running, access the dashboard at:

- **🎨 Frontend Dashboard:** http://localhost:3000
- **🔧 Backend API:** http://localhost:8080/api/v1/
- **🔌 WebSocket:** ws://localhost:8080/ws
- **💊 Health Check:** http://localhost:8080/api/v1/health

## 📡 API Reference

### REST Endpoints

#### Dashboard Data
```
GET /api/v1/dashboard
```
Returns complete dashboard data including portfolio, strategies, and transactions.

#### Portfolio Information
```
GET /api/v1/portfolio
```
Returns detailed portfolio information across all chains.

#### Trading Strategies
```
GET /api/v1/strategies
```
Returns all active and paused trading strategies.

#### Transaction History
```
GET /api/v1/transactions?page=1&limit=20
```
Returns paginated transaction history.

#### Chain Status
```
GET /api/v1/chains
```
Returns status of all supported blockchain networks.

### WebSocket Events

#### Client → Server
```json
{
  "type": "get_portfolio",
  "data": null
}
```

```json
{
  "type": "pause_strategy",
  "data": 1
}
```

```json
{
  "type": "resume_strategy", 
  "data": 1
}
```

#### Server → Client
```json
{
  "type": "initial_data",
  "data": { /* complete dashboard data */ }
}
```

```json
{
  "type": "portfolio_update",
  "data": { /* portfolio data */ }
}
```

```json
{
  "type": "strategies_update",
  "data": [ /* strategies array */ ]
}
```

## 🏗️ Architecture

```
┌─────────────────────┐    ┌─────────────────────┐    ┌─────────────────────┐
│                     │    │                     │    │                     │
│   React Frontend    │◄──►│   Go Backend API    │◄──►│   Sentinel Agent    │
│   (Next.js + TS)    │    │   (WebSocket + REST)│    │   (Main Process)    │
│                     │    │                     │    │                     │
└─────────────────────┘    └─────────────────────┘    └─────────────────────┘
         │                           │                           │
         ▼                           ▼                           ▼
┌─────────────────────┐    ┌─────────────────────┐    ┌─────────────────────┐
│   Browser Client    │    │   WebSocket Hub     │    │   Multi-Chain       │
│   (Real-time UI)    │    │   (Live Updates)    │    │   Connections       │
└─────────────────────┘    └─────────────────────┘    └─────────────────────┘
```

## 📱 Dashboard Sections

### 1. Overview Tab
- **Portfolio Stats** - Total value, change, distribution
- **Active Strategies** - Live strategy status
- **Recent Activity** - Latest transactions
- **Performance Charts** - Visual analytics

### 2. Portfolio Tab
- **Multi-Chain Assets** - Complete asset breakdown
- **Chain Distribution** - Value allocation across networks
- **Asset Performance** - Individual token performance
- **Balance History** - Historical portfolio changes

### 3. Strategies Tab
- **Strategy Cards** - Visual strategy management
- **Performance Metrics** - ROI and execution stats
- **Control Actions** - Pause/resume/edit strategies
- **Execution Schedule** - Next execution timers

### 4. Transactions Tab
- **Transaction Log** - Complete transaction history
- **Status Tracking** - Pending/completed/failed states
- **Chain Filtering** - Filter by blockchain network
- **Export Options** - Download transaction data

### 5. Multi-Chain Tab
- **Network Status** - Live blockchain health
- **Gas Prices** - Real-time gas monitoring
- **Block Heights** - Latest block information
- **Connection Health** - RPC connection status

### 6. Analytics Tab
- **Performance Charts** - Interactive data visualization
- **Profit/Loss** - Detailed P&L analysis
- **Strategy Comparison** - Strategy performance metrics
- **Historical Data** - Time-series analysis

## 🔧 Configuration

### Environment Variables

Create `.env.local` file with:

```env
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXT_PUBLIC_WS_URL=ws://localhost:8080/ws

# Blockchain RPC Endpoints
X_LAYER_RPC=https://testrpc.xlayer.tech
ETHEREUM_RPC=https://eth.llamarpc.com
POLYGON_RPC=https://polygon-rpc.com
ARBITRUM_RPC=https://arb1.arbitrum.io/rpc
OPTIMISM_RPC=https://mainnet.optimism.io
BASE_RPC=https://mainnet.base.org

# Dashboard Settings
REFRESH_INTERVAL=30000
MAX_TRANSACTIONS=100
ENABLE_NOTIFICATIONS=true
```

### Backend Configuration

The Go backend can be configured via environment variables:

```env
PORT=8080
X_LAYER_RPC=https://testrpc.xlayer.tech
# Add other RPC endpoints as needed
```

## 🛠️ Development

### Frontend Development
```bash
# Start development server with hot reload
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Run linting
npm run lint
```

### Backend Development
```bash
cd server

# Run with hot reload (requires air)
air

# Build manually
go build -o dashboard-server main.go

# Run tests
go test ./...
```

### Project Structure
```
web-dashboard/
├── src/
│   ├── app/                 # Next.js app directory
│   ├── components/          # React components
│   ├── hooks/              # Custom React hooks
│   └── globals.css         # Global styles
├── server/                 # Go backend
│   ├── main.go            # Server implementation
│   └── go.mod             # Go modules
├── public/                # Static assets
├── package.json           # Frontend dependencies
├── next.config.js         # Next.js configuration
├── tailwind.config.js     # Tailwind CSS config
└── start-dashboard.sh     # Startup script
```

## 🔍 Troubleshooting

### Common Issues

#### WebSocket Connection Failed
```bash
# Check if backend is running
curl http://localhost:8080/api/v1/health

# Check firewall settings
sudo ufw allow 8080
```

#### Frontend Build Errors
```bash
# Clear cache and reinstall
rm -rf node_modules package-lock.json
npm install
```

#### Backend API Errors
```bash
# Check Go version
go version

# Rebuild dependencies
cd server && go mod tidy
```

### Debug Mode

Enable debug logging:
```bash
# Frontend
NEXT_PUBLIC_DEBUG=true npm run dev

# Backend
DEBUG=true ./dashboard-server
```

## 🚀 Production Deployment

### Frontend (Vercel/Netlify)
```bash
# Build for production
npm run build

# Export static files
npm run export
```

### Backend (Docker)
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY server/ .
RUN go build -o dashboard-server main.go

FROM alpine:latest
COPY --from=builder /app/dashboard-server /dashboard-server
EXPOSE 8080
CMD ["./dashboard-server"]
```

### Environment Setup
```bash
# Production environment variables
export NODE_ENV=production
export PORT=8080
export X_LAYER_RPC=https://rpc.xlayer.tech
```

## 📈 Performance

- **Frontend**: Optimized React with Next.js SSG
- **Backend**: Efficient Go with connection pooling
- **WebSocket**: Real-time updates with automatic reconnection
- **Caching**: Strategic caching for better performance
- **Compression**: Gzip compression enabled

## 🔒 Security

- **CORS Protection** - Configured for production domains
- **Rate Limiting** - API endpoint protection
- **Input Validation** - Comprehensive input sanitization
- **Error Handling** - Graceful error management
- **Environment Variables** - Secure configuration management

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/dashboard-enhancement`)
3. Commit changes (`git commit -am 'Add new dashboard feature'`)
4. Push to branch (`git push origin feature/dashboard-enhancement`)
5. Create a Pull Request

## 📄 License

This project is part of the Sentinel Agent ecosystem and follows the same MIT license.

---

<div align="center">

**Built with ❤️ for the DeFi community**

[🚀 Back to Main Project](../README.md) • [📊 Live Demo](#) • [🐛 Report Issues](#)

</div>
