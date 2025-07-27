set -e

# Legacy deployment script - use blockchain/scripts/deploy-smart-account.sh instead
echo "⚠️  This script is deprecated. Use the new deployment script:"
echo "   cd blockchain && ./scripts/deploy-smart-account.sh"
echo ""
echo "🔄 Running new deployment script..."

cd ../blockchain
./scripts/deploy-smart-account.sh testnet
