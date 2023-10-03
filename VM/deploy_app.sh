#!/bin/bash

# Variables
SSH_USER="mkurmanalieva"
SSH_HOST="192.168.97.129"
SSH_PORT="2222"
# Define container details
CONTAINER_NAME="hello-web-app"
DOCKER_IMAGE="mkurmanalieva/hello-web-app:latest"
DOCKER_CONTAINER_ID="5762bfc6c9a3"
CONTAINER_PORT="8080"
HOST_PORT="8081"

# Customize the database run options
DOCKER_RUN_OPTIONS="-d -p $HOST_PORT:$CONTAINER_PORT"

# Connect to the Linux machine over SSH and run the database container
ssh -tt -p $SSH_PORT $SSH_USER@$SSH_HOST <<EOF
  # Stop and remove any existing container with the same name
  docker stop $CONTAINER_NAME || true
  docker rm $CONTAINER_NAME || true

  # Run the Docker container with customized options
  docker run $DOCKER_RUN_OPTIONS --name $CONTAINER_NAME $DOCKER_IMAGE
EOF
