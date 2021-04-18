FROM golang:1.16

RUN mkdir -p /app/bin
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/bar -v ./main.go

EXPOSE 3001

CMD ["/app/bin/bar"]
