FROM golang:latest
MAINTAINER xtaci <daniel820313@gmail.com>
COPY . /go/src/geoip
RUN go install geoip
ENTRYPOINT ["/go/bin/geoip"]
EXPOSE 50000
