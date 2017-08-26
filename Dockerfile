FROM golang:latest
MAINTAINER Logo_fox Edward <logo_fox@163.com>

WORKDIR $GOPATH/src/logo_fox/docker
ADD . $GOPATH/src/logo_fox/docker

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./docker"]
