# Use a minimal base image for Go applications
FROM golang:1.17-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o myapp .

# Use a lightweight base image for the final container
FROM alpine:latest

# Copy the compiled binary from the builder stage to the final container
COPY --from=builder /app/myapp /usr/local/bin/myapp

# Set the entry point for the container
ENTRYPOINT ["myapp"]
