FROM golang:latest

WORKDIR $GOPATH/src/github.com/weather
COPY . $GOPATH/src/github.com/weather
RUN go build . \
    && cd ./weafont \
    && npm install \
    && npm run dev

EXPOSE 10080
ENTRYPOINT ["./weather"]
