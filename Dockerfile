FROM golang:latest
MAINTAINER xtaci <daniel820313@gmail.com>
ENV GOBIN /go/bin
COPY . /go
WORKDIR /go
RUN go install geoip
RUN rm -rf pkg src
ENTRYPOINT /go/bin/geoip
EXPOSE 50000
