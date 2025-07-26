# Sentinel Agent Setup Guide

## Current Status ✅
- Go module and dependencies installed
- Smart contract compiled successfully
- Environment validation working
- Deployment script ready

## Next Steps to Complete Setup

### 1. Get Testnet Funds 💰
Your wallet `0xDf450Fc2AecB2eB2D1e1BD6C5D63F4628F937B61` needs testnet tokens to deploy the contract.

**Option A: Use X Layer Testnet Faucet**
- Visit: https://www.okx.com/xlayer/faucet
- Connect your wallet or paste your address
- Request testnet tokens

**Option B: Bridge from Ethereum Sepolia**
- Get Sepolia ETH from https://sepoliafaucet.com/
- Use X Layer bridge to transfer to testnet

### 2. Deploy Smart Account Contract 🚀
Once you have testnet funds, run:
```bash
./setup-and-run.sh
```

This will:
- Deploy your SmartAccount contract
- Show you the deployed address
- Prompt you to update `.env` with the contract address

### 3. Update Environment Variables 📝
After deployment, copy the contract address and update `.env`:
```bash
SMART_ACCOUNT=0xYourDeployedContractAddress
```

### 4. Run the Agent 🤖
Run the script again:
```bash
./setup-and-run.sh
```

The agent will:
- Get a swap quote from OKX DEX
- Execute a swap through your smart account
- Show transaction hash on success

## Manual Commands

If you prefer manual control:

**Deploy contract:**
```bash
cd deploy && ./deploy.sh
```

**Run agent:**
```bash
cd agent && go run .
```

**Load environment variables:**
```bash
export $(cat .env | grep -v '^#' | xargs)
```

## Troubleshooting

**"insufficient funds" error:**
- Get testnet tokens from faucet

**"Invalid private key" error:**
- Make sure PRIVATE_KEY is 64 character hex (no 0x prefix)

**"SMART_ACCOUNT not deployed" error:**
- Deploy contract first, then update .env

**OKX API errors:**
- API might need different token symbols or be temporarily unavailable
- The agent will continue with blockchain operations regardless
