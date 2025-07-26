#!/bin/bash

echo "🚀 Sentinel Agent Deployment & Execution Script"
echo "================================================"

# Load environment variables from .env file
if [ -f .env ]; then
    echo "📄 Loading environment variables from .env..."
    export $(cat .env | grep -v '^#' | xargs)
else
    echo "❌ .env file not found!"
    exit 1
fi

# Check if required variables are set
if [ "$PRIVATE_KEY" == "your_priv_key_here" ] || [ -z "$PRIVATE_KEY" ]; then
    echo "❌ Please set PRIVATE_KEY in .env file (64 character hex string)"
    exit 1
fi

if [ "$OWNER" == "your_wallet_address_here" ] || [ -z "$OWNER" ]; then
    echo "❌ Please set OWNER in .env file (your wallet address)"
    exit 1
fi

echo "✅ Environment variables loaded"
echo "   RPC: $X_LAYER_RPC"
echo "   Owner: $OWNER"

# Check if smart account is already deployed
if [ "$SMART_ACCOUNT" == "deployed_smart_account_address_here" ] || [ -z "$SMART_ACCOUNT" ]; then
    echo "🔨 Smart Account not deployed yet. Deploying..."
    
    # Run deployment
    cd deploy
    ./deploy.sh
    cd ..
    
    echo "⚠️  Please update SMART_ACCOUNT in .env with the deployed address and run this script again"
    exit 0
else
    echo "✅ Smart Account already deployed: $SMART_ACCOUNT"
fi

# Run the Go agent
echo "🤖 Starting Go Agent..."
cd agent && go run .
