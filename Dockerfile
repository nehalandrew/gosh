# docker build -t gosh .
# docker run -it --rm gosh

# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Enable Go modules and set working directory
WORKDIR /app

# Install git (for modules), and build tools
RUN apk add --no-cache git

# Copy go mod files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go application with optimizations
RUN go build -ldflags="-s -w" -o gosh main.go

# Stage 2: Create minimal final image
FROM scratch

# Copy the Go binary from the builder
COPY --from=builder /app/gosh /gosh

# Set binary as entrypoint
ENTRYPOINT ["/gosh"]
