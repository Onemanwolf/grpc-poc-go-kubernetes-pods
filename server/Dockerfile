FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 50051

CMD ["./server"]