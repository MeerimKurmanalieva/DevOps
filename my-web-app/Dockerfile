# Use a lightweight web server as the base image
FROM nginx:alpine

# Set the working directory in the container
WORKDIR /app

# Copy the application files to the web server directory
COPY index.html /usr/share/nginx/html


# Expose port 8080
EXPOSE 8080

# Start the web server
CMD ["nginx", "-g", "daemon off;"]
