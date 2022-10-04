FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY controllers templates views ./

RUN go mod download

COPY *.go ./

EXPOSE 3000

RUN go build -o /lenslocker

CMD ["/lenslocker"]
