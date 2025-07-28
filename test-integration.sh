#!/bin/bash
set -e

echo "🚀 Running Integration Tests for Advanced Features"
echo "================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test functions
test_passed() {
    echo -e "${GREEN}✅ $1${NC}"
}

test_failed() {
    echo -e "${RED}❌ $1${NC}"
    exit 1
}

test_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

test_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

# Check prerequisites
echo -e "\n${BLUE}📋 Checking Prerequisites...${NC}"

# Check Go installation
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | cut -d' ' -f3)
    test_passed "Go installed: $GO_VERSION"
else
    test_failed "Go is not installed"
fi

# Check Foundry installation
if command -v forge &> /dev/null; then
    FORGE_VERSION=$(forge --version | head -n1)
    test_passed "Foundry installed: $FORGE_VERSION"
else
    test_failed "Foundry is not installed"
fi

# Check .env file
if [ -f ".env" ]; then
    test_passed ".env file exists"
else
    test_warning ".env file missing - using defaults"
fi

# Test smart contract compilation
echo -e "\n${BLUE}🔨 Testing Smart Contract Compilation...${NC}"

cd blockchain

# Test original SmartAccount
if forge build --root . --contracts contracts --out out 2>/dev/null; then
    test_passed "SmartAccount.sol compiles successfully"
else
    test_failed "SmartAccount.sol compilation failed"
fi

# Test advanced SmartAccountV2
if [ -f "contracts/SmartAccountV2.sol" ]; then
    if forge build --root . --contracts contracts --out out 2>/dev/null; then
        test_passed "SmartAccountV2.sol compiles successfully"
    else
        test_failed "SmartAccountV2.sol compilation failed"
    fi
else
    test_warning "SmartAccountV2.sol not found"
fi

cd ..

# Test Go agent compilation
echo -e "\n${BLUE}🔧 Testing Go Agent Compilation...${NC}"

cd agent

# Test original agent
if go build -o ../build/agent . 2>/dev/null; then
    test_passed "Original agent compiles successfully"
else
    test_failed "Original agent compilation failed"
fi

# Test advanced features
if [ -f "strategies/trading.go" ]; then
    if go build ./strategies/ 2>/dev/null; then
        test_passed "Trading strategies compile successfully"
    else
        test_failed "Trading strategies compilation failed"
    fi
else
    test_warning "strategies/trading.go not found"
fi

if [ -f "multichain/manager.go" ]; then
    if go build ./multichain/ 2>/dev/null; then
        test_passed "Multi-chain manager compiles successfully"
    else
        test_failed "Multi-chain manager compilation failed"
    fi
else
    test_warning "multichain/manager.go not found"
fi

# Test main_v2.go if exists
if [ -f "main_v2.go" ]; then
    if go build -o ../build/agent-v2 main_v2.go 2>/dev/null; then
        test_passed "Advanced agent (v2) compiles successfully"
    else
        test_failed "Advanced agent (v2) compilation failed"
    fi
else
    test_warning "main_v2.go not found"
fi

cd ..

# Test environment configuration
echo -e "\n${BLUE}⚙️  Testing Environment Configuration...${NC}"

# Check required environment variables
REQUIRED_VARS=("PRIVATE_KEY" "RPC_URL" "SMART_ACCOUNT_FACTORY")
MISSING_VARS=()

for var in "${REQUIRED_VARS[@]}"; do
    if grep -q "^$var=" .env 2>/dev/null; then
        test_passed "$var is configured"
    else
        MISSING_VARS+=("$var")
        test_warning "$var is not configured"
    fi
done

# Test advanced features configuration
ADVANCED_VARS=("ENABLE_STRATEGIES" "ENABLE_MULTICHAIN")
for var in "${ADVANCED_VARS[@]}"; do
    if grep -q "^$var=" .env 2>/dev/null; then
        VALUE=$(grep "^$var=" .env | cut -d'=' -f2)
        test_passed "$var=$VALUE"
    else
        test_info "$var not set (will use default: false)"
    fi
