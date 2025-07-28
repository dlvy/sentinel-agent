#!/bin/bash

# Smart Account V2 Deployment Script
# ===================================

set -e

echo "🚀 Deploying Smart Account V2 with Advanced Features..."

# Check required environment variables
if [ -z "$PRIVATE_KEY" ]; then
    echo "❌ PRIVATE_KEY environment variable is required"
    exit 1
fi

if [ -z "$OWNER" ]; then
    echo "❌ OWNER environment variable is required"
    exit 1
fi

# Network selection
NETWORK=${1:-testnet}

if [ "$NETWORK" = "mainnet" ]; then
    RPC_URL="https://rpc.xlayer.tech"
    echo "🌐 Deploying to X Layer Mainnet"
elif [ "$NETWORK" = "testnet" ]; then
    RPC_URL="https://testrpc.xlayer.tech"
    echo "🧪 Deploying to X Layer Testnet"
else
    echo "❌ Invalid network. Use 'mainnet' or 'testnet'"
    exit 1
fi

echo "📋 Deployment Details:"
echo "   Network: $NETWORK"
echo "   RPC: $RPC_URL"
echo "   Owner: $OWNER"
echo ""

# Deploy V2 Smart Account with advanced features
echo "🔨 Starting Smart Account V2 deployment..."
forge create \
    --rpc-url "$RPC_URL" \
    --private-key "$PRIVATE_KEY" \
    --legacy \
    --broadcast \
    --verify \
    contracts/SmartAccountV2.sol:SmartAccountV2 \
    --constructor-args "$OWNER"

echo ""
echo "✅ Smart Account V2 deployed successfully!"
echo ""
echo "🔧 Next steps:"
echo "1. Copy the deployed address to your .env file as SMART_ACCOUNT"
echo "2. Set ENABLE_STRATEGIES=true to enable advanced trading strategies"
echo "3. Set ENABLE_MULTICHAIN=true to enable multi-chain portfolio tracking"
echo "4. Run: ./run-agent-v2.sh"
