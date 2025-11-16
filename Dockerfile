# Use a multi-stage build for smaller final image
FROM golang:1.21-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /src

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o codebase-interface \
    ./cmd/codebase-interface

# Final stage
FROM scratch

# Copy ca-certificates from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary
COPY --from=builder /src/codebase-interface /usr/local/bin/codebase-interface

# Set the working directory
WORKDIR /workspace

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/codebase-interface"]

# Default command
CMD ["--help"]

# Add labels for better container metadata
LABEL org.opencontainers.image.title="Codebase Interface CLI"
LABEL org.opencontainers.image.description="A CLI tool for validating codebase structure and development standards"
LABEL org.opencontainers.image.vendor="Codebase Interface"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/codebase-interface/cli"