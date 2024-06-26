# Start from the latest golang base image
FROM golang:1.22-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum .env ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./server ./server

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./server/cmd/ -o main .

####### Start a new stage from scratch #######
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage also copy the .env file
COPY --from=builder /app/server/cmd/main ./
COPY --from=builder /app/.env ./

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./main"]