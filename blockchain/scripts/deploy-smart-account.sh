#!/bin/bash

# Smart Account Deployment Script
# ===============================

set -e

echo "🚀 Deploying Smart Account to X Layer..."

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

# Deploy with Foundry
echo "🔨 Starting deployment..."
forge create \
    --rpc-url "$RPC_URL" \
    --private-key "$PRIVATE_KEY" \
    --legacy \
    --broadcast \
    --verify \
    contracts/SmartAccount.sol:SmartAccount \
    --constructor-args "$OWNER"

echo ""
echo "✅ Deployment completed!"
echo "💡 Update your .env file with the deployed contract address"
