FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make gcc sqlite-dev musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o main .

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache sqlite

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]