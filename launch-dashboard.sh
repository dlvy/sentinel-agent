#!/bin/bash

# 🚀 Sentinel Agent - Complete Dashboard Launch Script
# This script provides easy access to launch the web dashboard

set -e

echo "🚀 Sentinel Agent - Dashboard Launcher"
echo "======================================"

# Colors for output
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}📊 Launching Web Dashboard...${NC}"
echo ""

# Check if web-dashboard directory exists
if [ ! -d "web-dashboard" ]; then
    echo -e "${YELLOW}⚠️  Web dashboard directory not found!${NC}"
    echo "Make sure you're in the correct project directory."
    exit 1
fi

# Navigate to web-dashboard and launch
cd web-dashboard

# Check if start script exists and is executable
if [ ! -f "start-dashboard.sh" ]; then
    echo -e "${YELLOW}⚠️  Dashboard startup script not found!${NC}"
    exit 1
fi

if [ ! -x "start-dashboard.sh" ]; then
    chmod +x start-dashboard.sh
fi

echo -e "${GREEN}🚀 Starting Sentinel Agent Dashboard...${NC}"
echo ""

# Execute the dashboard startup script
./start-dashboard.sh
