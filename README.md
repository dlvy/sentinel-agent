# 🚀 Sentinel Agent - Your DeFi Guardian on X Layer 🛡️

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Chain](https://img.shields.io/badge/chain-X%20Layer-orange.svg)
![Language](https://img.shields.io/badge/language-Go%20%2B%20Solidity-green.svg)
![Status](https://img.shields.io/badge/status-Ready%20to%20Moon-gold.svg)

> **🔥 WAGMI Alert!** This isn't just another agent - this is your personal DeFi sentinel, deployed on the cutting-edge X Layer blockchain, ready to execute swaps with the precision of a diamond-handed chad! 💎🙌

## 🌟 What Makes This Special?

**Sentinel Agent** is the ultimate crypto automation beast that combines:
- 🤖 **Smart Account Architecture** - Account abstraction for the win!
- ⚡ **OKX DEX Integration** - Powered by OKX's lightning-fast aggregator
- 🔗 **X Layer Native** - Built for OKX's Layer 2 ecosystem
- 🛡️ **Session Key Security** - Automated trading without compromising your keys
- 🎯 **One-Click Deployment** - Because time is money, anon!

## 🏗️ Architecture That Hits Different

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Sentinel      │    │   Smart         │    │   OKX DEX       │
│   Agent         │◄──►│   Account       │◄──►│   Aggregator    │
│   (Go)          │    │   (Solidity)    │    │   API           │
└─────────────────┘    └─────────────────┘    └─────────────────┘
        │                       │                       │
        ▼                       ▼                       ▼
   🤖 Automation          🔐 Security             💱 Best Prices
```

## ✨ Features That Make Degens Cry (Happy Tears)

- 🎯 **Auto-Swap Execution** - Set it and forget it (but don't actually forget it)
- 🔐 **Session Key Management** - Temporary permissions for automated trading
- ⚡ **Gas Optimization** - Because every wei counts
- 🌊 **MEV Protection** - Front-running protection built-in
- 📊 **Real-time Price Feeds** - Always getting the alpha
- 🚀 **One-Click Deploy** - From zero to hero in minutes

## 🚀 Quick Start (Ape In Responsibly)

### Prerequisites 📋
- Go 1.21+ installed
- Foundry for smart contract magic
- X Layer testnet funds (we'll help you get them)
- A wallet with some testnet ETH

### 1. Clone & Setup 📥
```bash
git clone <repo-url>
cd sentinel-agent
chmod +x *.sh
```

### 2. Environment Configuration 🔧
Create your `.env` file:
```bash
cp .env.example .env
```

Edit `.env` with your details:
```env
X_LAYER_RPC=https://xlayerrpc.example.com
PRIVATE_KEY=your_64_character_hex_private_key
SMART_ACCOUNT=deployed_smart_account_address_here
```

### 3. Get Testnet Funds 💰
Your generated wallet needs some testnet tokens:
- 🌊 **X Layer Faucet**: https://www.okx.com/xlayer/faucet
- 🔗 **Bridge from Sepolia**: Get Sepolia ETH → Bridge to X Layer

### 4. Deploy & Run 🚀
```bash
# One command to rule them all
./setup-and-run.sh
```

This legendary script will:
1. ✅ Install all dependencies
2. 🔨 Compile smart contracts
3. 🚀 Deploy your Smart Account
4. 🤖 Start the Sentinel Agent

## 🏛️ Smart Contract Architecture

### SmartAccount.sol - The Foundation 🏗️
```solidity
contract SmartAccount {
    address public owner;
    mapping(address => uint256) public sessionKeys;
    
    function execute(address target, bytes calldata data) external;
    function setSessionKey(address key, uint256 expiresAt) external;
}
```

**Key Features:**
- 🔐 **Owner-based access control**
- ⏰ **Time-limited session keys** 
- 🎯 **Generic execution function**
- ⛽ **Gas-optimized operations**

## 🤖 Agent Capabilities

The Sentinel Agent is your 24/7 DeFi guardian that:

1. 📡 **Monitors Market Conditions**
   - Real-time price feeds from OKX DEX
   - Slippage protection mechanisms
   - Gas price optimization

2. 🎯 **Executes Smart Trades**
   - Automated swap execution
   - MEV protection strategies
   - Multi-hop routing when profitable

3. 🛡️ **Security First**
   - Session key rotation
   - Transaction simulation before execution
   - Fail-safe mechanisms

## 📊 Supported Operations

| Operation | Status | Description |
|-----------|--------|-------------|
| 💱 Token Swaps | ✅ | ETH ↔ ERC20, ERC20 ↔ ERC20 |
| 🔄 Auto-Rebalancing | 🚧 | Coming soon™ |
| 📈 DCA Strategies | 🚧 | Dollar cost averaging |
| 🌊 Liquidity Mining | 🚧 | Automated LP management |

## 🔧 Configuration Options

### Environment Variables
```env
# Required
X_LAYER_RPC=your_rpc_endpoint
PRIVATE_KEY=your_private_key
SMART_ACCOUNT=your_deployed_contract

# Optional
GAS_LIMIT=300000
MAX_SLIPPAGE=0.5
MIN_PROFIT_THRESHOLD=0.1
```

### Agent Parameters
```go
// Customize in agent/main.go
const (
    DefaultGasLimit = 300000
    MaxSlippage    = 0.005  // 0.5%
    PollInterval   = 30     // seconds
)
```

## 🚨 Security Best Practices

1. 🔐 **Private Key Management**
   - Never commit private keys to git
   - Use environment variables only
   - Consider hardware wallet integration

2. ⏰ **Session Key Rotation**
   - Set reasonable expiration times
   - Rotate keys regularly
   - Monitor unauthorized usage

3. 💰 **Fund Management**
   - Start with small amounts
   - Test on testnet first
   - Set reasonable limits

## 🛠️ Development & Testing

### Running Tests 🧪
```bash
# Smart contract tests
cd blockchain && forge test

# Agent tests (coming soon)
cd agent && go test ./...
```

### Local Development 💻
```bash
# Start local development
./run-agent.sh

# Deploy to different networks
NETWORK=mainnet ./deploy/deploy.sh
```

## 🌐 Supported Networks

| Network | Status | Chain ID | RPC |
|---------|--------|----------|-----|
| X Layer Testnet | ✅ | 195 | https://... |
| X Layer Mainnet | 🚧 | 196 | Coming Soon |

## 📚 Documentation Deep Dive

- 📖 **[Setup Guide](./SETUP.md)** - Detailed setup instructions
- 🏗️ **[Smart Contracts](./blockchain/README.md)** - Contract documentation
- 🤖 **[Agent API](./agent/README.md)** - Agent documentation

## 🚀 Roadmap to the Moon

- [x] ✅ Smart Account deployment
- [x] ✅ OKX DEX integration  
- [x] ✅ Basic swap functionality
- [ ] 🚧 Advanced trading strategies
- [ ] 🚧 Multi-chain support
- [ ] 🚧 Web dashboard
- [ ] 🚧 Mobile notifications
- [ ] 🚧 Social trading features

## 🤝 Contributing (Be a Chad)

We welcome all contributors! Whether you're a:
- 🧙‍♂️ Solidity wizard
- 🐹 Go gopher
- 📊 Data scientist
- 🎨 Frontend artist
- 📝 Documentation hero

Check out our [Contributing Guidelines](./CONTRIBUTING.md) to get started!

## ⚠️ Disclaimer (DYOR Always)

> **This software is experimental and for educational purposes. Always DYOR (Do Your Own Research) and never risk more than you can afford to lose. Past performance doesn't guarantee future results. Not financial advice. May contain traces of hopium.** 🚨

## 🆘 Support & Community

- 🐛 **Issues**: [GitHub Issues](../../issues)
- 💬 **Discussions**: [GitHub Discussions](../../discussions)  
- 🐦 **Twitter**: [@SentinelAgent](https://twitter.com/SentinelAgent)
- 📱 **Telegram**: [t.me/SentinelAgent](https://t.me/SentinelAgent)

## 📄 License

MIT License - Because sharing is caring! 💕

