FROM golang:1.21 AS build

# Install TrafficControl
RUN apt update && apt install -y iproute2 && apt clean

WORKDIR /go/src

# Install dependencies
COPY go.mod go.sum .
RUN go mod download

# Copy source code
COPY . .

# Run build
RUN go build cmd/server/main.go

EXPOSE 8080/tcp
ENTRYPOINT ["./main"]
