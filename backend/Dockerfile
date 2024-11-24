# Base image
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o harbor-view ./cmd/main.go

# Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set working directory in final image
WORKDIR /root/

# Copy built application and config file from builder
COPY --from=builder /app/harbor-view .
COPY config/config.yaml ./config/config.yaml

# Expose the backend port
EXPOSE 3000

# Command to run the application
CMD ["./harbor-view"]
