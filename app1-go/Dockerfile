FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o servidor

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/servidor .
EXPOSE 5000
CMD ["./servidor"]
