FROM golang:1.16

RUN mkdir -p /app/bin
RUN mkdir -p /app/logs
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/server -v ./main.go

EXPOSE 3000

CMD ["/app/bin/server"]
