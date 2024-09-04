# Use official Golang image as base
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize Go modules (this will create go.mod and go.sum files)
RUN go mod init myapp

# Copy the source code
COPY . .

# Ensure go.mod and go.sum are in the right place and download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
