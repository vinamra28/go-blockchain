FROM golang:1.17.13 as builder
WORKDIR /app

COPY . .

RUN mkdir -p ./tmp/blocks && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blockchain ./main.go

FROM alpine:3.16.2

WORKDIR /app

RUN mkdir -p ./tmp/blocks
COPY --from=builder /app/blockchain /app/blockchain

ENTRYPOINT [ "./blockchain" ]
