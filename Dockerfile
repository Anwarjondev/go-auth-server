    FROM golang:1.24 AS builder

    WORKDIR /app

    # Install SQLite dependencies
    RUN apt-get update && apt-get install -y gcc libc6-dev

    # Copy source code
    COPY . .

    # Enable CGO and build the binary
    ENV CGO_ENABLED=1
    RUN go build -o go-auth-server .

    # Create a lightweight final image
    FROM debian:bookworm-slim

    WORKDIR /app

    COPY --from=builder /app/go-auth-server .

    EXPOSE 8080

    CMD ["./go-auth-server"]
