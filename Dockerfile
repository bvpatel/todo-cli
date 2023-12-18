FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files for dependencies
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o todo-cli main.go

# Command to run the application
CMD ["/bin/sh"]
