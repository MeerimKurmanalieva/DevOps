
# Use an official Go runtime as the base image
FROM golang:1.16

# Set the working directory in the container
WORKDIR /app

# Initialize Go module
RUN go mod init example.com/myapp

# Copy the application source code to the container
COPY .. .

# Build the Go application
RUN go build -o main .

# Set the entry point command to run the application
CMD ["./main"]