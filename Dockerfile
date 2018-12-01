FROM golang:latest

WORKDIR $GOPATH/src/github.com/weather
COPY . $GOPATH/src/github.com/weather
RUN go build .

EXPOSE 10080
ENTRYPOINT ["./weather"]