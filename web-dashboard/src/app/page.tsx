'use client'

import React, { useState, useEffect } from 'react'
import { Sidebar, TopBar, StatCard, TransactionRow, LoadingSpinner } from '@/components/ui'
import {
  BarChart3,
  DollarSign,
  TrendingUp,
  Zap,
  Activity,
  Wallet,
  Bot,
  Globe,
  RefreshCw,
  ArrowUpDown,
  LineChart,
  PieChart,
  Settings,
  AlertTriangle,
  CheckCircle,
  Clock,
} from 'lucide-react'

// Mock data - In a real app, this would come from APIs
const mockPortfolioData = {
  totalValue: '$25,420.50',
  totalChange: '+12.5%',
  chains: [
    { name: 'X Layer', value: '$18,250.30', percentage: 72 },
    { name: 'Ethereum', value: '$4,120.10', percentage: 16 },
    { name: 'Polygon', value: '$2,050.10', percentage: 8 },
    { name: 'Arbitrum', value: '$1,000.00', percentage: 4 },
  ],
  assets: [
    { symbol: 'ETH', amount: '8.5', value: '$20,400.00', change: '+5.2%' },
    { symbol: 'USDC', value: '$3,000.00', change: '+0.1%' },
    { symbol: 'OKB', amount: '125', value: '$2,020.50', change: '+8.7%' },
  ]
}

const mockTransactions = [
  {
    hash: '0x1a2b3c4d5e6f7890abcdef1234567890abcdef12',
    type: 'DCA Buy',
    amount: '0.1',
    token: 'ETH',
    status: 'completed' as const,
    timestamp: '2 mins ago',
    chain: 'X Layer'
  },
  {
    hash: '0x2b3c4d5e6f7890abcdef1234567890abcdef123a',
    type: 'Grid Order',
    amount: '500',
    token: 'USDC',
    status: 'pending' as const,
    timestamp: '5 mins ago',
    chain: 'X Layer'
  },
  {
    hash: '0x3c4d5e6f7890abcdef1234567890abcdef123a2b',
    type: 'Rebalance',
    amount: '1.2',
    token: 'ETH',
    status: 'completed' as const,
    timestamp: '1 hour ago',
    chain: 'Polygon'
  },
]

const mockStrategies = [
  {
    id: 1,
    name: 'ETH DCA Strategy',
    type: 'DCA',
    status: 'active',
    nextExecution: '23 mins',
    performance: '+15.3%',
    invested: '$2,400',
    current: '$2,767'
  },
  {
    id: 2,
    name: 'Grid Trading Bot',
    type: 'Grid',
    status: 'active',
    nextExecution: '12 mins',
    performance: '+8.7%',
    invested: '$5,000',
    current: '$5,435'
  },
  {
    id: 3,
    name: 'Portfolio Rebalancer',
    type: 'Rebalance',
    status: 'paused',
    nextExecution: 'Paused',
    performance: '+3.2%',
    invested: '$10,000',
    current: '$10,320'
  },
]

