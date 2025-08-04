'use client'

import React, { useState, useEffect } from 'react'
import {
  Activity,
  BarChart3,
  Bot,
  DollarSign,
  Globe,
  Settings,
  Shield,
  TrendingUp,
  Wallet,
  Zap,
  AlertTriangle,
  CheckCircle,
  Clock,
  RefreshCw,
  ArrowUpDown,
  Eye,
  Filter,
  Search,
  Download,
  Bell,
  Menu,
  X,
} from 'lucide-react'

interface SidebarProps {
  activeTab: string
  setActiveTab: (tab: string) => void
  isMobile?: boolean
  isOpen?: boolean
  onClose?: () => void
}

interface NavigationItem {
  id: string
  label: string
  icon: React.ComponentType<any>
  description: string
}

const navigationItems: NavigationItem[] = [
  {
    id: 'overview',
    label: 'Overview',
    icon: BarChart3,
    description: 'Portfolio and performance overview'
  },
  {
    id: 'portfolio',
    label: 'Portfolio',
    icon: Wallet,
    description: 'Multi-chain asset tracking'
  },
  {
    id: 'strategies',
    label: 'Strategies',
    icon: Bot,
    description: 'Trading automation and DCA'
  },
  {
    id: 'transactions',
    label: 'Transactions',
    icon: ArrowUpDown,
    description: 'Transaction history and status'
  },
  {
    id: 'analytics',
    label: 'Analytics',
    icon: TrendingUp,
    description: 'Performance metrics and insights'
  },
  {
    id: 'multichain',
    label: 'Multi-Chain',
    icon: Globe,
    description: 'Cross-chain operations'
  },
  {
    id: 'settings',
    label: 'Settings',
    icon: Settings,
    description: 'Configuration and preferences'
  },
]

export function Sidebar({ activeTab, setActiveTab, isMobile = false, isOpen = true, onClose }: SidebarProps) {
  const [isCollapsed, setIsCollapsed] = useState(false)

  const sidebarClasses = `
    ${isMobile ? 'fixed inset-y-0 left-0 z-50' : 'relative'}
    ${isMobile && !isOpen ? '-translate-x-full' : 'translate-x-0'}
    ${isCollapsed && !isMobile ? 'w-16' : 'w-64'}
    transition-all duration-300 ease-in-out
    bg-white/10 backdrop-blur-xl border-r border-white/20
    flex flex-col h-full
  `

  return (
    <div className={sidebarClasses}>
      {/* Header */}
      <div className="flex items-center justify-between p-4 border-b border-white/20">
        {!isCollapsed && (
          <div className="flex items-center space-x-3">
            <div className="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
              <Shield className="w-6 h-6 text-white" />
            </div>
            <div>
              <h1 className="text-lg font-bold text-white">Sentinel</h1>
              <p className="text-xs text-blue-200">DeFi Guardian</p>
            </div>
          </div>
        )}
        
        <div className="flex items-center space-x-2">
          {!isMobile && (
            <button
              onClick={() => setIsCollapsed(!isCollapsed)}
              className="p-2 text-white/70 hover:text-white hover:bg-white/10 rounded-lg transition-colors"
            >
              <Menu className="w-4 h-4" />
            </button>
          )}
          
          {isMobile && onClose && (
            <button
              onClick={onClose}
              className="p-2 text-white/70 hover:text-white hover:bg-white/10 rounded-lg transition-colors"
            >
              <X className="w-4 h-4" />
            </button>
          )}
        </div>
      </div>

      {/* Navigation */}
      <nav className="flex-1 p-4 space-y-2 overflow-y-auto">
        {navigationItems.map((item) => {
          const isActive = activeTab === item.id
          const Icon = item.icon

          return (
            <button
              key={item.id}
              onClick={() => {
                setActiveTab(item.id)
                if (isMobile && onClose) onClose()
              }}
              className={`
                w-full flex items-center space-x-3 p-3 rounded-lg transition-all duration-200
                ${isActive 
                  ? 'bg-gradient-to-r from-blue-500/20 to-purple-500/20 border border-blue-400/30 text-white shadow-lg' 
                  : 'text-white/70 hover:text-white hover:bg-white/10'
                }
                ${isCollapsed ? 'justify-center' : ''}
              `}
              title={isCollapsed ? item.label : undefined}
            >
              <Icon className={`w-5 h-5 ${isActive ? 'text-blue-300' : ''}`} />
              {!isCollapsed && (
                <div className="flex-1 text-left">
                  <div className="text-sm font-medium">{item.label}</div>
                  {!isActive && (
                    <div className="text-xs text-white/50">{item.description}</div>
                  )}
                </div>
              )}
            </button>
          )
        })}
      </nav>

      {/* Footer */}
      {!isCollapsed && (
        <div className="p-4 border-t border-white/20">
          <div className="bg-gradient-to-r from-green-500/20 to-blue-500/20 border border-green-400/30 rounded-lg p-3">
            <div className="flex items-center space-x-2">
              <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
              <span className="text-sm text-white">Agent Active</span>
            </div>
            <p className="text-xs text-green-200 mt-1">
              Monitoring 7 chains • 3 active strategies
            </p>
          </div>
        </div>
      )}
    </div>
  )
}

