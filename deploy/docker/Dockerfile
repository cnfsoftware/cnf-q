FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cnf-q-service ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cnf-q-service .

EXPOSE 8080

CMD ["./cnf-q-service"]