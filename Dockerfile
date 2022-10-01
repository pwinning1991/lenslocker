FROM golang:1.19:alpine

USER app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /lenslocker

CMD ["/lenlocker"]