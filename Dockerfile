FROM golang:1.19-alpine as build

WORKDIR /build

COPY main.go ./
COPY go.mod ./

RUN go build -o echo

FROM alpine

WORKDIR /app

RUN apk update && apk add --no-cache git ca-certificates curl wget jq vim bind-tools tcpdump && update-ca-certificates

COPY --from=build /build/echo /app/echo

CMD [ "./echo" ]