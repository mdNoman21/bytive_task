# Use a Go base image with a specific version
FROM golang:1.21.0-alpine3.18

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]

