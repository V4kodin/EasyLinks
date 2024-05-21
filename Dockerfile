# Start from the latest golang base image
FROM --platform=linux/amd64 golang:latest as builder

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

ENTRYPOINT ls -l ./server/cmd/main

####### Start a new stage from scratch #######

# Expose port 8080 to the outside
EXPOSE 8080

#RUN chmod 777 example.txt

#ENTRYPOINT ls -l ./

ENTRYPOINT ["./server/cmd/main"]