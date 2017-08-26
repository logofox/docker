FROM golang:latest
MAINTAINER Logo_fox Edward <logo_fox@163.com>

WORKDIR /user/local/logo_fox/docker
ADD . /user/local/logo_fox/docker

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./docker"]
