# Start with a base Go language image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files to the container
COPY main.go .




# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point command for the container
CMD ["./main"]
