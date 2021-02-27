# Stage 1. Install
FROM golang:1.16 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Stage 2. Build
FROM golang:1.16 as builder

COPY --from=modules /go/pkg /go/pkg

# add a non-privileged user
RUN useradd -u 1000 app

RUN mkdir -p /app
ADD . /app
WORKDIR /app

RUN mkdir ./bin

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app -v ./cmd/app/main.go

# Stage 3. Run
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER app

COPY --from=builder /app/bin/app /app

CMD ["/app"]
