# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Stage 2: Run
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

RUN chmod +x ./server

EXPOSE 5000

CMD ["./server"]
