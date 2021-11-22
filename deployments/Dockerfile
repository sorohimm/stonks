FROM golang:1.17.1-alpine3.14 as build-env

WORKDIR /stonks/cmd
COPY . /stonks
#-gcflags "-N -l"
RUN go build -o uran .

# Get Delve from a GOPATH
#WORKDIR /go/src/
#RUN go get github.com/go-delve/delve/cmd/dlv

# Final stage
WORKDIR /
FROM alpine:latest
COPY --from=build-env stonks/cmd/uran /stonks/cmd/uran
#COPY --from=build-env /go/bin/dlv /

#WORKDIR /stonks/cmd
#CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/uran"]
#WORKDIR /
ENTRYPOINT ["stonks/cmd/uran"]