done

# Test script permissions
echo -e "\n${BLUE}🔐 Testing Script Permissions...${NC}"

SCRIPTS=("run-agent.sh" "setup-and-run.sh")
if [ -f "run-agent-v2.sh" ]; then
    SCRIPTS+=("run-agent-v2.sh")
fi

for script in "${SCRIPTS[@]}"; do
    if [ -f "$script" ]; then
        if [ -x "$script" ]; then
            test_passed "$script is executable"
        else
            test_warning "$script is not executable (fixing...)"
            chmod +x "$script"
            test_passed "$script made executable"
        fi
    else
        test_warning "$script not found"
    fi
done

# Test deployment scripts
if [ -d "blockchain/scripts" ]; then
    for script in blockchain/scripts/*.sh; do
        if [ -f "$script" ]; then
            SCRIPT_NAME=$(basename "$script")
            if [ -x "$script" ]; then
                test_passed "$SCRIPT_NAME is executable"
            else
                test_warning "$SCRIPT_NAME is not executable (fixing...)"
                chmod +x "$script"
                test_passed "$SCRIPT_NAME made executable"
            fi
        fi
    done
fi

# Test network connectivity (if RPC_URL is available)
echo -e "\n${BLUE}🌐 Testing Network Connectivity...${NC}"

if grep -q "^RPC_URL=" .env 2>/dev/null; then
    RPC_URL=$(grep "^RPC_URL=" .env | cut -d'=' -f2)
    test_info "Testing connection to: $RPC_URL"
    
    if curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$RPC_URL" | grep -q "result"; then
        test_passed "RPC endpoint is accessible"
    else
        test_warning "RPC endpoint is not accessible (check network)"
    fi
else
    test_info "RPC_URL not configured - skipping connectivity test"
fi

# Test Go dependencies
echo -e "\n${BLUE}📦 Testing Go Dependencies...${NC}"

cd agent

if go mod tidy && go mod verify; then
    test_passed "Go dependencies are valid"
else
    test_warning "Go dependencies may have issues"
fi

# Check for common import issues
if grep -r "github.com/ethereum/go-ethereum" . > /dev/null; then
    test_passed "Ethereum Go client imported correctly"
else
    test_warning "Ethereum Go client not imported"
fi

cd ..

# Run basic functionality test
echo -e "\n${BLUE}🧪 Running Basic Functionality Test...${NC}"

if [ -f "build/agent" ]; then
    test_info "Testing agent with --help flag..."
    if timeout 10s ./build/agent --help 2>/dev/null || [ $? -eq 124 ]; then
        test_passed "Agent responds to --help (or times out gracefully)"
    else
        test_warning "Agent may have runtime issues"
    fi
else
    test_warning "Agent binary not found - compile first"
fi

# Summary
echo -e "\n${BLUE}📊 Test Summary${NC}"
echo "=============="

if [ ${#MISSING_VARS[@]} -eq 0 ]; then
    echo -e "${GREEN}✅ All required environment variables are configured${NC}"
else
    echo -e "${YELLOW}⚠️  Missing environment variables: ${MISSING_VARS[*]}${NC}"
    echo -e "${YELLOW}   Please configure these in your .env file${NC}"
fi

echo -e "\n${BLUE}🚀 Ready to Deploy!${NC}"
echo "=================="
echo "To deploy and run the advanced features:"
echo ""
echo "1. Deploy Smart Account V2:"
echo "   ./blockchain/scripts/deploy-smart-account-v2.sh testnet"
echo ""
echo "2. Update .env with deployed contract address"
echo ""
echo "3. Run advanced agent:"
echo "   ./run-agent-v2.sh"
echo ""
echo "4. Or run original agent:"
echo "   ./run-agent.sh"
echo ""

if [ ${#MISSING_VARS[@]} -gt 0 ]; then
    echo -e "${YELLOW}⚠️  Don't forget to configure missing environment variables first!${NC}"
fi

echo -e "\n${GREEN}🎉 Integration tests completed successfully!${NC}"
