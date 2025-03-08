# Use a lightweight Go image
FROM golang:1.24.0 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files to the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final container
FROM alpine:latest  

WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
