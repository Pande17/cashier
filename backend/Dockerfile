# Step 1: Build stage for Go application
FROM golang:1.24.2 AS builder

# Set working directory in the container
WORKDIR /app

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Step 2: Final stage with minimal image to run the application
FROM debian:bullseye

# Install necessary dependencies for running Go, MySQL client, and wkhtmltopdf
RUN apt-get update && apt-get install -y \
    mysql-client \
    wkhtmltopdf \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set working directory for the application in the final image
WORKDIR /root

# Copy the Go binary from the builder stage to the final image
COPY --from=builder /app/backend /root/

# Expose the port your Go application will run on (default: 3000)
EXPOSE 3000

# Set environment variables for Go application (replace with your actual values)
ENV DB_USER=root
ENV DB_PASS=""
ENV DB_HOST=mysql
ENV DB_PORT=3306
ENV DB_NAME=machine-cashier

# Command to run the Go application with air for hot-reloading
CMD ["air"]
