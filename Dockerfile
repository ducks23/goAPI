FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates
ARG NUC_DB_VAR
ENV NUC_DB=${NUC_DB_VAR}

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port (adjust based on your app)
EXPOSE 3000

# Run the binary
CMD ["./main"]
