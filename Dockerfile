FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files into the workdir
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the workdir
COPY . .

# Build the application
RUN go build -o main cmd/server/main.go

# Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install ca-certificates and any other required libraries
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the application's listening port
EXPOSE 8080

# Run the application
CMD ["./main"]