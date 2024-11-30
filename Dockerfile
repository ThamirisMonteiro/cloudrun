FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o cloudrun ./main.go

FROM alpine:latest
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/cloudrun /cloudrun
ENTRYPOINT ["/cloudrun"]