interface TopBarProps {
  onMenuClick?: () => void
}

export function TopBar({ onMenuClick }: TopBarProps) {
  const [notifications, setNotifications] = useState(3)

  return (
    <header className="bg-white/10 backdrop-blur-xl border-b border-white/20 px-4 py-3">
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <button
            onClick={onMenuClick}
            className="lg:hidden p-2 text-white/70 hover:text-white hover:bg-white/10 rounded-lg transition-colors"
          >
            <Menu className="w-5 h-5" />
          </button>
          
          <div className="hidden sm:block">
            <h2 className="text-xl font-semibold text-white">Dashboard</h2>
            <p className="text-sm text-white/70">Welcome back, Agent Commander</p>
          </div>
        </div>

        <div className="flex items-center space-x-4">
          {/* Search */}
          <div className="hidden md:flex items-center space-x-2 bg-white/10 rounded-lg px-3 py-2">
            <Search className="w-4 h-4 text-white/50" />
            <input
              type="text"
              placeholder="Search transactions..."
              className="bg-transparent text-white placeholder-white/50 text-sm focus:outline-none w-48"
            />
          </div>

          {/* Notifications */}
          <button className="relative p-2 text-white/70 hover:text-white hover:bg-white/10 rounded-lg transition-colors">
            <Bell className="w-5 h-5" />
            {notifications > 0 && (
              <span className="absolute -top-1 -right-1 w-5 h-5 bg-red-500 text-white text-xs rounded-full flex items-center justify-center">
                {notifications}
              </span>
            )}
          </button>

          {/* Wallet Status */}
          <div className="hidden sm:flex items-center space-x-2 bg-green-500/20 border border-green-400/30 rounded-lg px-3 py-2">
            <div className="w-2 h-2 bg-green-400 rounded-full"></div>
            <span className="text-sm text-white">Connected</span>
          </div>
        </div>
      </div>
    </header>
  )
}

interface StatCardProps {
  title: string
  value: string
  change?: string
  changeType?: 'positive' | 'negative' | 'neutral'
  icon: React.ComponentType<any>
  description?: string
}

export function StatCard({ title, value, change, changeType = 'neutral', icon: Icon, description }: StatCardProps) {
  const changeColor = {
    positive: 'text-green-400',
    negative: 'text-red-400',
    neutral: 'text-yellow-400'
  }[changeType]

  return (
    <div className="bg-white/10 backdrop-blur-xl border border-white/20 rounded-xl p-6 hover:bg-white/15 transition-all duration-200">
      <div className="flex items-start justify-between">
        <div className="flex-1">
          <p className="text-sm text-white/70 font-medium">{title}</p>
          <p className="text-2xl font-bold text-white mt-1">{value}</p>
          {change && (
            <p className={`text-sm mt-1 ${changeColor}`}>
              {change}
            </p>
          )}
          {description && (
            <p className="text-xs text-white/50 mt-1">{description}</p>
          )}
        </div>
        <div className="w-12 h-12 bg-gradient-to-r from-blue-500/20 to-purple-500/20 rounded-lg flex items-center justify-center">
          <Icon className="w-6 h-6 text-blue-300" />
        </div>
      </div>
    </div>
  )
}

interface TransactionRowProps {
  hash: string
  type: string
  amount: string
  token: string
  status: 'pending' | 'completed' | 'failed'
  timestamp: string
  chain: string
}

export function TransactionRow({ hash, type, amount, token, status, timestamp, chain }: TransactionRowProps) {
  const statusConfig = {
    pending: { icon: Clock, color: 'text-yellow-400', bg: 'bg-yellow-500/20' },
    completed: { icon: CheckCircle, color: 'text-green-400', bg: 'bg-green-500/20' },
    failed: { icon: AlertTriangle, color: 'text-red-400', bg: 'bg-red-500/20' }
  }

  const config = statusConfig[status]
  const StatusIcon = config.icon

  return (
    <div className="bg-white/5 backdrop-blur-xl border border-white/10 rounded-lg p-4 hover:bg-white/10 transition-all duration-200">
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <div className={`w-10 h-10 ${config.bg} rounded-lg flex items-center justify-center`}>
            <StatusIcon className={`w-5 h-5 ${config.color}`} />
          </div>
          <div>
            <p className="text-white font-medium">{type}</p>
            <p className="text-sm text-white/70">{`${amount} ${token}`}</p>
            <p className="text-xs text-white/50">{chain}</p>
          </div>
        </div>
        
        <div className="text-right">
          <p className="text-sm text-white/70">{timestamp}</p>
          <p className="text-xs text-white/50 font-mono">{hash.slice(0, 8)}...{hash.slice(-6)}</p>
        </div>
      </div>
    </div>
  )
}

export function LoadingSpinner() {
  return (
    <div className="flex items-center justify-center p-8">
      <div className="w-8 h-8 border-4 border-blue-500/30 border-t-blue-500 rounded-full animate-spin"></div>
    </div>
  )
}
