FROM golang:latest

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum  ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
