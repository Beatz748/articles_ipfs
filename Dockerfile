# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy application files into the container
COPY . .

# Build the application
RUN go build -o main .

# Set the environment variable PORT
ENV PORT=8080

# Declare the port to be exposed
EXPOSE $PORT

# Command to run the application
CMD ["./main"]

