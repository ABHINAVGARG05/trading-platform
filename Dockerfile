# Dockerfile

# Step 1: Use Golang image as the base
FROM golang:1.23-alpine AS builder

# Step 2: Install dependencies (git, air)
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Step 3: Set working directory in the container
WORKDIR /app

# Step 4: Copy Go mod and sum files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Step 5: Copy rest of the application code
COPY . .

# Expose port (match this with the one in your Go app)
EXPOSE 3000

# Start the application with Air for live reloading
CMD ["air"]
