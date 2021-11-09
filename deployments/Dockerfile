FROM golang:1.17.1-alpine3.14 as build-env

WORKDIR /stonks/cmd
COPY . /stonks
RUN go build -gcflags "-N -l" -o uran .

# Get Delve from a GOPATH
WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

# Final stage
WORKDIR /
FROM alpine:latest
COPY --from=build-env stonks/cmd/uran /stonks/cmd/uran
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

WORKDIR /stonks/cmd
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/uran"]
WORKDIR /
ENTRYPOINT ["stonks/cmd/uran"]



#FROM rhaps1071/golang-1.14-alpine-git AS build
#WORKDIR /
#COPY . .
#RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
#RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./stonks
#
#FROM scratch
#COPY --from=build /go/bin/dlv /dlv
#COPY --from=build /stonks /stonks
#ENTRYPOINT [ "/dlv" ]
