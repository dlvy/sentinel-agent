// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract SmartAccountV2 is ReentrancyGuard {
    using ECDSA for bytes32;

    address public owner;
    mapping(address => uint256) public sessionKeys;
    
    // Advanced Trading Strategies Storage
    struct DCAStrategy {
        bool active;
        address tokenIn;
        address tokenOut;
        uint256 amountPerExecution;
        uint256 interval; // seconds between executions
        uint256 lastExecution;
        uint256 totalExecutions;
        uint256 maxExecutions;
    }
    
    struct GridStrategy {
        bool active;
        address tokenA;
        address tokenB;
        uint256 gridSize;
        uint256 priceStep;
        uint256 basePrice;
        mapping(uint256 => bool) gridLevels; // price level => has order
    }
    
    struct RebalanceStrategy {
        bool active;
        address[] tokens;
        uint256[] targetPercentages; // basis points (10000 = 100%)
        uint256 rebalanceThreshold; // percentage deviation to trigger rebalance
        uint256 lastRebalance;
        uint256 minInterval; // minimum time between rebalances
    }

    mapping(uint256 => DCAStrategy) public dcaStrategies;
    mapping(uint256 => GridStrategy) public gridStrategies;
    mapping(uint256 => RebalanceStrategy) public rebalanceStrategies;
    
    uint256 public nextStrategyId = 1;
    
    // Multi-chain support
    mapping(uint256 => bool) public supportedChains;
    mapping(uint256 => address) public crossChainContracts;
    
    // Events
    event StrategyCreated(uint256 indexed strategyId, string strategyType);
    event StrategyExecuted(uint256 indexed strategyId, address indexed executor);
    event StrategyPaused(uint256 indexed strategyId);
    event ChainAdded(uint256 indexed chainId, address contractAddress);

    constructor(address _owner) {
        owner = _owner;
        // Add initial supported chains
        supportedChains[1] = true;     // Ethereum Mainnet
        supportedChains[137] = true;   // Polygon
        supportedChains[42161] = true; // Arbitrum
        supportedChains[10] = true;    // Optimism
        supportedChains[8453] = true;  // Base
        supportedChains[195] = true;   // X Layer Testnet
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }
    
    modifier onlyAuthorized() {
        require(
            msg.sender == owner || sessionKeys[msg.sender] > block.timestamp,
            "Unauthorized"
        );
        _;
    }

    function setSessionKey(address key, uint256 expiresAt) external onlyOwner {
        sessionKeys[key] = expiresAt;
    }

    function execute(address target, bytes calldata data) external onlyAuthorized nonReentrant {
        (bool success, ) = target.call(data);
        require(success, "Call failed");
    }
    
    // === DCA Strategy Management ===
    
    function createDCAStrategy(
        address tokenIn,
        address tokenOut,
        uint256 amountPerExecution,
        uint256 interval,
        uint256 maxExecutions
    ) external onlyOwner returns (uint256 strategyId) {
        strategyId = nextStrategyId++;
        
        DCAStrategy storage strategy = dcaStrategies[strategyId];
        strategy.active = true;
        strategy.tokenIn = tokenIn;
        strategy.tokenOut = tokenOut;
        strategy.amountPerExecution = amountPerExecution;
        strategy.interval = interval;
        strategy.lastExecution = block.timestamp;
        strategy.maxExecutions = maxExecutions;
        
        emit StrategyCreated(strategyId, "DCA");
    }
    
    function executeDCA(uint256 strategyId) external onlyAuthorized {
        DCAStrategy storage strategy = dcaStrategies[strategyId];
        require(strategy.active, "Strategy not active");
        require(block.timestamp >= strategy.lastExecution + strategy.interval, "Too early");
        require(strategy.totalExecutions < strategy.maxExecutions, "Max executions reached");
        
        strategy.lastExecution = block.timestamp;
        strategy.totalExecutions++;
        
        if (strategy.totalExecutions >= strategy.maxExecutions) {
            strategy.active = false;
        }
        
        emit StrategyExecuted(strategyId, msg.sender);
    }
    
    // === Grid Trading Strategy ===
    
    function createGridStrategy(
        address tokenA,
        address tokenB,
        uint256 gridSize,
        uint256 priceStep,
        uint256 basePrice
    ) external onlyOwner returns (uint256 strategyId) {
        strategyId = nextStrategyId++;
        
        GridStrategy storage strategy = gridStrategies[strategyId];
        strategy.active = true;
        strategy.tokenA = tokenA;
        strategy.tokenB = tokenB;
        strategy.gridSize = gridSize;
        strategy.priceStep = priceStep;
        strategy.basePrice = basePrice;
        
        emit StrategyCreated(strategyId, "Grid");
    }
    
    // === Rebalancing Strategy ===
    
    function createRebalanceStrategy(
        address[] calldata tokens,
        uint256[] calldata targetPercentages,
        uint256 rebalanceThreshold,
        uint256 minInterval
    ) external onlyOwner returns (uint256 strategyId) {
        require(tokens.length == targetPercentages.length, "Array length mismatch");
        
        uint256 totalPercentage = 0;
        for (uint256 i = 0; i < targetPercentages.length; i++) {
            totalPercentage += targetPercentages[i];
        }
        require(totalPercentage == 10000, "Percentages must sum to 100%");
        
        strategyId = nextStrategyId++;
        
        RebalanceStrategy storage strategy = rebalanceStrategies[strategyId];
        strategy.active = true;
        strategy.tokens = tokens;
        strategy.targetPercentages = targetPercentages;
        strategy.rebalanceThreshold = rebalanceThreshold;
        strategy.minInterval = minInterval;
        strategy.lastRebalance = block.timestamp;
        
        emit StrategyCreated(strategyId, "Rebalance");
    }
    
    function executeRebalance(uint256 strategyId) external onlyAuthorized {
        RebalanceStrategy storage strategy = rebalanceStrategies[strategyId];
        require(strategy.active, "Strategy not active");
        require(block.timestamp >= strategy.lastRebalance + strategy.minInterval, "Too early");
        
        strategy.lastRebalance = block.timestamp;
        emit StrategyExecuted(strategyId, msg.sender);
    }
    
    // === Multi-chain Support ===
    
    function addSupportedChain(uint256 chainId, address contractAddress) external onlyOwner {
        supportedChains[chainId] = true;
        crossChainContracts[chainId] = contractAddress;
        emit ChainAdded(chainId, contractAddress);
    }
    
    function removeSupportedChain(uint256 chainId) external onlyOwner {
        supportedChains[chainId] = false;
        delete crossChainContracts[chainId];
    }
    
    // === Strategy Management ===
    
    function pauseStrategy(uint256 strategyId, string calldata strategyType) external onlyOwner {
        bytes32 typeHash = keccak256(abi.encodePacked(strategyType));
        
        if (typeHash == keccak256("DCA")) {
            dcaStrategies[strategyId].active = false;
        } else if (typeHash == keccak256("Grid")) {
            gridStrategies[strategyId].active = false;
        } else if (typeHash == keccak256("Rebalance")) {
            rebalanceStrategies[strategyId].active = false;
        }
        
        emit StrategyPaused(strategyId);
    }
    
    // === View Functions ===
    
    function getDCAStrategy(uint256 strategyId) external view returns (DCAStrategy memory) {
        return dcaStrategies[strategyId];
    }
    
    function isChainSupported(uint256 chainId) external view returns (bool) {
        return supportedChains[chainId];
    }
    
    function getCrossChainContract(uint256 chainId) external view returns (address) {
        return crossChainContracts[chainId];
    }
}
