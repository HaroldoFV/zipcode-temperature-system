FROM golang:1.23-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g internal/server/main.go -o ./docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./internal/server

RUN go test ./... -cover -v > test-report.txt

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/test-report.txt .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
CMD ["./main"]
