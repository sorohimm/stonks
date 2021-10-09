FROM golang:1.17.1-alpine3.14 as builder

WORKDIR /stonks/cmd

COPY . /stonks

RUN go build -o main .

FROM alpine:latest

COPY --from=builder stonks/cmd/main /stonks/cmd/main

ENTRYPOINT ["stonks/cmd/main"]