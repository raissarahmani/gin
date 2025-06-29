#Build stage
FROM golang:1.24.4-alpine3.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd

#Rebuild
FROM alpine:3.21.3
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 80

CMD ["./server"]
