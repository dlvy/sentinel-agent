# Multi-stage build for Sentinel Agent Dashboard

# Frontend build stage
FROM node:18-alpine AS frontend-builder
WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci --only=production

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Backend build stage
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy source code
COPY server/ .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dashboard-server main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Copy the frontend build
COPY --from=frontend-builder /app/out ./static/

# Copy the backend binary
COPY --from=backend-builder /app/dashboard-server .

# Change ownership to appuser
RUN chown -R appuser:appgroup /root/

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/api/v1/health || exit 1

# Start the application
CMD ["./dashboard-server"]
