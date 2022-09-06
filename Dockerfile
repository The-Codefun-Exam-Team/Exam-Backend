# syntax=docker/dockerfile:1

# Use the latest version
FROM golang:1.19-bullseye

# Create a directory inside the image
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code to the folder
COPY . .

# Compile to the executable
RUN go build -o ./main

# Run the executable when the image is run
CMD [ "./main" ]
