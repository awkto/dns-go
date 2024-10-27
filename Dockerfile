# Step 1: Build the Go app using the golang base image
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy Go modules first to leverage Docker's caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code and build the Go app
COPY . .
RUN go build -o dns-server main.go

# Step 2: Create a lightweight final image for running the app
FROM alpine:latest

# Optional: Use SQLite CLI for Debugging (Only If Needed)
# RUN apk add --no-cache sqlite

# Set the working directory
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/dns-server .

# Expose the DNS port
EXPOSE 8053/udp

# Run the DNS server binary when the container starts
CMD ["./dns-server"]


