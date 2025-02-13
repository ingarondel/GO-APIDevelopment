FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o go-apidevelopment ./cmd/main.go

FROM golang:1.21

WORKDIR /root/

COPY --from=builder /app/go-apidevelopment .
COPY --from=builder /app/.env .  

EXPOSE 3000
CMD ["./go-apidevelopment"]
