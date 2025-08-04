'use client'

import { useState, useEffect, useRef } from 'react'

export interface DashboardData {
  portfolio: {
    totalValue: string
    totalChange: string
    assets: Array<{
      symbol: string
      amount?: string
      value: string
      change: string
      address?: string
      chainId?: number
    }>
    distribution: Array<{
      chainId: number
      name: string
      value: string
      percentage: number
    }>
  }
  strategies: Array<{
    id: number
    name: string
    type: string
    status: string
    nextExecution: string
    performance: string
    invested: string
    currentValue: string
    lastExecution?: string
  }>
  transactions: Array<{
    hash: string
    type: string
    amount: string
    token: string
    status: 'pending' | 'completed' | 'failed'
    timestamp: string
    chain: string
    chainId?: number
    gasUsed?: string
    gasPrice?: string
    from?: string
    to?: string
  }>
  stats: {
    totalPortfolio: string
    activeStrategies: number
    volume24h: string
    gasOptimized: string
    totalTransactions: number
    successRate: string
    avgGasPrice: string
  }
  chains: Array<{
    chainId: number
    name: string
    status: string
    isActive: boolean
    lastChecked: string
    lastBlock?: number
    gasPrice?: string
  }>
  lastUpdated: string
}

interface WebSocketMessage {
  type: string
  data: any
}

export function useWebSocket(url: string) {
  const [isConnected, setIsConnected] = useState(false)
  const [data, setData] = useState<DashboardData | null>(null)
  const [error, setError] = useState<string | null>(null)
  const ws = useRef<WebSocket | null>(null)
  const reconnectTimeout = useRef<NodeJS.Timeout | null>(null)

  const connect = () => {
    try {
      ws.current = new WebSocket(url)

      ws.current.onopen = () => {
        setIsConnected(true)
        setError(null)
        console.log('🔌 WebSocket connected')
      }

      ws.current.onmessage = (event) => {
        try {
          const message: WebSocketMessage = JSON.parse(event.data)
          handleMessage(message)
        } catch (err) {
          console.error('Failed to parse WebSocket message:', err)
        }
      }

      ws.current.onclose = () => {
        setIsConnected(false)
        console.log('🔌 WebSocket disconnected')
        // Attempt to reconnect after 3 seconds
        reconnectTimeout.current = setTimeout(connect, 3000)
      }

      ws.current.onerror = (err) => {
        setError('WebSocket connection error')
        console.error('WebSocket error:', err)
      }
    } catch (err) {
      setError('Failed to connect to WebSocket')
      console.error('WebSocket connection failed:', err)
    }
  }

  const handleMessage = (message: WebSocketMessage) => {
    switch (message.type) {
      case 'initial_data':
      case 'data_update':
        setData(message.data)
        break
      case 'portfolio_update':
        setData(prev => prev ? { ...prev, portfolio: message.data } : null)
        break
      case 'strategies_update':
        setData(prev => prev ? { ...prev, strategies: message.data } : null)
        break
      case 'transactions_update':
        setData(prev => prev ? { ...prev, transactions: message.data } : null)
        break
    }
  }

  const sendMessage = (type: string, data?: any) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify({ type, data }))
    }
  }

  useEffect(() => {
    connect()

    return () => {
      if (reconnectTimeout.current) {
        clearTimeout(reconnectTimeout.current)
      }
      if (ws.current) {
        ws.current.close()
      }
    }
  }, [url])

  return {
    isConnected,
    data,
    error,
    sendMessage
  }
}

export function useApi(baseUrl: string) {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const fetchData = async (endpoint: string, options?: RequestInit) => {
    setLoading(true)
    setError(null)

    try {
      const response = await fetch(`${baseUrl}${endpoint}`, {
        headers: {
          'Content-Type': 'application/json',
          ...options?.headers
        },
        ...options
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      setLoading(false)
      return data
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Unknown error'
      setError(errorMessage)
      setLoading(false)
      throw err
    }
  }

  const get = (endpoint: string) => fetchData(endpoint)
  const post = (endpoint: string, data: any) => 
    fetchData(endpoint, { method: 'POST', body: JSON.stringify(data) })
  const put = (endpoint: string, data: any) => 
    fetchData(endpoint, { method: 'PUT', body: JSON.stringify(data) })
  const del = (endpoint: string) => 
    fetchData(endpoint, { method: 'DELETE' })

  return {
    loading,
    error,
    get,
    post,
    put,
    del
  }
}

export function useLocalStorage<T>(key: string, initialValue: T) {
  const [storedValue, setStoredValue] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(key)
      return item ? JSON.parse(item) : initialValue
    } catch (error) {
      console.error(`Error reading localStorage key "${key}":`, error)
      return initialValue
    }
  })

  const setValue = (value: T | ((val: T) => T)) => {
    try {
      const valueToStore = value instanceof Function ? value(storedValue) : value
      setStoredValue(valueToStore)
      window.localStorage.setItem(key, JSON.stringify(valueToStore))
    } catch (error) {
      console.error(`Error setting localStorage key "${key}":`, error)
    }
  }

  return [storedValue, setValue] as const
}

export function useNotifications() {
  const [notifications, setNotifications] = useState<Array<{
    id: string
    type: 'success' | 'error' | 'warning' | 'info'
    title: string
    message: string
    timestamp: Date
  }>>([])

  const addNotification = (type: 'success' | 'error' | 'warning' | 'info', title: string, message: string) => {
    const notification = {
      id: Date.now().toString(),
      type,
      title,
      message,
      timestamp: new Date()
    }

    setNotifications(prev => [notification, ...prev.slice(0, 9)]) // Keep last 10

    // Auto-remove after 5 seconds for success/info, 10 seconds for warnings/errors
    const timeout = type === 'error' || type === 'warning' ? 10000 : 5000
    setTimeout(() => {
      removeNotification(notification.id)
    }, timeout)
  }

  const removeNotification = (id: string) => {
    setNotifications(prev => prev.filter(n => n.id !== id))
  }

  const clearAll = () => {
    setNotifications([])
  }

  return {
    notifications,
    addNotification,
    removeNotification,
    clearAll
  }
}

export function useInterval(callback: () => void, delay: number | null) {
  const savedCallback = useRef<() => void>()

  useEffect(() => {
    savedCallback.current = callback
  }, [callback])

  useEffect(() => {
    function tick() {
      savedCallback.current?.()
    }

    if (delay !== null) {
      const id = setInterval(tick, delay)
      return () => clearInterval(id)
    }
  }, [delay])
}

export function useDebounce<T>(value: T, delay: number): T {
  const [debouncedValue, setDebouncedValue] = useState<T>(value)

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value)
    }, delay)

    return () => {
      clearTimeout(handler)
    }
  }, [value, delay])

  return debouncedValue
}
