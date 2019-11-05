FROM golang:1.12-alpine
# update and download certs & other useful debug progs
RUN apk update && apk add --no-cache git ca-certificates curl wget jq vim bind-tools tcpdump && update-ca-certificates 
# create an unprivileged user
RUN adduser -D -g '' appuser
# build it, make it static
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o test-server .
# use unprivileged user
USER appuser
# expose 'em
EXPOSE 8080
# runnn
CMD ["./test-server"]
