# Blockchain Module

This folder contains all Solidity smart contracts and Foundry-related files for the Sentinel Agent project.

## Structure

```
blockchain/
├── contracts/          # Smart contracts
│   └── SmartAccount.sol
├── test/              # Contract tests
│   └── SmartAccount.t.sol
├── scripts/           # Deployment and utility scripts
│   ├── deploy-smart-account.sh
│   └── deploy.sh
├── out/              # Compiled contracts (auto-generated)
├── foundry.toml      # Foundry configuration
└── README.md         # This file
```

## Quick Start

### Prerequisites
- Foundry installed
- Environment variables set in project root `.env`

### Deploy Smart Account

```bash
# From the blockchain directory
cd blockchain

# Deploy to testnet
./scripts/deploy-smart-account.sh testnet

# Deploy to mainnet
./scripts/deploy-smart-account.sh mainnet
```

### Run Tests

```bash
# From the blockchain directory
forge test

# Run with verbosity
forge test -vv

# Run specific test
forge test --match-test testOwnerIsSet
```

### Compile Contracts

```bash
forge build
```

### Gas Usage Report

```bash
forge test --gas-report
```

## Smart Contracts

### SmartAccount.sol
A simple smart account implementation with:
- Owner-based access control
- Session key management for temporary access
- Execute function for calling other contracts
- Gas-efficient design

## Environment Variables

Required in project root `.env`:
- `PRIVATE_KEY` - Deployer private key (64 char hex, no 0x)
- `OWNER` - Smart account owner address
- `XLAYER_ETHERSCAN_API_KEY` - For contract verification (optional)

## Networks

- **X Layer Testnet**: Chain ID 195, RPC: https://testrpc.xlayer.tech
- **X Layer Mainnet**: Chain ID 196, RPC: https://rpc.xlayer.tech
