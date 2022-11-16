FROM golang:1.18-bullseye as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags "-w -s" -o gras cmd/gras/main.go

FROM debian:bullseye-slim as deploy
RUN apt-get update
COPY --from=builder /app/gras ./
CMD ["./gras", "server"]

FROM golang:1.18-bullseye as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
