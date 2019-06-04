# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.12.5

# Add Maintainer Info
LABEL maintainer="Zuri Pabón <zuripabon@spotahome.com>"

# Set the Current Working Directory inside the container
WORKDIR /usr/src/

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

WORKDIR ./app

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 9092 to the outside world
EXPOSE 9092

# Run the executable
CMD ["go", "run", "main.go"]
