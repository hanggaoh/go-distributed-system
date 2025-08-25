# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/node ./cmd/node

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/node .

ENV NODE_ID=0

CMD ["./node"]
