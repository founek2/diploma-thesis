FROM golang:1.21 AS build
WORKDIR /go/src

# Install dependencies
COPY go.mod go.sum .
RUN go mod download

# Copy source code
COPY . .

# Run build
RUN go build cmd/server/main.go

FROM scratch AS runtime
COPY --from=build /go/src/main ./
EXPOSE 8080/tcp
ENTRYPOINT ["./main"]
