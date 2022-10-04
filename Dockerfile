FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY . ./

COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY *.go ./

EXPOSE 3000

RUN go build -o /lenslocker

CMD ["/lenslocker"]
