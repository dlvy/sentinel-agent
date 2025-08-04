/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  trailingSlash: true,
  images: {
    unoptimized: true,
  },
  // Remove assetPrefix and basePath for development
  // assetPrefix: process.env.NODE_ENV === 'production' ? '/sentinel-agent/' : '',
  // basePath: process.env.NODE_ENV === 'production' ? '/sentinel-agent' : '',
  env: {
    CUSTOM_KEY: 'my-value',
  },
}

module.exports = nextConfig
