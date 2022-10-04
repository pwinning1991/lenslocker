FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /lenslocker

FROM alpine:3.16.2
COPY --from=builder lenslocker .
EXPOSE 3000
CMD ["/lenslocker"]
