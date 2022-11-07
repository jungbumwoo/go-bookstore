FROM golang:latest

# Change working dir -> for locating configrations
RUN mkdir -p /app
WORKDIR /app

# COPY go.mod and go.sum files to the workspace
COPY go.mod ./
COPY go.sum ./

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Copy the local package files to the container's workspace.
COPY . .

RUN go build -o /go-bookstore

CMD ["/go-bookstore", "--host", "0.0.0.0"]