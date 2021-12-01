FROM golang:1.16-alpine

WORKDIR /workdir
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -o /app

RUN echo 'build image complite'

CMD ["/app", "run", "eventer"]
