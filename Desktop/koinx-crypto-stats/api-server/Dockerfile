FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o api-server ./cmd/api

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./api-server"] 