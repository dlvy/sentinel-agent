#!/bin/bash

# 🔄 Sentinel Agent Web Dashboard Restart Script
# This script stops any running dashboard processes and restarts them

set -e

echo "🔄 Restarting Sentinel Agent Web Dashboard..."
echo "==============================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

# Function to print colored output
print_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in the correct directory
if [ ! -f "package.json" ]; then
    print_error "Please run this script from the web-dashboard directory"
    exit 1
fi

print_step "Stopping existing dashboard processes..."

# Kill existing dashboard-server processes
DASHBOARD_PIDS=$(ps aux | grep "dashboard-server" | grep -v grep | awk '{print $2}')
if [ ! -z "$DASHBOARD_PIDS" ]; then
    echo "$DASHBOARD_PIDS" | xargs kill -9 2>/dev/null || true
    print_success "Stopped dashboard server processes"
else
    print_warning "No dashboard server processes found"
fi

# Kill existing npm/node processes related to our dashboard
NPM_PIDS=$(ps aux | grep -E "(npm.*dev|next-server|next dev)" | grep -v grep | awk '{print $2}')
if [ ! -z "$NPM_PIDS" ]; then
    echo "$NPM_PIDS" | xargs kill -9 2>/dev/null || true
    print_success "Stopped frontend processes"
else
    print_warning "No frontend processes found"
fi

# Wait a moment for processes to stop
sleep 2

print_step "Cleaning up build artifacts..."
# Remove any existing build server
if [ -f "server/dashboard-server" ]; then
    rm server/dashboard-server
    print_success "Removed old server binary"
fi

# Clean Next.js cache if needed
if [ -d ".next" ]; then
    rm -rf .next
    print_success "Cleaned Next.js cache"
fi

print_step "Starting fresh dashboard..."
# Start the main dashboard script
./start-dashboard.sh
