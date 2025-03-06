FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o worker .

FROM ubuntu:latest

# Set working directory
WORKDIR /root

# Copy built binary from builder
COPY --from=builder /app/worker ./

# Run the binary
CMD ["./worker"]


