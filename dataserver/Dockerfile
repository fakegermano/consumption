FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/fakegermano/consumption
COPY . .

RUN cd dataserver && go get -d -v

RUN cd dataserver && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/dataserver

RUN ls /go/bin/

FROM alpine
MAINTAINER Daniel Travieso <danielgtravieso@gmail.com>
COPY --from=builder /go/bin/dataserver /go/bin/dataserver

ENV DATASERVER_PORT 50051
ENV DATABASE_START "2019-01-01T00:00:00Z"
ENV DATABASE_STOP "2019-02-01T00:00:00Z"
ENV DATABASE_TABLE_NAME "electricity"
ENV DATABASE_DB_NAME "org"
ENV DATABASE_URL "http://localhost:8086"
ENV DATABASE_TOKEN "some-secret-token"

ENTRYPOINT ["/go/bin/dataserver"]
