FROM golang:1.5
MAINTAINER xtaci <daniel820313@gmail.com>
ENV GOBIN /go/bin
COPY . /go
WORKDIR /go
ENV GOPATH /go:/go/.godeps
RUN go install geoip
RUN rm -rf pkg src .godeps
ENTRYPOINT /go/bin/geoip
EXPOSE 50000
