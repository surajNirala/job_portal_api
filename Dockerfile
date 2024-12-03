########## Stage 1: #################

# Build
FROM golang:1.23-alpine AS builder


# Set the working directory in builder stage
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the builder container
COPY . .

# Build the Go application 
RUN go build -o main ./cmd/server/main.go



############ Stage 2: #################

# Run
FROM alpine:3.18

# Set working directory in the final stage
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]




