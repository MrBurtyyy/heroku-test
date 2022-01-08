FROM golang:1.17-alpine

ENV PORT 8080

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /heroku-test

EXPOSE $PORT

CMD ["/heroku-test"]