function OverviewTab() {
  return (
    <div className="space-y-6">
      {/* Stats Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <StatCard
          title="Total Portfolio"
          value={mockPortfolioData.totalValue}
          change={mockPortfolioData.totalChange}
          changeType="positive"
          icon={Wallet}
          description="Across 7 chains"
        />
        <StatCard
          title="Active Strategies"
          value="3"
          change="2 executing"
          changeType="positive"
          icon={Bot}
          description="DCA • Grid • Rebalance"
        />
        <StatCard
          title="24h Volume"
          value="$8,420"
          change="+23.1%"
          changeType="positive"
          icon={BarChart3}
          description="Automated trades"
        />
        <StatCard
          title="Gas Optimized"
          value="$127"
          change="Saved today"
          changeType="positive"
          icon={Zap}
          description="Multi-chain routing"
        />
      </div>

      {/* Charts Row */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Portfolio Distribution */}
        <div className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
          <h3 className="text-lg font-semibold text-white mb-4">Portfolio Distribution</h3>
          <div className="space-y-3">
            {mockPortfolioData.chains.map((chain, index) => (
              <div key={index} className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-3 h-3 rounded-full bg-gradient-to-r from-blue-400 to-purple-500"></div>
                  <span className="text-white">{chain.name}</span>
                </div>
                <div className="text-right">
                  <div className="text-white font-medium">{chain.value}</div>
                  <div className="text-white/60 text-sm">{chain.percentage}%</div>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Recent Activity */}
        <div className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-lg font-semibold text-white">Recent Activity</h3>
            <button className="text-blue-400 hover:text-blue-300 transition-colors">
              <RefreshCw className="w-4 h-4" />
            </button>
          </div>
          <div className="space-y-3">
            {mockTransactions.slice(0, 3).map((tx, index) => (
              <TransactionRow key={index} {...tx} />
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

function PortfolioTab() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-white">Multi-Chain Portfolio</h2>
        <button className="flex items-center space-x-2 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg transition-colors">
          <RefreshCw className="w-4 h-4" />
          <span>Refresh</span>
        </button>
      </div>

      {/* Portfolio Summary */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2 bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
          <h3 className="text-lg font-semibold text-white mb-4">Asset Breakdown</h3>
          <div className="space-y-4">
            {mockPortfolioData.assets.map((asset, index) => (
              <div key={index} className="flex items-center justify-between p-4 bg-white/5 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-orange-400 to-red-500 rounded-full flex items-center justify-center">
                    <span className="text-white font-bold text-sm">{asset.symbol}</span>
                  </div>
                  <div>
                    <div className="text-white font-medium">{asset.symbol}</div>
                    {asset.amount && (
                      <div className="text-white/60 text-sm">{asset.amount} {asset.symbol}</div>
                    )}
                  </div>
                </div>
                <div className="text-right">
                  <div className="text-white font-medium">{asset.value}</div>
                  <div className="text-green-400 text-sm">{asset.change}</div>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
          <h3 className="text-lg font-semibold text-white mb-4">Chain Distribution</h3>
          <div className="space-y-3">
            {mockPortfolioData.chains.map((chain, index) => (
              <div key={index} className="space-y-2">
                <div className="flex justify-between text-sm">
                  <span className="text-white">{chain.name}</span>
                  <span className="text-white/70">{chain.percentage}%</span>
                </div>
                <div className="w-full bg-white/20 rounded-full h-2">
                  <div 
                    className="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full"
                    style={{ width: `${chain.percentage}%` }}
                  ></div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

function StrategiesTab() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-white">Trading Strategies</h2>
        <button className="flex items-center space-x-2 bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg transition-colors">
          <Bot className="w-4 h-4" />
          <span>New Strategy</span>
        </button>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-6">
        {mockStrategies.map((strategy) => (
          <div key={strategy.id} className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg font-semibold text-white">{strategy.name}</h3>
              <span className={`px-2 py-1 text-xs rounded-full ${
                strategy.status === 'active' 
                  ? 'bg-green-500/20 text-green-400 border border-green-500/30' 
                  : 'bg-yellow-500/20 text-yellow-400 border border-yellow-500/30'
              }`}>
                {strategy.status}
              </span>
            </div>
            
            <div className="space-y-3">
              <div className="flex justify-between">
                <span className="text-white/70">Type</span>
                <span className="text-white">{strategy.type}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-white/70">Performance</span>
                <span className="text-green-400">{strategy.performance}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-white/70">Invested</span>
                <span className="text-white">{strategy.invested}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-white/70">Current Value</span>
                <span className="text-white font-medium">{strategy.current}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-white/70">Next Execution</span>
                <span className="text-blue-400">{strategy.nextExecution}</span>
              </div>
            </div>

            <div className="mt-4 pt-4 border-t border-white/20">
              <div className="flex space-x-2">
                <button className="flex-1 bg-blue-500/20 hover:bg-blue-500/30 text-blue-400 border border-blue-500/30 px-3 py-2 rounded-lg text-sm transition-colors">
                  Edit
                </button>
                <button className="flex-1 bg-red-500/20 hover:bg-red-500/30 text-red-400 border border-red-500/30 px-3 py-2 rounded-lg text-sm transition-colors">
                  Pause
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

function TransactionsTab() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-white">Transaction History</h2>
        <div className="flex items-center space-x-2">
          <select className="bg-white/10 border border-white/20 text-white rounded-lg px-3 py-2 text-sm">
            <option>All Chains</option>
            <option>X Layer</option>
            <option>Ethereum</option>
            <option>Polygon</option>
          </select>
          <button className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg text-sm transition-colors">
            Export
          </button>
        </div>
      </div>

      <div className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6">
        <div className="space-y-4">
          {mockTransactions.map((tx, index) => (
            <TransactionRow key={index} {...tx} />
          ))}
        </div>
      </div>
    </div>
  )
}

export default function Dashboard() {
  const [activeTab, setActiveTab] = useState('overview')
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false)
  const [isLoading, setIsLoading] = useState(false)

  const renderActiveTab = () => {
    switch (activeTab) {
      case 'overview':
        return <OverviewTab />
      case 'portfolio':
        return <PortfolioTab />
      case 'strategies':
        return <StrategiesTab />
      case 'transactions':
        return <TransactionsTab />
      default:
        return <div className="text-white text-center py-12">
          <h2 className="text-2xl font-bold mb-2">Coming Soon</h2>
          <p className="text-white/70">This feature is under development</p>
        </div>
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <div className="flex h-screen">
        {/* Mobile Overlay */}
        {isMobileMenuOpen && (
          <div 
            className="fixed inset-0 bg-black/50 z-40 lg:hidden"
            onClick={() => setIsMobileMenuOpen(false)}
          />
        )}
        
        {/* Sidebar */}
        <Sidebar
          activeTab={activeTab}
          setActiveTab={setActiveTab}
          isMobile={true}
          isOpen={isMobileMenuOpen}
          onClose={() => setIsMobileMenuOpen(false)}
        />
        
        {/* Desktop Sidebar */}
        <div className="hidden lg:block">
          <Sidebar
            activeTab={activeTab}
            setActiveTab={setActiveTab}
            isMobile={false}
          />
        </div>

        {/* Main Content */}
        <div className="flex-1 flex flex-col overflow-hidden">
          <TopBar onMenuClick={() => setIsMobileMenuOpen(true)} />
          
          <main className="flex-1 overflow-y-auto p-6 bg-hero-pattern">
            {isLoading ? (
              <LoadingSpinner />
            ) : (
              renderActiveTab()
            )}
          </main>
        </div>
      </div>
    </div>
  )
}
