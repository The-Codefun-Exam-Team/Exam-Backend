# syntax=docker/dockerfile:1

### STAGE 1: Build the executable ###

# Use the latest version
FROM golang:1.19-alpine as build

# Create a directory inside the image
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code to the folder
COPY . .

# Compile to the executable
RUN go build -o ./main

### STAGE 2: Run the server ###

FROM alpine

WORKDIR /

COPY --from=build /app/main ./main

EXPOSE 80

# Run the executable when the image is run
ENTRYPOINT [ "./main" ]
