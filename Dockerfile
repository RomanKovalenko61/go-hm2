FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

RUN apk update && apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

