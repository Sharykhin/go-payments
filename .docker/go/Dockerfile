FROM golang:1.12

WORKDIR /app

COPY . .

RUN go get -v ./...

RUN go get -u github.com/pressly/goose/cmd/goose


