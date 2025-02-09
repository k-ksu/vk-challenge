FROM golang:1.23 AS builder
WORKDIR /technotest
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN make migration-up
RUN go mod tidy
RUN go build -o technotest cmd/technotest/main.go
CMD ["./technotest"]
