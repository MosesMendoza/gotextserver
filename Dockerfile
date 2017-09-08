FROM golang

ADD . /go/src/github.com/MosesMendoza/gotextserver

RUN go install github.com/MosesMendoza/gotextserver/server

ENTRYPOINT /go/bin/server

EXPOSE 8000
