# FROM golang:1.18.5-bullseye

# # Change working dir -> for locating configrations
# RUN mkdir -p /app
# WORKDIR /app

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# # Copy the local package files to the container's workspace.
# COPY . .

# CMD /go/bin/go-bookstore


################################################################

#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git


# Change working dir -> for locating configrations
RUN mkdir -p /app
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates


EXPOSE 3000
