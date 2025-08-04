#!/bin/bash

# 🚀 Sentinel Agent Web Dashboard Startup Script
# This script sets up and starts the complete web dashboard with backend API

set -e

echo "🚀 Starting Sentinel Agent Web Dashboard..."
echo "=============================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

# Function to print colored output
print_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in the correct directory
if [ ! -f "package.json" ]; then
    print_error "Please run this script from the web-dashboard directory"
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js 18+ first."
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go 1.21+ first."
    exit 1
fi

print_step "Checking environment variables..."
if [ -z "$X_LAYER_RPC" ]; then
    print_warning "X_LAYER_RPC not set. Using default testnet RPC."
    export X_LAYER_RPC="https://testrpc.xlayer.tech"
fi

# Create .env.local file if it doesn't exist
if [ ! -f ".env.local" ]; then
    print_step "Creating .env.local file..."
    cat > .env.local << EOF
# Sentinel Agent Dashboard Environment Variables
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXT_PUBLIC_WS_URL=ws://localhost:8080/ws
X_LAYER_RPC=${X_LAYER_RPC:-https://testrpc.xlayer.tech}
ETHEREUM_RPC=${ETHEREUM_RPC:-}
POLYGON_RPC=${POLYGON_RPC:-}
ARBITRUM_RPC=${ARBITRUM_RPC:-}
OPTIMISM_RPC=${OPTIMISM_RPC:-}
BASE_RPC=${BASE_RPC:-}
EOF
    print_success "Created .env.local file"
else
    print_success ".env.local file already exists"
fi

# Install frontend dependencies
print_step "Installing frontend dependencies..."
if [ ! -d "node_modules" ]; then
    npm install
    print_success "Frontend dependencies installed"
else
    print_success "Frontend dependencies already installed"
fi

# Install backend dependencies
print_step "Installing backend dependencies..."
cd server
if [ ! -f "go.sum" ]; then
    go mod tidy
    print_success "Backend dependencies installed"
else
    print_success "Backend dependencies already installed"
fi
cd ..

# Build frontend
print_step "Building frontend..."
npm run build
print_success "Frontend built successfully"

# Function to start backend server
start_backend() {
    print_step "Starting backend server..."
    cd server
    
    # Set environment variables for the server
    export PORT=8080
    export X_LAYER_RPC="${X_LAYER_RPC:-https://testrpc.xlayer.tech}"
    
    # Build and run the server
    go build -o dashboard-server main.go
    print_success "Backend server built"
    
    echo -e "${PURPLE}[INFO]${NC} Backend server starting on http://localhost:8080"
    echo -e "${PURPLE}[INFO]${NC} API available at http://localhost:8080/api/v1/"
    echo -e "${PURPLE}[INFO]${NC} WebSocket available at ws://localhost:8080/ws"
    
    ./dashboard-server &
    SERVER_PID=$!
    cd ..
    
    return $SERVER_PID
}

# Function to start frontend
start_frontend() {
    print_step "Starting frontend..."
    echo -e "${PURPLE}[INFO]${NC} Frontend starting on http://localhost:3000"
    
    npm run dev &
    FRONTEND_PID=$!
    
    return $FRONTEND_PID
}

# Cleanup function
cleanup() {
    echo ""
    print_step "Shutting down services..."
    
    if [ ! -z "$SERVER_PID" ]; then
        kill $SERVER_PID 2>/dev/null || true
        print_success "Backend server stopped"
    fi
    
    if [ ! -z "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null || true
        print_success "Frontend server stopped"
    fi
    
    echo -e "${GREEN}[SHUTDOWN]${NC} Sentinel Agent Dashboard stopped"
    exit 0
}

# Set up signal handlers
trap cleanup INT TERM

# Start services
start_backend
SERVER_PID=$!

# Give backend time to start
sleep 3

start_frontend
FRONTEND_PID=$!

echo ""
echo -e "${GREEN}✅ Sentinel Agent Dashboard is now running!${NC}"
echo ""
echo -e "${BLUE}📊 Frontend Dashboard:${NC} http://localhost:3000"
echo -e "${BLUE}🔧 Backend API:${NC}       http://localhost:8080/api/v1/"
echo -e "${BLUE}🔌 WebSocket:${NC}         ws://localhost:8080/ws"
echo -e "${BLUE}💡 Health Check:${NC}      http://localhost:8080/api/v1/health"
echo ""
echo -e "${YELLOW}📋 Available API Endpoints:${NC}"
echo "   • GET  /api/v1/dashboard    - Complete dashboard data"
echo "   • GET  /api/v1/portfolio    - Portfolio information"
echo "   • GET  /api/v1/strategies   - Trading strategies"
echo "   • GET  /api/v1/transactions - Transaction history"
echo "   • GET  /api/v1/chains       - Blockchain status"
echo "   • GET  /api/v1/health       - Health check"
echo ""
echo -e "${PURPLE}🚀 Press Ctrl+C to stop all services${NC}"

# Wait for both processes
wait
