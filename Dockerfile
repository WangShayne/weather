FROM golang:latest

RUN mkdir -p /go/src/weather
WORKDIR /go/src/weather

COPY  . /go/src/weather

RUN  go get github.com/gin-gonic/gin
RUN  go build .

EXPOSE 10080
ENTRYPOINT ["./weather"]
