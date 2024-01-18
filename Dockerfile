FROM golang:1.21-alpine 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy all files from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o geoproperty .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./geoproperty", "server", "-i", "0.0.0.0", "-p", "3000"]

