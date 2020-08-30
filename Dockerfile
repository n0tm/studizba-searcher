FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

CMD tail -f /dev/null