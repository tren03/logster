# Use official Golang image as the builder
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /root/logster

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go binary
RUN go build -o logster

# Use a smaller image for the runtime stage
FROM alpine:latest

# Set working directory in the runtime container
WORKDIR /root

# Copy the compiled binary from the builder stage
COPY --from=builder /root/logster/logster .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./logster"]

