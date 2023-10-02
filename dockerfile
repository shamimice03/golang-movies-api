# Stage 1: Build the Go application
FROM golang:1.21.1 AS build

WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Create the final Docker image
FROM golang:1.21.1-alpine3.18

# Copy the built binary from the "build" stage to the final stage
COPY --from=build /app/main .

# Expose port 8080 (optional)
EXPOSE 8080

# Command to run the Go application
CMD ["./main"] 
