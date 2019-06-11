# Create a builder image for installation and build tasks 
FROM golang:alpine AS builder

# The go command disregards the GOPATH and requires the use of modules instead.
ENV GO111MODULE=on

# Set the working directory and add the go source code
WORKDIR /app
ADD ./ /app

# Run a OS update, install git and compile the golang-test binary  
RUN apk update --no-cache && \
    apk add git --no-cache && \
    go build -o api-test .

# Build the 
FROM alpine

ARG version
ARG last_commit_sha 

# Set the working directory 
WORKDIR /app

# Copy the the golang-test binary over from the builder image
COPY --from=builder /app/api-test .

ENV VERSION=$version
ENV LAST_COMMIT_SHA=$last_commit_sha

# Set the container start the golang-binary to the 
ENTRYPOINT ["/app/api-test"]

# Open port 8000 to the container
EXPOSE 8000

