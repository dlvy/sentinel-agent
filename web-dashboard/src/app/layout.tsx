import type { Metadata, Viewport } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Sentinel Agent Dashboard | Your DeFi Guardian on X Layer',
  description: 'Advanced trading dashboard for Sentinel Agent with multi-chain portfolio tracking, DCA strategies, and real-time monitoring.',
  keywords: 'DeFi, Trading, Blockchain, Dashboard, X Layer, OKX, Automated Trading',
  authors: [{ name: 'Sentinel Agent Team' }],
  metadataBase: new URL('https://sentinel-agent.app'),
  openGraph: {
    title: 'Sentinel Agent Dashboard',
    description: 'Your DeFi Guardian on X Layer - Advanced Trading & Multi-Chain Support',
    type: 'website',
    images: [
      {
        url: '/og-image.png',
        width: 1200,
        height: 630,
        alt: 'Sentinel Agent Dashboard',
      },
    ],
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Sentinel Agent Dashboard',
    description: 'Your DeFi Guardian on X Layer',
    images: ['/og-image.png'],
  },
  robots: {
    index: true,
    follow: true,
  },
}

export const viewport: Viewport = {
  width: 'device-width',
  initialScale: 1,
  maximumScale: 1,
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className="h-full">
      <body className="h-full font-sans antialiased bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
        <div className="min-h-full">
          {children}
        </div>
      </body>
    </html>
  )
}
