# Use an official Go runtime as the base image
FROM golang:1.16

# Set the working directory in the container
WORKDIR /app

# Copy the application source code to the container
COPY . .


# Build the Go application
RUN go build -o main .

# Set the entry point command to run the application
CMD ["./main"]
