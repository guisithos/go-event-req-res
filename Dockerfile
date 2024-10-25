# Start from the official Go image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
