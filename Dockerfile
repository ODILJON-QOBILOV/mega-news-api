# Use official Go image with CGO support
FROM golang:1.21

# Install dependencies for sqlite3 C library
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

WORKDIR /app

# Copy go.mod and go.sum and download deps first (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Enable CGO and build the app
ENV CGO_ENABLED=1
RUN go build -o app .

# Run the app
CMD ["./app"]
