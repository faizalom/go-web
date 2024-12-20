# syntax=docker/dockerfile:1
FROM golang:1.21.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in https://docs.docker.com/engine/reference/builder/
COPY *.go ./

# Build the Go app
RUN go build -o app

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./app"]
